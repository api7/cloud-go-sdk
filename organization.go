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
	"time"
)

// Organization is the specification of an API7 Cloud organization.
type Organization struct {
	// ID is the unique identify to mark an object.
	ID ID `json:"id,inline" yaml:"id"`
	// Name is the object name.
	Name string `json:"name" yaml:"name"`
	// CreatedAt is the object creation time.
	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
	// UpdatedAt is the last modified time of this object.
	UpdatedAt time.Time `json:"updated_at" yaml:"updated_at"`
	// PlanID indicates which plan is used by this organization.
	// PlanID should refer to a valid Plan object.
	PlanID ID `json:"plan_id" yaml:"plan_id"`
	// PlanExpireTime indicates the binding plan expire time for this organization.
	PlanExpireTime time.Time `json:"plan_expire_time" yaml:"plan_expire_time"`
	// SubscriptionStartedAt is the time when the organization subscribed to the plan.
	SubscriptionStartedAt *time.Time `json:"subscription_started_at" yaml:"subscription_started_at"`
	// OwnerID indicates who create the organization.
	OwnerID string `json:"owner_id" yaml:"owner_id"`
}

// OrganizationInterface is the interface for manipulating Organization.
type OrganizationInterface interface {
	// GetOrganization gets an existing API7 Cloud Organization.
	// The given `orgID` parameter should specify the Organization that you want to get.
	// Currently, the `opts` parameter doesn't matter, users can pass the `nil` value.
	GetOrganization(ctx context.Context, orgID ID, opts *ResourceGetOptions) (*Organization, error)
}

type organizationImpl struct {
	client httpClient
}

func newOrganization(cli httpClient) OrganizationInterface {
	return &organizationImpl{
		client: cli,
	}
}

func (impl *organizationImpl) GetOrganization(ctx context.Context, orgID ID, _ *ResourceGetOptions) (*Organization, error) {
	var org Organization

	uri := path.Join(_apiPathPrefix, "orgs", orgID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&org))
	if err != nil {
		return nil, err
	}
	return &org, nil
}
