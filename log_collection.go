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
	CreateLogCollection(ctx context.Context, lc *LogCollection, opts *ResourceCreateOptions) (*LogCollection, error)
	UpdateLogCollection(ctx context.Context, lc *LogCollection, opts *ResourceUpdateOptions) (*LogCollection, error)
	DeleteLogCollection(ctx context.Context, lcID ID, opts *ResourceDeleteOptions) error
	GetLogCollection(ctx context.Context, lcID ID, opts *ResourceGetOptions) (*LogCollection, error)
	ListLogCollection(ctx context.Context, opts *ResourceListOptions) (LogCollectionIterator, error)
}

type LogCollectionIterator interface {
	Next() (*LogCollection, error)
}

type logCollectionIterator struct {
	iter listIterator
}

type logCollectionImpl struct {
	client httpClient
}

func (iter *logCollectionIterator) Next() (*LogCollection, error) {
	app, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if app == nil {
		return nil, nil
	}
	return app.(*LogCollection), nil
}

func newLogCollection(cli httpClient) LogCollectionInterface {
	return &logCollectionImpl{
		client: cli,
	}
}

func (impl *logCollectionImpl) CreateLogCollection(ctx context.Context, lc *LogCollection, opts *ResourceCreateOptions) (*LogCollection, error) {
	var createdLogCollection LogCollection

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "log_collections")
	err := impl.client.sendPostRequest(ctx, uri, "", lc, jsonPayloadDecodeFactory(createdLogCollection))
	if err != nil {
		return nil, err
	}
	return &createdLogCollection, nil
}

func (impl *logCollectionImpl) UpdateLogCollection(ctx context.Context, lc *LogCollection, opts *ResourceUpdateOptions) (*LogCollection, error) {
	var createdLogCollection LogCollection

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "log_collections", lc.ID.String())
	err := impl.client.sendPutRequest(ctx, uri, "", lc, jsonPayloadDecodeFactory(createdLogCollection))
	if err != nil {
		return nil, err
	}
	return &createdLogCollection, nil
}

func (impl *logCollectionImpl) DeleteLogCollection(ctx context.Context, lcID ID, opts *ResourceDeleteOptions) error {
	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "log_collections", lcID.String())
	return impl.client.sendDeleteRequest(ctx, uri, "", nil)
}

func (impl *logCollectionImpl) GetLogCollection(ctx context.Context, lcID ID, opts *ResourceGetOptions) (*LogCollection, error) {
	var logcollection LogCollection

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "log_collections", lcID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&logcollection))
	if err != nil {
		return nil, err
	}
	return &logcollection, nil
}

func (impl *logCollectionImpl) ListLogCollection(ctx context.Context, opts *ResourceListOptions) (LogCollectionIterator, error) {
	iter := listIterator{
		ctx:      ctx,
		resource: "logcollection",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "controlplanes", opts.ControlPlane.ID.String(), "log_collections"),
		paging:   mergePagination(opts.Pagination),
	}

	return &logCollectionIterator{iter: iter}, nil
}
