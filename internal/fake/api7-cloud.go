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

// Package fake provide fake objects. This package is used only for test purpose.
package fake

import (
	"github.com/pkg/errors"
	"net"
	"net/http"
	"net/url"

	"golang.org/x/net/nettest"
)

// API7Cloud is a fake API7 Cloud server side implementation.
type API7Cloud struct {
	listener net.Listener

	expect map[string]expect
}

type expect struct {
	statusCode int
	body       []byte
}

// NewAPI7Cloud creates a fake API7 Cloud server.
func NewAPI7Cloud() (*API7Cloud, error) {
	listener, err := nettest.NewLocalListener("tcp")
	if err != nil {
		return nil, errors.Wrap(err, "new local listener")
	}

	api7 := &API7Cloud{
		listener: listener,
		expect:   make(map[string]expect),
	}

	return api7, nil
}

// Expect defines an expected response (status code, response body) for an URI.
func (api7 *API7Cloud) Expect(uri string, statusCode int, body []byte) {
	api7.expect[uri] = expect{
		statusCode: statusCode,
		body:       body,
	}
}

// Addr returns the API7 Cloud addr.
func (api7 *API7Cloud) Addr() string {
	url := url.URL{
		Scheme: "http",
		Host:   api7.listener.Addr().String(),
	}
	return url.String()
}

// Serve starts to accept HTTP requests.
func (api7 *API7Cloud) Serve() error {
	return http.Serve(api7.listener, api7)
}

// Close closes the API7 Cloud.
func (api7 *API7Cloud) Close() error {
	return api7.listener.Close()
}

// ServeHTTP implements the HTTP.Handler interface.
func (api7 *API7Cloud) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	exp, ok := api7.expect[req.URL.Path]
	if !ok {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	rw.WriteHeader(exp.statusCode)
	_, _ = rw.Write(exp.body)
}
