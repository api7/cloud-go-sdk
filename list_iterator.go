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
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

var (
	// DefaultPagination is the default paging.
	DefaultPagination = Pagination{
		Page:     1,
		PageSize: 10,
	}
)

// Pagination indicates the paging.
type Pagination struct {
	// Page is the start page.
	Page int
	// PageSize is the page size (how many items will be in a page).
	PageSize int
}

type listIterator struct {
	ctx      context.Context
	resource string
	client   httpClient
	path     string
	paging   Pagination
	eof      bool
	items    []interface{}
}

func (iter *listIterator) Next() (interface{}, error) {
	if iter.eof {
		return nil, nil
	}

	if len(iter.items) == 0 {
		var (
			items []interface{}
		)
		query := make(url.Values)
		query.Set("page", strconv.Itoa(iter.paging.Page))
		query.Set("page_size", strconv.Itoa(iter.paging.PageSize))

		err := iter.client.sendGetRequest(iter.ctx, iter.path, query.Encode(), jsonPayloadDecodeFactory(&items))
		if err != nil {
			return nil, errors.Wrap(err, "list resources")
		}

		iter.items = items
	}

	if len(iter.items) > 0 {
		res := iter.items[0]
		iter.items[0] = nil
		iter.items = iter.items[1:]
		return res, nil
	} else {
		iter.eof = true
		return nil, nil
	}
}
