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
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
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
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newApplication(cli, &store{}).CreateApplication(context.Background(), tc.pendingApp, &ResourceCreateOptions{
				Cluster: &Cluster{
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

func TestUpdateApplication(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		pendingApp    *Application
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "update successfully",
			pendingApp: &Application{
				ApplicationSpec: ApplicationSpec{
					Name:       "test app",
					Labels:     nil,
					Protocols:  []string{ProtocolHTTP},
					PathPrefix: "/v1",
				},
				ID: 12,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps/12"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
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
				ID: 12,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps/12"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newApplication(cli, &store{}).UpdateApplication(context.Background(), tc.pendingApp, &ResourceUpdateOptions{
				Cluster: &Cluster{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check application update error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestDeleteApplication(t *testing.T) {
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
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps/12"), "", nil, gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps/12"), "", nil, gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			err := newApplication(cli, &store{}).DeleteApplication(context.Background(), 12, &ResourceDeleteOptions{
				Cluster: &Cluster{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check application delete error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestGetApplication(t *testing.T) {
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
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps/12"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps/12"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
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
			_, err := newApplication(cli, &store{}).GetApplication(context.Background(), 12, &ResourceGetOptions{
				Cluster: &Cluster{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check application get error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestPublishApplication(t *testing.T) {
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
				cli.EXPECT().sendPatchRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps/12"), "", []byte(`{"active":0}`), gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPatchRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps/12"), "", []byte(`{"active":0}`), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
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
			_, err := newApplication(cli, &store{}).PublishApplication(context.Background(), 12, &ResourceUpdateOptions{
				Cluster: &Cluster{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check application publish error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestUnpublishApplication(t *testing.T) {
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
				cli.EXPECT().sendPatchRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps/12"), "", []byte(`{"active":1}`), gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPatchRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/apps/12"), "", []byte(`{"active":1}`), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
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
			_, err := newApplication(cli, &store{}).UnpublishApplication(context.Background(), 12, &ResourceUpdateOptions{
				Cluster: &Cluster{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check application publish error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestListApplications(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		iterator *applicationListIterator
	}{
		{
			name: "create iterator successfully",
			iterator: &applicationListIterator{
				iter: listIterator{
					resource: "application",
					path:     "/api/v1/clusters/123/apps",
					paging: Pagination{
						Page:     14,
						PageSize: 25,
					},
					eof: false,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// ignore the application check since currently we don't mock it, and the app is always a zero value.
			raw, err := newApplication(nil, &store{}).ListApplications(context.Background(), &ResourceListOptions{
				Cluster: &Cluster{
					ID: 123,
				},
				Pagination: &Pagination{
					Page:     14,
					PageSize: 25,
				},
			})
			assert.Nil(t, err, "check list application error")
			iter := raw.(*applicationListIterator)
			assert.Equal(t, tc.iterator.iter.resource, iter.iter.resource, "check resource")
			assert.Equal(t, tc.iterator.iter.path, iter.iter.path, "check path")
			assert.Equal(t, tc.iterator.iter.paging.Page, iter.iter.paging.Page, "check page")
			assert.Equal(t, tc.iterator.iter.paging.PageSize, iter.iter.paging.PageSize, "check page size")
		})
	}
}
