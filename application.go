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

// Application is the definition of API7 Cloud Application, which also contains
// some management fields.
type Application struct {
	ApplicationSpec `json:",inline"`

	// ID is the unique identify to mark an object.
	ID ID `json:"id"`
	// ControlPlaneID is id of control plane that current app belong with
	ControlPlaneID ID `json:"control_plane_id"`
	// Status is status of app
	Status EntityStatus `json:"status"`
	// CreatedAt is the object creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the last modified time of this object.
	UpdatedAt time.Time `json:"updated_at"`
	// AvailableCertIDs records the available cert ids for this app.
	AvailableCertIDs []ID `json:"available_cert_ids" gorm:"-"`
	// CanaryReleaseID is the canary release id that in progress
	CanaryReleaseID []ID `json:"canary_release_id" gorm:"-"`
	// CanaryUpstreamVersionList is the canary upstream version list that in progress or paused
	CanaryUpstreamVersionList []string `json:"canary_upstream_version_list" gorm:"-"`
}

// ApplicationSpec is the specification of the Application.
type ApplicationSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// Labels are used for resource classification and indexing
	Labels []string `json:"labels,omitempty"`
	// Protocols contains all the support protocols that this Application exposes.
	Protocols []string `json:"protocols,omitempty"`
	// The listening path prefix for this application
	PathPrefix string `json:"path_prefix"`
	// Hosts contains all the hosts that this Application uses.
	Hosts []string `json:"hosts"`
	// Plugins settings on Application level.
	Plugins Plugins `json:"plugins,omitempty"`
	// Upstream settings for the Application
	Upstreams []UpstreamAndVersion `json:"upstreams"`
	// DefaultUpstreamVersion settings for the upstream that should be used
	DefaultUpstreamVersion string `json:"default_upstream_version,omitempty"`
	// Active is status of application
	// Optional values can be:
	// * ActiveStatus: the object is active.
	// * InactiveStatus: the object is inactive.
	Active int `json:"active"`
}

// ApplicationInterface is the interface for manipulating Applications.
type ApplicationInterface interface {
	// CreateApplication creates an API7 Cloud Application in the specified control plane.
	// The given `app` parameter should specify the desired Application specification.
	// Users need to specify the ControlPlane in the `opts`.
	// The returned Application will contain the same Application specification plus some
	// management fields and default values.
	CreateApplication(ctx context.Context, app *Application, opts *ResourceCreateOptions) (*Application, error)
	// UpdateApplication updates an existing API7 Cloud Application in the specified control plane.
	// The given `app` parameter should specify the desired Application specification.
	// Users need to specify the ControlPlane in the `opts`.
	// The returned Application will contain the same Application specification plus some
	// management fields and default values.
	UpdateApplication(ctx context.Context, app *Application, opts *ResourceUpdateOptions) (*Application, error)
	// DeleteApplication deletes an existing API7 Cloud Application in the specified control plane.
	// The given `appID` parameter should specify the Application that you want to delete.
	// Users need to specify the ControlPlane in the `opts`.
	DeleteApplication(ctx context.Context, appID ID, opts *ResourceDeleteOptions) error
	// GetApplication gets an existing API7 Cloud Application in the specified control plane.
	// The given `appID` parameter should specify the Application that you want to get.
	// Users need to specify the ControlPlane in the `opts`.
	GetApplication(ctx context.Context, appID ID, opts *ResourceGetOptions) (*Application, error)
	// PublishApplication publishes the Application in the specified control plane (which is
	// a shortcut of UpdateApplication and set ApplicationSpec.Active to ActiveStatus).
	// The given `appID` parameter should specify the Application that you want to operate.
	// Users need to specify the ControlPlane in the `opts`.
	// The updated Application will be returned and the ApplicationSpec.Active field should be ActiveStatus.
	PublishApplication(ctx context.Context, appID ID, opts *ResourceUpdateOptions) (*Application, error)
	// UnpublishApplication publishes the Application in the specified control plane (which is
	// a shortcut of UpdateApplication and set ApplicationSpec.Active to InactiveStatus).
	// The given `appID` parameter should specify the Application that you want to operate.
	// Users need to specify the ControlPlane in the `opts`.
	// The updated Application will be returned and the ApplicationSpec.Active field should be InactiveStatus.
	UnpublishApplication(ctx context.Context, appID ID, opts *ResourceUpdateOptions) (*Application, error)
	// ListApplications returns an iterator for listing Applications in the specified control plane with the
	// given list conditions.
	// Users need to specify the ControlPlane, Paging and Filter conditions (if necessary)
	// in the `opts`.
	ListApplications(ctx context.Context, opts *ResourceListOptions) (ApplicationListIterator, error)
}

// ApplicationListIterator is an iterator for listing Applications.
type ApplicationListIterator interface {
	// Next returns the next Application according to the filter conditions.
	Next() (*Application, error)
}

type applicationImpl struct {
	client httpClient
}

type applicationListIterator struct {
	iter listIterator
}

func (iter *applicationListIterator) Next() (*Application, error) {
	var app Application
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	if err = json.Unmarshal(rawData, &app); err != nil {
		return nil, err
	}
	return &app, nil
}

func newApplication(cli httpClient) ApplicationInterface {
	return &applicationImpl{
		client: cli,
	}
}

func (impl *applicationImpl) CreateApplication(ctx context.Context, app *Application, opts *ResourceCreateOptions) (*Application, error) {
	var createdApp Application

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "apps")
	err := impl.client.sendPostRequest(ctx, uri, "", app, jsonPayloadDecodeFactory(&createdApp))
	if err != nil {
		return nil, err
	}
	return &createdApp, nil
}

func (impl *applicationImpl) UpdateApplication(ctx context.Context, app *Application, opts *ResourceUpdateOptions) (*Application, error) {
	var updatedApp Application

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "apps", app.ID.String())
	err := impl.client.sendPutRequest(ctx, uri, "", app, jsonPayloadDecodeFactory(&updatedApp))
	if err != nil {
		return nil, err
	}
	return &updatedApp, nil
}

func (impl *applicationImpl) DeleteApplication(ctx context.Context, appID ID, opts *ResourceDeleteOptions) error {
	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "apps", appID.String())
	return impl.client.sendDeleteRequest(ctx, uri, "", nil)
}

func (impl *applicationImpl) GetApplication(ctx context.Context, appID ID, opts *ResourceGetOptions) (*Application, error) {
	var app Application

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "apps", appID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&app))
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (impl *applicationImpl) PublishApplication(ctx context.Context, appID ID, opts *ResourceUpdateOptions) (*Application, error) {
	var app Application

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "apps", appID.String())
	body := []byte(`{"active":0}`)
	err := impl.client.sendPatchRequest(ctx, uri, "", body, jsonPayloadDecodeFactory(&app))
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (impl *applicationImpl) UnpublishApplication(ctx context.Context, appID ID, opts *ResourceUpdateOptions) (*Application, error) {
	var app Application

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "apps", appID.String())
	body := []byte(`{"active":1}`)
	err := impl.client.sendPatchRequest(ctx, uri, "", body, jsonPayloadDecodeFactory(&app))
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (impl *applicationImpl) ListApplications(ctx context.Context, opts *ResourceListOptions) (ApplicationListIterator, error) {
	iter := listIterator{
		ctx:      ctx,
		resource: "application",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "controlplanes", opts.ControlPlane.ID.String(), "apps"),
		paging:   mergePagination(opts.Pagination),
		filter:   opts.Filter,
	}

	return &applicationListIterator{iter: iter}, nil
}
