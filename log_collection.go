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
)

// LogCollectionType is the type of log collection
type LogCollectionType string

var (
	// HTTPLogCollection means http log collection
	HTTPLogCollection LogCollectionType = "http-logger"
	// KakfaLogCollection means kafka log collection
	KakfaLogCollection LogCollectionType = "kafka-logger"
)

// LogCollection is the abstraction of log storage
type LogCollection struct {
	// ID is the unique identify to mark an object.
	ID ID `json:"id"`
	// Name is the name of log collection
	Name string `json:"name"`
	// Description is the description of log collection
	Description string `json:"description"`
	// Type is the type of log collection
	Type LogCollectionType `json:"type"`
	// Spec is the specification of log collection
	Spec interface{} `json:"spec" gorm:"serializer:json"`
}

// LogCollectionInterface is the interface of the LogCollection
type LogCollectionInterface interface {
	// CreateLogCollection creates an API7 Cloud Log Collection in the specified control plane.
	// The given `lc` parameter should specify the desired LogCollection specification.
	// Users need to specify the Cluster in the `opts`.
	// The returned LogCollection will contain the same LogCollection specification plus some
	// management fields and default values.
	CreateLogCollection(ctx context.Context, lc *LogCollection, opts *ResourceCreateOptions) (*LogCollection, error)
	// UpdateLogCollection updates an existing API7 Cloud Log Collection in the specified control plane.
	// The given `lc` parameter should specify the desired LogCollection specification.
	// Users need to specify the Cluster in the `opts`.
	// The returned LogCollection will contain the same LogCollection specification plus some
	// management fields and default values.
	UpdateLogCollection(ctx context.Context, lc *LogCollection, opts *ResourceUpdateOptions) (*LogCollection, error)
	// DeleteLogCollection deletes an existing API7 Cloud Log Collection in the specified control plane.
	// The given `lcID` parameter should specify the LogCollection that you want to delete.
	// Users need to specify the Cluster in the `opts`.
	DeleteLogCollection(ctx context.Context, lcID ID, opts *ResourceDeleteOptions) error
	// GetLogCollection gets an existing API7 Cloud Log Collection in the specified control plane.
	// The given `lcID` parameter should specify the LogCollection that you want to get.
	// Users need to specify the Cluster in the `opts`.
	GetLogCollection(ctx context.Context, lcID ID, opts *ResourceGetOptions) (*LogCollection, error)
	// ListLogCollections returns an iterator for listing Log Collections in the specified control plane with the
	// given list conditions.
	// Users need to specify the Cluster, Paging, Filter conditions (if necessary)
	// in the `opts`.
	ListLogCollections(ctx context.Context, opts *ResourceListOptions) (LogCollectionIterator, error)
}

// LogCollectionIterator is an iterator for listing Log Collections.
type LogCollectionIterator interface {
	// Next returns the next Log Collection according to the filter conditions.
	Next() (*LogCollection, error)
}

type logCollectionIterator struct {
	iter listIterator
}

type logCollectionImpl struct {
	client httpClient
}

func (iter *logCollectionIterator) Next() (*LogCollection, error) {
	var lc LogCollection
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	if err = json.Unmarshal(rawData, &lc); err != nil {
		return nil, err
	}
	return &lc, nil
}

func newLogCollection(cli httpClient) LogCollectionInterface {
	return &logCollectionImpl{
		client: cli,
	}
}

func (impl *logCollectionImpl) CreateLogCollection(ctx context.Context, lc *LogCollection, opts *ResourceCreateOptions) (*LogCollection, error) {
	var createdLogCollection LogCollection

	cpID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", cpID.String(), "log_collections")
	err := impl.client.sendPostRequest(ctx, uri, "", lc, jsonPayloadDecodeFactory(&createdLogCollection))
	if err != nil {
		return nil, err
	}
	return &createdLogCollection, nil
}

func (impl *logCollectionImpl) UpdateLogCollection(ctx context.Context, lc *LogCollection, opts *ResourceUpdateOptions) (*LogCollection, error) {
	var createdLogCollection LogCollection

	cpID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", cpID.String(), "log_collections", lc.ID.String())
	err := impl.client.sendPutRequest(ctx, uri, "", lc, jsonPayloadDecodeFactory(&createdLogCollection))
	if err != nil {
		return nil, err
	}
	return &createdLogCollection, nil
}

func (impl *logCollectionImpl) DeleteLogCollection(ctx context.Context, lcID ID, opts *ResourceDeleteOptions) error {
	cpID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", cpID.String(), "log_collections", lcID.String())
	return impl.client.sendDeleteRequest(ctx, uri, "", nil)
}

func (impl *logCollectionImpl) GetLogCollection(ctx context.Context, lcID ID, opts *ResourceGetOptions) (*LogCollection, error) {
	var logcollection LogCollection

	cpID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", cpID.String(), "log_collections", lcID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&logcollection))
	if err != nil {
		return nil, err
	}
	return &logcollection, nil
}

func (impl *logCollectionImpl) ListLogCollections(ctx context.Context, opts *ResourceListOptions) (LogCollectionIterator, error) {
	iter := listIterator{
		ctx:      ctx,
		resource: "logcollection",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "clusters", opts.Cluster.ID.String(), "log_collections"),
		paging:   mergePagination(opts.Pagination),
		filter:   opts.Filter,
	}

	return &logCollectionIterator{iter: iter}, nil
}
