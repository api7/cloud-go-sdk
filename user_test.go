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

func TestUserMe(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		mockFunc      func(*testing.T) httpClient
		expectedUser  *User
		expectedError string
	}{
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/user/me"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli
			},
			expectedUser:  nil,
			expectedError: "/api/v1/user/me: mock error",
		},
		{
			name: "success",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/user/me"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli
			},
			expectedUser:  &User{},
			expectedError: "",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			user, err := newUser(tc.mockFunc(t), &store{}).Me(context.Background())
			if tc.expectedError != "" {
				assert.EqualError(t, err, tc.expectedError, "check error")
			} else {
				assert.Nil(t, err, "check if error is nil")
				assert.Equal(t, tc.expectedUser, user, "check user")
			}
		})
	}
}
