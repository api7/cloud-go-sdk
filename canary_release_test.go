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
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func TestCreateCanaryRelease(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		pendingCr     *CanaryRelease
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "create successfully",
			pendingCr: &CanaryRelease{
				CanaryReleaseSpec: CanaryReleaseSpec{
					Name:                  "test canary release",
					State:                 "pause",
					Type:                  "percent",
					CanaryUpstreamVersion: "v1",
					Percent:               50,
				},
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/canary_releases"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli
			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingCr: &CanaryRelease{
				CanaryReleaseSpec: CanaryReleaseSpec{
					Name:                  "test canary release",
					State:                 "pause",
					Type:                  "percent",
					CanaryUpstreamVersion: "v1",
					Percent:               50,
				},
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/canary_releases"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli
			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newCanaryRelease(cli).CreateCanaryRelease(context.Background(), tc.pendingCr, &ResourceCreateOptions{
				Application: &Application{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check canary release create error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}
