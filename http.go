// Copyright 2022 API7.ai, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloud

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net"
	"net/http"
	"net/url"
)

type payloadDecodeFunc func(data json.RawMessage) error

// responseWrapper wraps the response with its original payload,
// and sets the Status field to codes.OK if everything is OK, but when
// the response is invalid, ErrorReason could be filled to show the error
// details and in such a case, Status is not codes.OK but a specific error
// code to show the kind.
type responseWrapper struct {
	// Payload carries the original data.
	Payload json.RawMessage `json:"payload,omitempty"`
	// Status shows the operation status for current request.
	Status status `json:"status"`
	// ErrorReason is the error details, it's exclusive with Payload.
	ErrorReason string `json:"error,omitempty"`
	// Warning attaches a warning message to the response.
	Warning string `json:"warning,omitempty"`
}

// status represents an error type, it contains the error code and its
// description.
type status struct {
	// Code is the error code number.
	// example: 0
	Code int `json:"code"`
	// Message describes the error code.
	// example: OK
	Message string `json:"message"`
}

func jsonPayloadDecodeFactory(container interface{}) payloadDecodeFunc {
	return func(data json.RawMessage) error {
		if err := json.Unmarshal(data, container); err != nil {
			return errors.Wrap(err, "decode payload to json")
		}
		return nil
	}
}

// httpClient is an interface which abstracts behaviors that the Cloud Go SDK
// will perform.
type httpClient interface {
	sendGetRequest(path, query string, payloadDecodeFunc payloadDecodeFunc) error
	sendPostRequest(path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error
	sendPutRequest(path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error
	sendDeleteRequest(path, query string, payloadDecodeFunc payloadDecodeFunc) error
	sendRequest(req *http.Request, payloadDecodeFunc payloadDecodeFunc) error
}

type httpClientImpl struct {
	url    *url.URL
	token  *AccessToken
	client *http.Client
}

func constructHTTPClient(opts *Options, token *AccessToken) (httpClient, error) {
	u, err := url.Parse(opts.ServerAddr)
	if err != nil {
		return nil, errors.Wrap(err, "construct http client")
	}

	tr := &http.Transport{
		DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, opts.DialTimeout)
		},

		TLSClientConfig: &tls.Config{
			ServerName:         opts.ServerNameIndication,
			InsecureSkipVerify: opts.InsecureSkipTLSVerify,
			MinVersion:         tls.VersionTLS11,
			MaxVersion:         tls.VersionTLS13,
		},
		TLSHandshakeTimeout: opts.TLSHandshakeTimeout,
	}

	return &httpClientImpl{
		url: u,
		client: &http.Client{
			Transport: tr,
		},
		token: token,
	}, nil
}

func (impl *httpClientImpl) sendGetRequest(path, query string, payloadDecodeFunc payloadDecodeFunc) error {
	u := *impl.url
	u.Path = path
	u.RawQuery = query

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return errors.Wrap(err, "construct http request")
	}
	return impl.sendRequest(req, payloadDecodeFunc)
}

func (impl *httpClientImpl) sendPostRequest(path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error {
	u := *impl.url
	u.Path = path
	u.RawQuery = query

	data, err := json.Marshal(body)
	if err != nil {
		return errors.Wrap(err, "encode http request body")
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(data))
	if err != nil {
		return errors.Wrap(err, "construct http request")
	}

	req.Header.Set("Content-Type", "application/json")
	return impl.sendRequest(req, payloadDecodeFunc)
}

func (impl *httpClientImpl) sendPutRequest(path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error {
	u := *impl.url
	u.Path = path
	u.RawQuery = query

	data, err := json.Marshal(body)
	if err != nil {
		return errors.Wrap(err, "encode http request body")
	}

	req, err := http.NewRequest(http.MethodPut, u.String(), bytes.NewReader(data))
	if err != nil {
		return errors.Wrap(err, "construct http request")
	}

	req.Header.Set("Content-Type", "application/json")
	return impl.sendRequest(req, payloadDecodeFunc)
}

func (impl *httpClientImpl) sendDeleteRequest(path, query string, payloadDecodeFunc payloadDecodeFunc) error {
	u := *impl.url
	u.Path = path
	u.RawQuery = query

	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return errors.Wrap(err, "construct http request")
	}

	return impl.sendRequest(req, payloadDecodeFunc)
}

func (impl *httpClientImpl) sendRequest(req *http.Request, payloadDecodeFunc payloadDecodeFunc) error {
	var (
		rw responseWrapper
	)

	token := "Bearer " + impl.token.Token
	req.Header.Set("Authorization", token)

	resp, err := impl.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "send http request")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read http response body")
	}

	// API7 Cloud won't encapsulate response body into the ResponseWrapper,
	// so for these responses, we handle it alone.
	if resp.StatusCode >= 500 {
		return fmt.Errorf("status code: %d, message: %s", resp.StatusCode, string(body))
	}

	if err = json.Unmarshal(body, &rw); err != nil {
		return errors.Wrap(err, "decode response body")
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code: %d, error code: %d, error reason: %s, details: %s", resp.StatusCode, rw.Status.Code, rw.Status.Message, rw.ErrorReason)
	}

	if payloadDecodeFunc != nil {
		if err = payloadDecodeFunc(rw.Payload); err != nil {
			return err
		}
	}
	return nil
}
