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

func TestCreateConsumer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		pendingConsumer *Consumer
		expectedError   string
		mockFunc        func(t *testing.T) httpClient
	}{
		{
			name: "create successfully",
			pendingConsumer: &Consumer{
				Name: "test consumer",
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/consumers"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingConsumer: &Consumer{
				Name: "test consumer",
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/consumers"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newConsumer(cli).CreateConsumer(context.Background(), tc.pendingConsumer, &ResourceCreateOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check consumer create error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestUpdateConsumer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		pendingConsumer *Consumer
		expectedError   string
		mockFunc        func(t *testing.T) httpClient
	}{
		{
			name: "update successfully",
			pendingConsumer: &Consumer{
				Name: "test consumer",
				ID:   12,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/consumers/12"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingConsumer: &Consumer{
				Name: "test consumer",
				ID:   12,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/consumers/12"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newConsumer(cli).UpdateConsumer(context.Background(), tc.pendingConsumer, &ResourceUpdateOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check consumer update error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestDeleteConsumer(t *testing.T) {
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
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/consumers/12"), "", nil).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/consumers/12"), "", nil).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			err := newConsumer(cli).DeleteConsumer(context.Background(), 12, &ResourceDeleteOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check consumer delete error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestGetConsumer(t *testing.T) {
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
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/consumers/12"), "", gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/consumers/12"), "", gomock.Any()).Return(errors.New("mock error"))
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
			_, err := newConsumer(cli).GetConsumer(context.Background(), 12, &ResourceGetOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check consumer get error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestListConsumers(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		iterator *consumerListIterator
	}{
		{
			name: "create iterator successfully",
			iterator: &consumerListIterator{
				iter: listIterator{
					resource: "consumer",
					path:     "/api/v1/controlplanes/123/consumers",
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
			// ignore the consumer check since currently we don't mock it, and the app is always a zero value.
			raw, err := newConsumer(nil).ListConsumers(context.Background(), &ResourceListOptions{
				ControlPlane: &ControlPlane{
					ID: 123,
				},
				Pagination: &Pagination{
					Page:     14,
					PageSize: 25,
				},
			})
			assert.Nil(t, err, "check list consumer error")
			iter := raw.(*consumerListIterator)
			assert.Equal(t, tc.iterator.iter.resource, iter.iter.resource, "check resource")
			assert.Equal(t, tc.iterator.iter.path, iter.iter.path, "check path")
			assert.Equal(t, tc.iterator.iter.paging.Page, iter.iter.paging.Page, "check page")
			assert.Equal(t, tc.iterator.iter.paging.PageSize, iter.iter.paging.PageSize, "check page size")
		})
	}
}
