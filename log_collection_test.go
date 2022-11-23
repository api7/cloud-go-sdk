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

func TestCreateLogCollection(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		name                 string
		pendingLogCollection *LogCollection
		expectedError        string
		mockFunc             func(t *testing.T) httpClient
	}{
		{
			name: "create successfully",
			pendingLogCollection: &LogCollection{
				Name: "test log collection",
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/log_collections"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli
			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingLogCollection: &LogCollection{
				Name: "test log collection",
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/log_collections"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli
			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCase {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newLogCollection(cli).CreateLogCollection(context.Background(), tc.pendingLogCollection, &ResourceCreateOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check log collection create error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestDeleteLogCollection(t *testing.T) {
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
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/log_collections/2"), "", nil).Return(nil)
				return cli
			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/log_collections/2"), "", nil).Return(errors.New("mock error"))
				return cli
			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			err := newLogCollection(cli).DeleteLogCollection(context.Background(), 2, &ResourceDeleteOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check delete log collection error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestGetLogCollection(t *testing.T) {
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
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/log_collections/2"), "", gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/log_collections/2"), "", gomock.Any()).Return(errors.New("mock error"))
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
			_, err := newLogCollection(cli).GetLogCollection(context.Background(), 2, &ResourceGetOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check get log collection detail error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestListLogCollection(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		iterator *logCollectionIterator
	}{
		{
			name: "create iterator successfully",
			iterator: &logCollectionIterator{
				iter: listIterator{
					resource: "logcollection",
					path:     "/api/v1/controlplanes/1/log_collections",
					paging: Pagination{
						Page:     2,
						PageSize: 14,
					},
					eof: false,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			raw, err := newLogCollection(nil).ListLogCollection(context.Background(), &ResourceListOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
				Pagination: &Pagination{
					Page:     2,
					PageSize: 14,
				},
			})
			assert.Nil(t, err, "check list log collection error")
			iter := raw.(*logCollectionIterator)
			assert.Equal(t, tc.iterator.iter.resource, iter.iter.resource, "check resource")
			assert.Equal(t, tc.iterator.iter.path, iter.iter.path, "check path")
			assert.Equal(t, tc.iterator.iter.paging.Page, iter.iter.paging.Page, "check page")
			assert.Equal(t, tc.iterator.iter.paging.PageSize, iter.iter.paging.PageSize, "check page size")
		})
	}
}

func TestPutLogCollection(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		name                 string
		pendingLogCollection *LogCollection
		expectedError        string
		mockFunc             func(t *testing.T) httpClient
	}{
		{
			name: "update successfully",
			pendingLogCollection: &LogCollection{
				Name: "test log collection",
				ID:   2,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/log_collections/2"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli
			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingLogCollection: &LogCollection{
				Name: "test log collection",
				ID:   2,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/log_collections/2"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli
			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCase {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newLogCollection(cli).UpdateLogCollection(context.Background(), tc.pendingLogCollection, &ResourceUpdateOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check log collection create error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}
