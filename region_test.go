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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRegions(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		iterator *regionListIterator
	}{
		{
			name: "create iterator successfully",
			iterator: &regionListIterator{
				iter: listIterator{
					resource: "region",
					path:     "/api/v1/regions",
					paging: Pagination{
						Page:     1,
						PageSize: 10,
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
			raw, err := newRegion(nil, &store{}).ListRegions(context.Background(), nil)
			assert.Nil(t, err, "check list cluster error")
			iter := raw.(*regionListIterator)
			assert.Equal(t, tc.iterator.iter.resource, iter.iter.resource, "check resource")
			assert.Equal(t, tc.iterator.iter.path, iter.iter.path, "check path")
			assert.Equal(t, tc.iterator.iter.paging.Page, iter.iter.paging.Page, "check page")
			assert.Equal(t, tc.iterator.iter.paging.PageSize, iter.iter.paging.PageSize, "check page size")
		})
	}
}
