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
	"io"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/url"

	"github.com/pkg/errors"
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
	sendGetRequest(ctx context.Context, path, query string, payloadDecodeFunc payloadDecodeFunc) error
	sendPostRequest(ctx context.Context, path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error
	sendPutRequest(ctx context.Context, path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error
	sendPatchRequest(ctx context.Context, path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error
	sendDeleteRequest(ctx context.Context, path, query string, payloadDecodeFunc payloadDecodeFunc) error
	sendRequest(req *http.Request, payloadDecodeFunc payloadDecodeFunc) error
}

type httpClientImpl struct {
	url             *url.URL
	token           *AccessToken
	client          *http.Client
	tracer          TraceInterface
	idGenerator     IDGenerator
	genIDForCalls   bool
	enableHTTPTrace bool
}

type httpClientConstructOptions struct {
	configOptions *Options
	token         *AccessToken
	tracer        TraceInterface
	idGenerator   IDGenerator
}

func constructHTTPClient(opts *httpClientConstructOptions) (httpClient, error) {
	url, err := url.Parse(opts.configOptions.ServerAddr)
	if err != nil {
		return nil, errors.Wrap(err, "construct http client")
	}

	tr := &http.Transport{
		DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, opts.configOptions.DialTimeout)
		},

		TLSClientConfig: &tls.Config{
			ServerName:         opts.configOptions.ServerNameIndication,
			InsecureSkipVerify: opts.configOptions.InsecureSkipTLSVerify,
			MinVersion:         tls.VersionTLS11,
			MaxVersion:         tls.VersionTLS13,
		},
		TLSHandshakeTimeout: opts.configOptions.TLSHandshakeTimeout,
	}

	if opts.configOptions.ClientCert != "" && opts.configOptions.ClientPrivateKey != "" {
		cert, err := tls.X509KeyPair([]byte(opts.configOptions.ClientCert), []byte(opts.configOptions.ClientPrivateKey))
		if err != nil {
			return nil, errors.Wrap(err, "load client certificate")
		}
		tr.TLSClientConfig.Certificates = []tls.Certificate{cert}
	}

	return &httpClientImpl{
		url: url,
		client: &http.Client{
			Transport: tr,
		},
		token:           opts.token,
		tracer:          opts.tracer,
		idGenerator:     opts.idGenerator,
		genIDForCalls:   opts.configOptions.GenIDForCalls,
		enableHTTPTrace: opts.configOptions.EnableHTTPTrace,
	}, nil
}

func newClientTrace(series *TraceSeries) *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			ev := generateEvent("plan to connect to %s", hostPort)
			series.appendEvent(ev)
		},
		GotConn: func(info httptrace.GotConnInfo) {
			ev := generateEvent("connected to %s", info.Conn.RemoteAddr())
			series.appendEvent(ev)
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			ev := generateEvent("plan to resolve domain %s", info.Host)
			series.appendEvent(ev)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			ev := generateEvent("resolved domain, ip addrs: %v, error: %v", info.Addrs, info.Err)
			series.appendEvent(ev)
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			ev := generateEvent("TLS handshake done, sni: %s, success: %v", state.ServerName, state.HandshakeComplete)
			series.appendEvent(ev)
		},
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			ev := generateEvent("request sent, error: %v", info.Err)
			series.appendEvent(ev)
		},
	}
}

func (impl *httpClientImpl) sendGetRequest(ctx context.Context, path, query string, payloadDecodeFunc payloadDecodeFunc) error {
	url := *impl.url
	url.Path = path
	url.RawQuery = query

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return errors.Wrap(err, "construct http request")
	}

	if impl.enableHTTPTrace {
		series := &TraceSeries{
			ID:      impl.idGenerator.NextID(),
			Request: req.Clone(context.TODO()),
		}
		req = req.WithContext(context.WithValue(req.Context(), TraceSeriesKey{}, series))
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), newClientTrace(series)))
		defer func() {
			impl.tracer.sendSeries(series)
		}()
	}

	return impl.sendRequest(req, payloadDecodeFunc)
}

func (impl *httpClientImpl) sendPostRequest(ctx context.Context, path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error {
	url := *impl.url
	url.Path = path
	url.RawQuery = query

	data, err := json.Marshal(body)
	if err != nil {
		return errors.Wrap(err, "encode http request body")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), bytes.NewReader(data))
	if err != nil {
		return errors.Wrap(err, "construct http request")
	}

	if impl.enableHTTPTrace {
		series := &TraceSeries{
			ID:          impl.idGenerator.NextID(),
			Request:     req.Clone(context.TODO()),
			RequestBody: data,
		}
		req = req.WithContext(context.WithValue(req.Context(), TraceSeriesKey{}, series))
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), newClientTrace(series)))
		defer func() {
			impl.tracer.sendSeries(series)
		}()
	}

	req.Header.Set("Content-Type", "application/json")
	return impl.sendRequest(req, payloadDecodeFunc)
}

