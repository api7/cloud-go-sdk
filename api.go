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

// API is the definition of API7 Cloud API, which also contains
// some management fields.
//
//	API is an affiliated resource of Application.
type API struct {
	APISpec `json:",inline" gorm:"column:spec"`

	// ID is the unique identify to mark an object.
	ID ID `json:"id"`
	// AppID is id of app that current api belong with
	AppID ID `json:"app_id"`
	// Status is status of api
	Status EntityStatus `json:"status"`
	// CreatedAt is the object creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the last modified time of this object.
	UpdatedAt time.Time `json:"updated_at"`
}

// APISpec is the specification of the API.
type APISpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// Labels are used for resource classification and indexing
	Labels []string `json:"labels,omitempty"`
	// Methods are allowed HTTP methods to access this API, if absent, all HTTP methods are allowed
	Methods []string `json:"methods"`
	// Paths indicates which URI paths can be matched (prefix or exact) to this API.
	Paths []APIPath `json:"paths"`
	// FineGrainedRouteControl is used to control the route matching.
	FineGrainedRouteControl *FineGrainedRouteControl `json:"fine_grained_route_control,omitempty"`
	// StripPathPrefix indicates whether to strip the path prefix (defined in the Application)
	// before the gateway forwards API requests to upstream.
	StripPathPrefix bool `json:"strip_path_prefix"`
	// Plugins settings on API level, it'll override the same one on Application level
	// (instead of running them twice).
	Plugins Plugins `json:"plugins,omitempty"`
	// Active is the status of the API
	// Optional values can be:
	// * ActiveStatus: the object is active.
	// * InactiveStatus: the object is inactive.
	Active int `json:"active"`
	// Type is the type of the API
	// Optional values can be:
	// * APITypeRest means this is an RESTful API.
	// * APITypeWebSocket means this is an WebSocket API.
	Type string `json:"type,omitempty"`
}

// APIPath is the path definition for an API.
type APIPath struct {
	// Path is the URL path (after the Application path prefix) that the API will listen,
	// when Path is empty, the whole path is equal to the Application path prefix.
	Path string `json:"path"`
	// PathType determines the match type, by default it is prefix match.
	// Optional values can be:
	// * PathPrefixMatch means the requests' URL path leads with the API path will match this API;
	// * PathExactMatch means the requests' URL path has to be same to the API path.
	PathType string `json:"path_type"`
}

// FineGrainedRouteControl is fine grained route control settings.
type FineGrainedRouteControl struct {
	// Enabled indicates whether to enable fine-grained route control.
	Enabled bool `json:"enabled,omitempty"`
	// LogicalRelationship indicates the logical relationship between expressions.
	LogicalRelationship ExpressionLogicalRelationship `json:"logical_relationship,omitempty"`
	// Expressions is a list of expressions.
	Expressions []Expression `json:"expressions,omitempty"`
}

// APIInterface is the interface for manipulating API.
type APIInterface interface {
	// CreateAPI creates an API7 Cloud API in the specified Application.
	// The given `api` parameter should specify the desired API specification.
	// Users need to specify the Application in the `opts`.
	// The returned APi will contain the same API specification plus some
	// management fields and default values.
	CreateAPI(ctx context.Context, api *API, opts *ResourceCreateOptions) (*API, error)
	// UpdateAPI updates an existing API7 Cloud API in the specified Application.
	// The given `api` parameter should specify the desired API specification.
	// Users need to specify the Application in the `opts`.
	// The returned API will contain the same API specification plus some
	// management fields and default values.
	UpdateAPI(ctx context.Context, api *API, opts *ResourceUpdateOptions) (*API, error)
	// DeleteAPI deletes an existing API7 Cloud API in the specified Application.
	// The given `apiID` parameter should specify the API that you want to delete.
	// Users need to specify the Application in the `opts`.
	DeleteAPI(ctx context.Context, apiID ID, opts *ResourceDeleteOptions) error
	// GetAPI gets an existing API7 Cloud API in the specified Application.
	// The given `apiID` parameter should specify the API that you want to get.
	// Users need to specify the Application in the `opts`.
	GetAPI(ctx context.Context, apiID ID, opts *ResourceGetOptions) (*API, error)
	// PublishAPI publishes the APi in the specified Application (which is
	// a shortcut of UpdateAPI and set APISpec.Active to ActiveStatus).
	// The given `apiID` parameter should specify the API that you want to operate.
	// Users need to specify the Application in the `opts`.
	// The updated API will be returned and the APISpec.Active field should be ActiveStatus.
	PublishAPI(ctx context.Context, apiID ID, opts *ResourceUpdateOptions) (*API, error)
	// UnpublishAPI publishes the API in the specified Application (which is
	// a shortcut of UpdateAPI and set APISpec.Active to InactiveStatus).
	// The given `apiID` parameter should specify the API that you want to operate.
	// Users need to specify the Application in the `opts`.
	// The updated APi will be returned and the APISpec.Active field should be InactiveStatus.
	UnpublishAPI(ctx context.Context, apiID ID, opts *ResourceUpdateOptions) (*API, error)
	// ListAPIs returns an iterator for listing APIs in the specified Application with the
	// given list conditions.
	// Users need to specify the Application, Paging and Filter conditions (if necessary)
	// in the `opts`.
	ListAPIs(ctx context.Context, opts *ResourceListOptions) (APIListIterator, error)
	// DebugAPIResources returns the corresponding translated APISIX resources for this API.
	// The given `apiID` parameter should specify the API that you want to operate.
	// Users need to specify the Cluster.ID in the `opts`.
	DebugAPIResources(ctx context.Context, apiID ID, opts *ResourceGetOptions) (string, error)
}

