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
	"testing"
)

func TestListIterator(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		mockFn        func(t *testing.T) *listIterator
		expectedError string
		getItems      []interface{}
	}{
		{
			name: "mock error",
			mockFn: func(t *testing.T) *listIterator {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), "/api/v1/controlplanes/1/apps", "page=1&page_size=10", gomock.Any()).Return(errors.New("mock error"))

				iter := &listIterator{
					ctx:      context.Background(),
					resource: "applications",
					client:   cli,
					path:     "/api/v1/controlplanes/1/apps",
					paging:   DefaultPagination,
				}

				return iter
			},
			expectedError: "mock error",
		},
		{
			name: "iterate gradually",
			mockFn: func(t *testing.T) *listIterator {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), "/api/v1/controlplanes/1/apps", "page=1&page_size=3", gomock.Any()).Return(nil)

				iter := &listIterator{
					ctx:      context.Background(),
					resource: "applications",
					client:   cli,
					path:     "/api/v1/controlplanes/1/apps",
					paging: Pagination{
						Page:     1,
						PageSize: 3,
					},
					items: []interface{}{
						&Application{
							ID: 1,
						},
						&Application{
							ID: 2,
						},
						&Application{
							ID: 3,
						},
					},
				}
				return iter
			},
			getItems: []interface{}{
				&Application{
					ID: 1,
				},
				&Application{
					ID: 2,
				},
				&Application{
					ID: 3,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var items []interface{}
			iter := tc.mockFn(t)
			for {
				item, err := iter.Next()
				if tc.expectedError != "" {
					assert.Contains(t, err.Error(), tc.expectedError, "check error for iterating next item")
					return
				} else {
					assert.Nil(t, err, "check if error is nil")
					if item == nil {
						break
					}
					items = append(items, item)
				}
			}
			assert.Len(t, items, len(tc.getItems), "check the number of items")
			for i, item := range items {
				assert.Equalf(t, tc.getItems[i].(*Application).ID, item.(*Application).ID, "check ID of item #%d", i)
			}
		})
	}
}
