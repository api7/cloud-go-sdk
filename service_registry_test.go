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

func TestCreateServiceRegistry(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		pendingRegistry *ServiceRegistry
		expectedError   string
		mockFunc        func(t *testing.T) httpClient
	}{
		{
			name: "create successfully",
			pendingRegistry: &ServiceRegistry{
				ServiceRegistrySpec: ServiceRegistrySpec{
					Name:       "test registry",
					Enabled:    true,
					Type:       ServiceRegistryKubernetes,
					Kubernetes: &KubernetesServiceRegistry{},
				},
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/service_registries"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingRegistry: &ServiceRegistry{
				ServiceRegistrySpec: ServiceRegistrySpec{
					Name:       "test registry",
					Enabled:    true,
					Type:       ServiceRegistryKubernetes,
					Kubernetes: &KubernetesServiceRegistry{},
				},
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/service_registries"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newServiceDiscovery(cli, &store{}).CreateServiceRegistry(context.Background(), tc.pendingRegistry, &ResourceCreateOptions{
				Cluster: &Cluster{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check service registry create error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestUpdateServiceRegistry(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		pendingRegistry *ServiceRegistry
		expectedError   string
		mockFunc        func(t *testing.T) httpClient
	}{
		{
			name: "create successfully",
			pendingRegistry: &ServiceRegistry{
				ServiceRegistrySpec: ServiceRegistrySpec{
					Name:       "test registry",
					Enabled:    true,
					Type:       ServiceRegistryKubernetes,
					Kubernetes: &KubernetesServiceRegistry{},
				},
				ID: 12,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/service_registries/12"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingRegistry: &ServiceRegistry{
				ServiceRegistrySpec: ServiceRegistrySpec{
					Name:       "test registry",
					Enabled:    true,
					Type:       ServiceRegistryKubernetes,
					Kubernetes: &KubernetesServiceRegistry{},
				},
				ID: 12,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/service_registries/12"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newServiceDiscovery(cli, &store{}).UpdateServiceRegistry(context.Background(), tc.pendingRegistry, &ResourceUpdateOptions{
				Cluster: &Cluster{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check service registry update error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestDeleteServiceRegistry(t *testing.T) {
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
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/service_registries/12"), "", nil, gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/service_registries/12"), "", nil, gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			err := newServiceDiscovery(cli, &store{}).DeleteServiceRegistry(context.Background(), 12, &ResourceDeleteOptions{
				Cluster: &Cluster{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check service discovery delete error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestGetServiceRegistry(t *testing.T) {
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
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/service_registries/12"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/clusters/1/service_registries/12"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
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
			_, err := newServiceDiscovery(cli, &store{}).GetServiceRegistry(context.Background(), 12, &ResourceGetOptions{
				Cluster: &Cluster{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check service registry get error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestListServiceRegistries(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		iterator *serviceRegistryListIterator
	}{
		{
			name: "create iterator successfully",
			iterator: &serviceRegistryListIterator{
				iter: listIterator{
					resource: "service_registry",
					path:     "/api/v1/clusters/123/service_registries",
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
			// ignore the service registry check since currently we don't mock it, and the app is always a zero value.
			raw, err := newServiceDiscovery(nil, &store{}).ListServiceRegistries(context.Background(), &ResourceListOptions{
				Cluster: &Cluster{
					ID: 123,
				},
				Pagination: &Pagination{
					Page:     14,
					PageSize: 25,
				},
			})
			assert.Nil(t, err, "check list application error")
			iter := raw.(*serviceRegistryListIterator)
			assert.Equal(t, tc.iterator.iter.resource, iter.iter.resource, "check resource")
			assert.Equal(t, tc.iterator.iter.path, iter.iter.path, "check path")
			assert.Equal(t, tc.iterator.iter.paging.Page, iter.iter.paging.Page, "check page")
			assert.Equal(t, tc.iterator.iter.paging.PageSize, iter.iter.paging.PageSize, "check page size")
		})
	}
}