// APIListIterator is an iterator for listing APIs.
type APIListIterator interface {
	// Next returns the next API according to the filter conditions.
	Next() (*API, error)
}

type apiImpl struct {
	client httpClient
	store  StoreInterface
}

type apiListIterator struct {
	iter listIterator
}

func (iter *apiListIterator) Next() (*API, error) {
	var api API
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	if err = json.Unmarshal(rawData, &api); err != nil {
		return nil, err
	}
	return &api, nil
}

func newAPI(cli httpClient, store StoreInterface) APIInterface {
	return &apiImpl{
		client: cli,
		store:  store,
	}
}

func (impl *apiImpl) CreateAPI(ctx context.Context, api *API, opts *ResourceCreateOptions) (*API, error) {
	var createdAPI API
	if !ensureClusterID(impl.store, opts) {
		return nil, ErrClusterIDNotExist
	}
	appID := opts.Application.ID
	uri := path.Join(_apiPathPrefix, "apps", appID.String(), "apis")
	err := impl.client.sendPostRequest(ctx, uri, "", api, jsonPayloadDecodeFactory(&createdAPI), appendHeader(mapClusterIdFromStore(impl.store), mapClusterIdFromOpts(opts)))
	if err != nil {
		return nil, err
	}
	return &createdAPI, nil
}

func (impl *apiImpl) UpdateAPI(ctx context.Context, api *API, opts *ResourceUpdateOptions) (*API, error) {
	var updatedAPI API
	if !ensureClusterID(impl.store, opts) {
		return nil, ErrClusterIDNotExist
	}
	appID := opts.Application.ID
	uri := path.Join(_apiPathPrefix, "apps", appID.String(), "apis", api.ID.String())
	err := impl.client.sendPutRequest(ctx, uri, "", api, jsonPayloadDecodeFactory(&updatedAPI), appendHeader(mapClusterIdFromStore(impl.store), mapClusterIdFromOpts(opts)))
	if err != nil {
		return nil, err
	}
	return &updatedAPI, nil
}

func (impl *apiImpl) DeleteAPI(ctx context.Context, apiID ID, opts *ResourceDeleteOptions) error {
	appID := opts.Application.ID
	if !ensureClusterID(impl.store, opts) {
		return ErrClusterIDNotExist
	}
	uri := path.Join(_apiPathPrefix, "apps", appID.String(), "apis", apiID.String())
	return impl.client.sendDeleteRequest(ctx, uri, "", nil, appendHeader(mapClusterIdFromStore(impl.store), mapClusterIdFromOpts(opts)))
}

func (impl *apiImpl) GetAPI(ctx context.Context, apiID ID, opts *ResourceGetOptions) (*API, error) {
	var api API
	if !ensureClusterID(impl.store, opts) {
		return nil, ErrClusterIDNotExist
	}
	appID := opts.Application.ID
	uri := path.Join(_apiPathPrefix, "apps", appID.String(), "apis", apiID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&api), appendHeader(mapClusterIdFromStore(impl.store), mapClusterIdFromOpts(opts)))
	if err != nil {
		return nil, err
	}
	return &api, nil
}

func (impl *apiImpl) PublishAPI(ctx context.Context, apiID ID, opts *ResourceUpdateOptions) (*API, error) {
	var api API
	if !ensureClusterID(impl.store, opts) {
		return nil, ErrClusterIDNotExist
	}
	appID := opts.Application.ID
	uri := path.Join(_apiPathPrefix, "apps", appID.String(), "apis", apiID.String())
	body := []byte(`{"active":0}`)
	err := impl.client.sendPatchRequest(ctx, uri, "", body, jsonPayloadDecodeFactory(&api), appendHeader(mapClusterIdFromStore(impl.store), mapClusterIdFromOpts(opts)))
	if err != nil {
		return nil, err
	}
	return &api, nil
}

func (impl *apiImpl) UnpublishAPI(ctx context.Context, apiID ID, opts *ResourceUpdateOptions) (*API, error) {
	var api API
	if !ensureClusterID(impl.store, opts) {
		return nil, ErrClusterIDNotExist
	}
	appID := opts.Application.ID
	uri := path.Join(_apiPathPrefix, "apps", appID.String(), "apis", apiID.String())
	body := []byte(`{"active":1}`)
	err := impl.client.sendPatchRequest(ctx, uri, "", body, jsonPayloadDecodeFactory(&api), appendHeader(mapClusterIdFromStore(impl.store), mapClusterIdFromOpts(opts)))
	if err != nil {
		return nil, err
	}
	return &api, nil
}

func (impl *apiImpl) ListAPIs(ctx context.Context, opts *ResourceListOptions) (APIListIterator, error) {
	if !ensureClusterID(impl.store, opts) {
		return nil, ErrClusterIDNotExist
	}
	iter := listIterator{
		ctx:      ctx,
		resource: "api",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "apps", opts.Application.ID.String(), "apis"),
		paging:   mergePagination(opts.Pagination),
		filter:   opts.Filter,
		headers:  appendHeader(mapClusterIdFromStore(impl.store), mapClusterIdFromOpts(opts)),
	}

	return &apiListIterator{iter: iter}, nil
}

func (impl *apiImpl) DebugAPIResources(ctx context.Context, apiID ID, opts *ResourceGetOptions) (string, error) {
	var rawData json.RawMessage
	if !ensureClusterID(impl.store, opts) {
		return "", ErrClusterIDNotExist
	}
	uri := path.Join(_apiPathPrefix, "debug", "config", "clusters", opts.Cluster.ID.String(), "api", apiID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&rawData), appendHeader(mapClusterIdFromStore(impl.store), mapClusterIdFromOpts(opts)))
	if err != nil {
		return "", err
	}
	return formatJSONData(rawData)
}
