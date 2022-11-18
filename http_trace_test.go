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
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/api7/cloud-go-sdk/internal/fake"
)

func TestHTTPTrace(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		enableTrace         bool
		createFakeAPI7Cloud func(t *testing.T) *fake.API7Cloud
		expectedError       string
	}{
		{
			name:        "enable http trace",
			enableTrace: true,
			createFakeAPI7Cloud: func(t *testing.T) *fake.API7Cloud {
				api7, err := fake.NewAPI7Cloud()
				assert.Nil(t, err, "check create fake api7 cloud error")
				return api7
			},
			expectedError: "",
		},
		{
			name:        "disable http trace",
			enableTrace: false,
			createFakeAPI7Cloud: func(t *testing.T) *fake.API7Cloud {
				api7, err := fake.NewAPI7Cloud()
				assert.Nil(t, err, "check create fake api7 cloud error")
				return api7
			},
			expectedError: "",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			api7 := tc.createFakeAPI7Cloud(t)
			go func() {
				_ = api7.Serve()
			}()
			defer func() {
				_ = api7.Close()
			}()
			app := `
{
  "status": {
    "code": 0,
    "message": "OK"
  },
  "payload": {
    "id": "123",
    "name": "first app",
	"path_prefix": "/v1",
	"protocols": ["HTTP", "HTTPS"]
  }
}
`
			api7.Expect("/api/v1/controlplanes/123/apps", http.StatusOK, []byte(app))

			sdk, err := NewInterface(&Options{
				ServerAddr:      api7.Addr(),
				Token:           "fake token",
				EnableHTTPTrace: tc.enableTrace,
			})
			assert.Nil(t, err)

			var (
				seriel *TraceSeries
			)
			received := make(chan struct{})

			go func() {
				seriel = <-sdk.TraceChan()
				close(received)
			}()

			_, err = sdk.CreateApplication(context.Background(), &Application{
				ApplicationSpec: ApplicationSpec{
					Name:       "first app",
					Protocols:  []string{ProtocolHTTP, ProtocolHTTPS},
					PathPrefix: "/v1",
				},
			}, &ResourceCreateOptions{
				ControlPlane: &ControlPlane{
					ID: 123,
				},
			})
			assert.Nil(t, err, "check application create result")
			select {
			case <-received:
				if !tc.enableTrace {
					assert.Fail(t, "received trace series when enable http trace is disabled")
				}
				assert.Len(t, seriel.Events, 3, "check events number")
				assert.NotEqual(t, int(seriel.ID), 0, "check if the ID is not zero")
				assert.Contains(t, seriel.Events[0].Message, "plan to connect to 127.0.0.1:", "check first event")
				assert.Contains(t, seriel.Events[1].Message, "connected to 127.0.0.1:", "check second event")
				assert.Contains(t, seriel.Events[2].Message, "request sent", "check third event")
				assert.Equal(t, http.MethodPost, seriel.Request.Method, "check request method")
				assert.Equal(t, "/api/v1/controlplanes/123/apps", seriel.Request.URL.Path, "check request URI path")
			case <-time.After(time.Second):
				if tc.enableTrace {
					assert.Fail(t, "didn't receive trace series when enable http trace is enabled")
				}
			}
		})
	}
}