func (impl *httpClientImpl) sendPutRequest(ctx context.Context, path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error {
	var (
		reader io.Reader
		data   []byte
		err    error
	)

	url := *impl.url
	url.Path = path
	url.RawQuery = query

	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return errors.Wrap(err, "encode http request body")
		}
		reader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), reader)
	if err != nil {
		return errors.Wrap(err, "construct http request")
	}

	if impl.enableHTTPTrace {
		series := &TraceSeries{
			ID:          impl.idGenerator.NextID(),
			Request:     req.Clone(context.TODO()),
			RequestBody: data,
		}
		req = req.WithContext(context.WithValue(req.Context(), TraceSeriesKey{}, series))
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), newClientTrace(series)))
		defer func() {
			impl.tracer.sendSeries(series)
		}()
	}

	req.Header.Set("Content-Type", "application/json")
	return impl.sendRequest(req, payloadDecodeFunc)
}

func (impl *httpClientImpl) sendPatchRequest(ctx context.Context, path, query string, body interface{}, payloadDecodeFunc payloadDecodeFunc) error {
	url := *impl.url
	url.Path = path
	url.RawQuery = query

	data, err := json.Marshal(body)
	if err != nil {
		return errors.Wrap(err, "encode http request body")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url.String(), bytes.NewReader(data))
	if err != nil {
		return errors.Wrap(err, "construct http request")
	}

	if impl.enableHTTPTrace {
		series := &TraceSeries{
			ID:          impl.idGenerator.NextID(),
			Request:     req.Clone(context.TODO()),
			RequestBody: data,
		}
		req = req.WithContext(context.WithValue(req.Context(), TraceSeriesKey{}, series))
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), newClientTrace(series)))
		defer func() {
			impl.tracer.sendSeries(series)
		}()
	}

	req.Header.Set("Content-Type", "application/json")
	return impl.sendRequest(req, payloadDecodeFunc)
}

func (impl *httpClientImpl) sendDeleteRequest(ctx context.Context, path, query string, payloadDecodeFunc payloadDecodeFunc) error {
	url := *impl.url
	url.Path = path
	url.RawQuery = query

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)
	if err != nil {
		return errors.Wrap(err, "construct http request")
	}

	if impl.enableHTTPTrace {
		series := &TraceSeries{
			ID:      impl.idGenerator.NextID(),
			Request: req.Clone(context.TODO()),
		}
		req = req.WithContext(context.WithValue(req.Context(), TraceSeriesKey{}, series))
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), newClientTrace(series)))
		defer func() {
			impl.tracer.sendSeries(series)
		}()
	}

	return impl.sendRequest(req, payloadDecodeFunc)
}

func (impl *httpClientImpl) sendRequest(req *http.Request, payloadDecodeFunc payloadDecodeFunc) error {
	var (
		rw       responseWrapper
		errTrace error
		resp     *http.Response
		body     []byte
	)

	token := "Bearer " + impl.token.Token
	req.Header.Set("Authorization", token)

	if impl.genIDForCalls {
		req.Header.Set("X-Request-ID", impl.idGenerator.NextID().String())
	}

	if impl.enableHTTPTrace {
		series, ok := req.Context().Value(TraceSeriesKey{}).(*TraceSeries)
		deferFunc := func() {
			series.Response = resp
			series.ResponseBody = body
			if errTrace == nil {
				return
			}

			ev := generateEvent("response error %s", errTrace)
			series.appendEvent(ev)
		}
		if ok {
			defer deferFunc()
		}
	}

	resp, err := impl.client.Do(req)
	if err != nil {
		errTrace = errors.Wrap(err, "send http request")
		return errTrace
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		errTrace = errors.Wrap(err, "read http response body")
		return errTrace
	}
	// API7 Cloud won't encapsulate response body into the ResponseWrapper,
	// so for these responses, we handle it alone.
	if resp.StatusCode >= 500 {
		errTrace = fmt.Errorf("status code: %d, message: %s", resp.StatusCode, string(body))
		return errTrace
	}

	if err = json.Unmarshal(body, &rw); err != nil {
		errTrace = errors.Wrap(err, "decode response body")
		return errTrace
	}

	if resp.StatusCode != 200 {
		errTrace = fmt.Errorf("status code: %d, error code: %d, error reason: %s, details: %s", resp.StatusCode, rw.Status.Code, rw.Status.Message, rw.ErrorReason)
		return errTrace
	}

	if payloadDecodeFunc != nil {
		if err = payloadDecodeFunc(rw.Payload); err != nil {
			errTrace = err
			return errTrace
		}
	}
	return nil
}
