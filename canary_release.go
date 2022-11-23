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
	"fmt"
	"path"
	"time"
)

// CanaryRelease is the definition of API7 Cloud CanaryRelease.
type CanaryRelease struct {
	CanaryReleaseSpec `json:",inline"`
	// ID is the unique identify to mark an object.
	ID ID `json:"id"`
	// AppID is id of current application
	AppID ID `json:"app_id"`
	// Status is status of canary release
	Status EntityStatus `json:"status"`
	// CreatedAt is the object creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the last modified time of this object.
	UpdatedAt time.Time `json:"updated_at"`
}

// CanaryReleaseRule is the rule that used in canary release
type CanaryReleaseRule struct {
	// Position means where we should get the key.Can be "header", "query" etc.
	Position string `json:"position"`
	// Key means the name of the key
	Key string `json:"key"`
	// Operator means the operator that used between key and value
	Operator string `json:"operator"`
	// Value means the values that used in the expression.
	Value interface{} `json:"value,omitempty"`
}

// CanaryReleaseSpec is the specification of the CanaryRelease
type CanaryReleaseSpec struct {
	// Name specify the name of canary release
	Name string `json:"name"`
	// State specify the state of the canary release.
	// Optional values can be:
	// * CanaryReleaseStatePause:  the object state is pause.
	// * CanaryReleaseStateInProgress:  the object state is in_progress.
	// * CanaryReleaseStateInFinish:  the object state is finish.
	State string `json:"state"`
	// Type specify the type of canary release.Can be "percent", "rules"
	Type string `json:"type"`
	// CanaryUpstreamVersion specify the version of canary release
	CanaryUpstreamVersion string `json:"canary_upstream_version"`
	// Percent specifies the percent of request will be transferred to canary upstream. Can be 0-100
	Percent int `json:"percent,omitempty"`
	// CanaryReleaseRule specify the matched rules of request that should be transferred to canary upstream
	CanaryReleaseRule []CanaryReleaseRule `json:"rules,omitempty"`
}

// CanaryReleaseInterface is the interface for manu
type CanaryReleaseInterface interface {
	// CreateCanaryRelease creates an API7 Cloud Canary Release in the specified Application.
	// The given `cr` parameter should specify the desired Canary Release specification.
	// Users need to specify the Application in the `opts`.
	// The returned CanaryRelease will contain the same CanaryRelease specification plus some
	// management fields and default values
	CreateCanaryRelease(ctx context.Context, cr *CanaryRelease, opts *ResourceCreateOptions) (*CanaryRelease, error)

	// StartCanaryRelease makes the Canary Release in progress in the specified Application
	// (a shortcut of UpdateCanaryRelease and set CanaryReleaseSpec.State to CanaryReleaseStateInProgress).
	// The given `crID` parameter should specify the desired Canary Release ID.
	// Users need to specify the Application in the `opts`.
	// The updated Canary Release will be returned and the CanaryReleaseSpec.State field should be
	// CanaryReleaseStateInProgress.
	StartCanaryRelease(ctx context.Context, crID ID, opts *ResourceUpdateOptions) (*CanaryRelease, error)
	// PauseCanaryRelease makes the Canary Release paused in the specified Application
	// (a shortcut of UpdateCanaryRelease and set CanaryReleaseSpec.State to CanaryReleaseStatePaused).
	// The given `crID` parameter should specify the desired Canary Release ID.
	// Users need to specify the Application in the `opts`.
	// The updated Canary Release will be returned and the CanaryReleaseSpec.State field should be
	// CanaryReleaseStatePaused.
	PauseCanaryRelease(ctx context.Context, crID ID, opts *ResourceUpdateOptions) (*CanaryRelease, error)
	// FinishCanaryRelease makes the Canary Release finished in the specified Application
	// (a shortcut of UpdateCanaryRelease and set CanaryReleaseSpec.State to CanaryReleaseStateFinished).
	// The given `crID` parameter should specify the desired Canary Release ID.
	// Users need to specify the Application in the `opts`.
	// The updated Canary Release will be returned and the CanaryReleaseSpec.State field should be
	// CanaryReleaseStateFinished.
	FinishCanaryRelease(ctx context.Context, crID ID, opts *ResourceUpdateOptions) (*CanaryRelease, error)
}

type canaryReleaseImpl struct {
	client httpClient
}

func newCanaryRelease(cli httpClient) CanaryReleaseInterface {
	return &canaryReleaseImpl{
		client: cli,
	}
}

func (impl *canaryReleaseImpl) CreateCanaryRelease(ctx context.Context, cr *CanaryRelease, opts *ResourceCreateOptions) (*CanaryRelease, error) {
	var createCr CanaryRelease

	appID := opts.Application.ID
	uri := path.Join(_apiPathPrefix, "apps", appID.String(), "canary_releases")
	err := impl.client.sendPostRequest(ctx, uri, "", cr, jsonPayloadDecodeFactory(&createCr))
	if err != nil {
		return nil, err
	}
	return &createCr, nil
}

func (impl *canaryReleaseImpl) StartCanaryRelease(ctx context.Context, crID ID, opts *ResourceUpdateOptions) (*CanaryRelease, error) {
	var cr CanaryRelease

	appID := opts.Application.ID
	uri := path.Join(_apiPathPrefix, "apps", appID.String(), "canary_releases", crID.String())
	body := []byte(fmt.Sprintf(`{"state":"%s"}`, CanaryReleaseStateInProgress))
	err := impl.client.sendPatchRequest(ctx, uri, "", body, jsonPayloadDecodeFactory(&cr))
	if err != nil {
		return nil, err
	}
	return &cr, nil
}

func (impl *canaryReleaseImpl) PauseCanaryRelease(ctx context.Context, crID ID, opts *ResourceUpdateOptions) (*CanaryRelease, error) {
	var cr CanaryRelease

	appID := opts.Application.ID
	uri := path.Join(_apiPathPrefix, "apps", appID.String(), "canary_releases", crID.String())
	body := []byte(fmt.Sprintf(`{"state":"%s"}`, CanaryReleaseStatePaused))
	err := impl.client.sendPatchRequest(ctx, uri, "", body, jsonPayloadDecodeFactory(&cr))
	if err != nil {
		return nil, err
	}
	return &cr, nil
}

func (impl *canaryReleaseImpl) FinishCanaryRelease(ctx context.Context, crID ID, opts *ResourceUpdateOptions) (*CanaryRelease, error) {
	var cr CanaryRelease

	appID := opts.Application.ID
	uri := path.Join(_apiPathPrefix, "apps", appID.String(), "canary_releases", crID.String())
	body := []byte(fmt.Sprintf(`{"state":"%s"}`, CanaryReleaseStateFinished))
	err := impl.client.sendPatchRequest(ctx, uri, "", body, jsonPayloadDecodeFactory(&cr))
	if err != nil {
		return nil, err
	}
	return &cr, nil
}
