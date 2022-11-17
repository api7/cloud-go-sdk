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

import "time"

var (
	// DefaultOptions contains the default settings for all the options.
	DefaultOptions = &Options{
		ServerAddr:            "https://api.api7.cloud",
		Token:                 "",
		TokenPath:             "",
		DialTimeout:           5 * time.Second,
		InsecureSkipTLSVerify: false,
		ServerNameIndication:  "",
		EnableHTTPTrace:       false,
	}
)

// Options contains all related configurations for the SDK to communicate with API7 Cloud.
type Options struct {
	// ServerAddr indicates the URL for accessing API7 Cloud API.
	// e.g. https://api.api7.cloud.
	ServerAddr string `json:"server_addr" yaml:"server_addr"`
	// Token is a personal access token for accessing API7 Cloud API.
	// You can skip filling this field in turn using `TokenPath` field to
	// configure the token from filesystem.
	// Note, when you configure both of the `Token` and `TokenPath` field, the `Token`
	// field takes the precedence.
	Token string `json:"token" yaml:"token"`
	// TokenPath indicates the filepath where the API7 Cloud access token stores.
	// You can skip filling this field in turn using `Token` field to configure
	// the token literally.
	// Note, when you configure both of the `Token` and `TokenPath` field, the `Token`
	// field takes the precedence.
	TokenPath string `json:"token_path" yaml:"token_path"`
	// DialTimeout indicates the timeout for the TCP handshake.
	DialTimeout time.Duration `json:"dial_timeout" yaml:"dial_timeout"`
	// TLSHandshakeTimeout indicates the timeout for TLS handshake.
	TLSHandshakeTimeout time.Duration `json:"tls_handshake_timeout" yaml:"tls_handshake_timeout"`
	// InsecureSkipTLSVerify indicates if Cloud Go SDK should skip verifying
	// API7 Cloud server's TLS certificate.
	InsecureSkipTLSVerify bool `json:"insecure_skip_tls_verify" yaml:"insecure_skip_tls_verify"`
	// ServerNameIndication indicates the TLS SNI extension.
	ServerNameIndication string `json:"server_name_indication" yaml:"server_name_indication"`
	// EnableHTTPTrace indicate if collect events occur during an HTTP call.
	// Events are generated by the net/http/httptrace package.
	// Note, set this field to true might cause more memory usage. Please only
	// enable it if you want to troubleshoot some problems.
	EnableHTTPTrace bool `json:"enable_http_trace" yaml:"enable_http_trace"`
	// GenIDForCalls indicates if an ID (generated by snowflake algorithm)
	// should be added for each API requests.
	// Note, the ID will be put in the `X-Request-ID` header.
	GenIDForCalls bool `json:"gen_id_for_calls" yaml:"gen_id_for_calls"`
}

func (o *Options) merge(o2 *Options) {
	if o.ServerAddr == "" {
		o.ServerAddr = o2.ServerAddr
	}
	if o.Token == "" {
		o.Token = o2.Token
	}
	if o.DialTimeout == 0 {
		o.DialTimeout = o2.DialTimeout
	}
	if !o.InsecureSkipTLSVerify {
		o.InsecureSkipTLSVerify = o2.InsecureSkipTLSVerify
	}
	if o.ServerNameIndication == "" {
		o.ServerNameIndication = o2.ServerNameIndication
	}
}
