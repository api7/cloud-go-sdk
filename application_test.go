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
	"errors"
	"path"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateApplication(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		pendingApp    *Application
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "create successfully",
			pendingApp: &Application{
				ApplicationSpec: ApplicationSpec{
					Name:       "test app",
					Labels:     nil,
					Protocols:  []string{ProtocolHTTP},
					PathPrefix: "/v1",
				},
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/apps"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingApp: &Application{
				ApplicationSpec: ApplicationSpec{
					Name:       "test app",
					Labels:     nil,
					Protocols:  []string{ProtocolHTTP},
					PathPrefix: "/v1",
				},
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/apps"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newApplication(cli).CreateApplication(context.Background(), tc.pendingApp, &ApplicationCreateOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check application create error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}
