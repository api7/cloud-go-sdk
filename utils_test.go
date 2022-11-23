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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergePagination(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		paging         *Pagination
		expectedPaging *Pagination
	}{
		{
			name:           "paging is nil",
			expectedPaging: &DefaultPagination,
		},
		{
			name: "page size is not specified",
			paging: &Pagination{
				Page: 3,
			},
			expectedPaging: &Pagination{
				Page:     3,
				PageSize: 10,
			},
		},
		{
			name: "page is not specified",
			paging: &Pagination{
				PageSize: 5,
			},
			expectedPaging: &Pagination{
				Page:     1,
				PageSize: 5,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			paging := mergePagination(tc.paging)
			assert.Equal(t, tc.expectedPaging.Page, paging.Page, "check page")
			assert.Equal(t, tc.expectedPaging.PageSize, paging.PageSize, "check page size")
		})
	}
}
