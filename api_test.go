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

func TestCreateAPI(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		pendingAPI    *API
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "create successfully",
			pendingAPI: &API{
				APISpec: APISpec{
					Name: "test api",
				},
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingAPI: &API{
				APISpec: APISpec{
					Name: "test api",
				},
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newAPI(cli).CreateAPI(context.Background(), tc.pendingAPI, &ResourceCreateOptions{
				Application: &Application{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check api create error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestUpdateAPI(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		pendingAPI    *API
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "update successfully",
			pendingAPI: &API{
				APISpec: APISpec{
					Name: "test api",
				},
				ID: 12,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis/12"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingAPI: &API{
				APISpec: APISpec{
					Name: "test api",
				},
				ID: 12,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis/12"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newAPI(cli).UpdateAPI(context.Background(), tc.pendingAPI, &ResourceUpdateOptions{
				Application: &Application{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check api update error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestDeleteAPI(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "delete successfully",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis/12"), "", nil).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis/12"), "", nil).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			err := newAPI(cli).DeleteAPI(context.Background(), 12, &ResourceDeleteOptions{
				Application: &Application{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check api delete error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestGetAPI(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "get successfully",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis/12"), "", gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis/12"), "", gomock.Any()).Return(errors.New("mock error"))
				return cli
			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			// ignore the application check since currently we don't mock it, and the app is always a zero value.
			_, err := newAPI(cli).GetAPI(context.Background(), 12, &ResourceGetOptions{
				Application: &Application{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check api get error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestPublishAPI(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "publish successfully",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPatchRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis/12"), "", []byte(`{"active":0}`), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPatchRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis/12"), "", []byte(`{"active":0}`), gomock.Any()).Return(errors.New("mock error"))
				return cli
			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			// ignore the api check since currently we don't mock it, and the app is always a zero value.
			_, err := newAPI(cli).PublishAPI(context.Background(), 12, &ResourceUpdateOptions{
				Application: &Application{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check api publish error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestUnpublishAPI(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "publish successfully",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPatchRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis/12"), "", []byte(`{"active":1}`), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPatchRequest(gomock.Any(), path.Join(_apiPathPrefix, "/apps/1/apis/12"), "", []byte(`{"active":1}`), gomock.Any()).Return(errors.New("mock error"))
				return cli
			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			// ignore the api check since currently we don't mock it, and the app is always a zero value.
			_, err := newAPI(cli).UnpublishAPI(context.Background(), 12, &ResourceUpdateOptions{
				Application: &Application{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check api publish error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}
