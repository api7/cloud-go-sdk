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

func TestGetOrganization(t *testing.T) {
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
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/12"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/12"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
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
			_, err := newOrganization(cli, &store{}).GetOrganization(context.Background(), 12, nil)
			if tc.expectedError == "" {
				assert.Nil(t, err, "check api get error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestListMembers(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		iterator *memberListIterator
	}{
		{
			name: "create iterator successfully",
			iterator: &memberListIterator{
				iter: listIterator{
					resource: "member",
					path:     "/api/v1/orgs/123/members",
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
			// ignore the application check since currently we don't mock it, and the api is always a zero value.
			raw, err := newOrganization(nil, &store{}).ListMembers(context.Background(), &ResourceListOptions{
				Organization: &Organization{
					ID: 123,
				},
				Pagination: &Pagination{
					Page:     14,
					PageSize: 25,
				},
			})
			assert.Nil(t, err, "check list api error")
			iter := raw.(*memberListIterator)
			assert.Equal(t, tc.iterator.iter.resource, iter.iter.resource, "check resource")
			assert.Equal(t, tc.iterator.iter.path, iter.iter.path, "check path")
			assert.Equal(t, tc.iterator.iter.paging.Page, iter.iter.paging.Page, "check page")
			assert.Equal(t, tc.iterator.iter.paging.PageSize, iter.iter.paging.PageSize, "check page size")
		})
	}
}

func TestListRoles(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		iterator *roleListIterator
	}{
		{
			name: "create iterator successfully",
			iterator: &roleListIterator{
				iter: listIterator{
					resource: "role",
					path:     "/api/v1/orgs/123/roles",
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
			// ignore the application check since currently we don't mock it, and the api is always a zero value.
			raw, err := newOrganization(nil, &store{}).ListRoles(context.Background(), &ResourceListOptions{
				Organization: &Organization{
					ID: 123,
				},
				Pagination: &Pagination{
					Page:     14,
					PageSize: 25,
				},
			})
			assert.Nil(t, err, "check list api error")
			iter := raw.(*roleListIterator)
			assert.Equal(t, tc.iterator.iter.resource, iter.iter.resource, "check resource")
			assert.Equal(t, tc.iterator.iter.path, iter.iter.path, "check path")
			assert.Equal(t, tc.iterator.iter.paging.Page, iter.iter.paging.Page, "check page")
			assert.Equal(t, tc.iterator.iter.paging.PageSize, iter.iter.paging.PageSize, "check page size")
		})
	}
}

func TestInviteMember(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "update successfully",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newOrganization(cli, &store{}).InviteMember(context.Background(), "foo@test.org", &Role{}, &ResourceCreateOptions{
				Organization: &Organization{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check member invite error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestReInviteMember(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "update successfully",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members/1/re_invite"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members/1/re_invite"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newOrganization(cli, &store{}).ReInviteMember(context.Background(), 1, &ResourceUpdateOptions{
				Organization: &Organization{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check member re-invite error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestGetMember(t *testing.T) {
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
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members/12"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members/12"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli
			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			// ignore the API check since currently we don't mock it, and the app is always a zero value.
			_, err := newOrganization(cli, &store{}).GetMember(context.Background(), 12, &ResourceGetOptions{
				Organization: &Organization{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check member get error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestUpdateMemberRoles(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "update successfully",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members/12"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members/12"), "", gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			err := newOrganization(cli, &store{}).UpdateMemberRoles(context.Background(), 12, []RoleBinding{
				{
					RoleID: 1,
				},
			}, &ResourceUpdateOptions{
				Organization: &Organization{
					ID: 1,
				},
			})

			if tc.expectedError == "" {
				assert.Nil(t, err, "check roles update error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestDeleteMember(t *testing.T) {
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
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members/12"), "", nil, gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members/12"), "", nil, gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			err := newOrganization(cli, &store{}).RemoveMember(context.Background(), 12, &ResourceDeleteOptions{
				Organization: &Organization{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check member delete error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestTransferOwnership(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "transfer successfully",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members/12/transfer_ownership"), "", nil, gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/orgs/1/members/12/transfer_ownership"), "", nil, gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			err := newOrganization(cli, &store{}).TransferOwnership(context.Background(), 12, &ResourceUpdateOptions{
				Organization: &Organization{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check member transfer error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}
