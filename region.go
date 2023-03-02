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
	"encoding/json"
	"path"
	"time"
)

// Region is the specification of the deploy region for Cloud.
type Region struct {
	// ID is the unique identify to mark an object.
	ID ID `json:"id"`
	// Name is the object name.
	Name string `json:"name"`
	// CreatedAt is the object creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the last modified time of this object.
	UpdatedAt time.Time `json:"updated_at"`
	// Provider is the cloud vendor we use, like aws, gcp
	Provider string `json:"provider"`
	// Status is the region status
	Status EntityStatus `json:"status,omitempty"`
}

// RegionInterface is the interface for manipulating Region.
type RegionInterface interface {
	// ListRegions returns an iterator for listing Regions with the
	// given list conditions.
	// Users need to specify the Paging and Filter conditions (if necessary)
	// in the `opts`.
	ListRegions(ctx context.Context, opts *ResourceListOptions) (RegionListIterator, error)
}

// RegionListIterator is an iterator for listing Regions.
type RegionListIterator interface {
	// Next returns the next Region according to the filter conditions.
	Next() (*Region, error)
}

type regionImpl struct {
	client httpClient
	store  StoreInterface
}

type regionListIterator struct {
	iter listIterator
}

func newRegion(client httpClient, store StoreInterface) RegionInterface {
	return &regionImpl{
		client: client,
		store:  store,
	}
}

func (iter *regionListIterator) Next() (*Region, error) {
	var region Region
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	if err = json.Unmarshal(rawData, &region); err != nil {
		return nil, err
	}
	return &region, nil
}

func (impl *regionImpl) ListRegions(ctx context.Context, opts *ResourceListOptions) (RegionListIterator, error) {
	var (
		paging *Pagination
		filter *Filter
	)
	if opts != nil {
		paging = opts.Pagination
		filter = opts.Filter
	}

	iter := listIterator{
		ctx:      ctx,
		resource: "region",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "regions"),
		paging:   mergePagination(paging),
		filter:   filter,
		headers:  appendHeader(mapClusterIdFromStore(impl.store), mapClusterIdFromOpts(opts)),
	}

	return &regionListIterator{iter: iter}, nil
}
