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

const (
	// RoleScopeOrganization indicates an organization scoped role.
	RoleScopeOrganization = "organization"
	// RoleScopeControlPlane indicates an control plane scoped role.
	RoleScopeControlPlane = "control_plane"

	// MemberStatePending means the member is still in pending state.
	MemberStatePending = "Pending"
	// MemberStateActive means the member is active.
	MemberStateActive = "Active"
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

// Member is the member of organization.
// It contains the member specification and some management fields.
type Member struct {
	MemberSpec `json:"-"`

	// ID is the unique identify to mark an object.
	ID ID `json:"id,omitempty,inline" yaml:"id"`
	// CreatedAt is the object creation time.
	CreatedAt time.Time `json:"created_at,omitempty" yaml:"created_at"`
	// UpdatedAt is the last modified time of this object.
	UpdatedAt time.Time `json:"updated_at,omitempty" yaml:"updated_at"`
	// OrgId indicates the organization where the member is in.
	OrgId ID `json:"org_id,omitempty,inline" yaml:"org_id"`
	// Status is the user data status.
	Status EntityStatus `json:"status,omitempty" yaml:"status"`
}

// MemberSpec contains the information
type MemberSpec struct {
	// FirstName is the member first name
	FirstName string `json:"first_name,omitempty" yaml:"first_name"`
	// LastName is the member last name
	LastName string `json:"last_name,omitempty" yaml:"last_name"`
	// Roles indicates the roles of the member.
	Roles []Role `json:"roles"`
	// Email is the email address of the member.
	Email string `json:"email"`
	// UserId refers to a user, since a 3rd party User Management
	// Service might be used so the type is not uint64.
	UserId string `json:"user_id,omitempty" yaml:"user_id"`
	// State is the user state. Optional values can be:
	// * MemberStatePending
	// * MemberStateActive
	State string `json:"state" yaml:"state"`
}

// Methods means the operations that can be performed on an organization.
type Methods struct {
	// Get is a get method
	Get bool `json:"get,omitempty"`
	// Put	is a put method
	Put bool `json:"put,omitempty"`
	// Post	is a post method
	Post bool `json:"post,omitempty"`
	// Patch is a patch method
	Patch bool `json:"patch,omitempty"`
	// Delete is a delete method
	Delete bool `json:"delete,omitempty"`
}

// Permissions means the permissions that can be performed on an organization.
type Permissions struct {
	// Organization is the organization scope of permission
	Organization map[string]Methods `json:"organization"`
	// ControlPlane is the control plane scope of permission
	ControlPlane map[string]Methods `json:"control_plane"`
	// Billing is the billing scope of permission
	Billing map[string]Methods `json:"billing"`
	// APIManagement is the API management scope of permission
	APIManagement map[string]Methods `json:"api_management"`
}

// Role is the role of a member.
type Role struct {
	// ID is the id of role
	ID ID `json:"id" gorm:"primaryKey"`
	// Name is the name of role
	Name string `json:"name" gorm:"column:name"`
	// OrgID is the id of organization
	OrgID ID `json:"org_id" gorm:"column:org_id"`
	// Owner is the owner of role
	Owner bool `json:"owner" gorm:"column:owner"`
	// Permissions is the permissions of role
	// Key means the name of resource
	// Value means the permissions of resource
	Permissions Permissions `json:"permissions" gorm:"serializer:json;"`
	// Scope is the scope of role. Optional values can be:
	// * RoleScopeOrganization
	// * RoleScopeControlPlane
	Scope     string    `json:"scope" gorm:"column:scope"`
	CreatedAt time.Time `json:"-" yaml:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" yaml:"updated_at" gorm:"autoUpdateTime"`
}

// RoleBinding binds a role to an organization member.
type RoleBinding struct {
	// RoleID is the id of role
	RoleID ID `json:"role_id"`
	// ControlPlaneID is the id of control plane
	// This field is used only if the role is not
	// organization scoped.
	ControlPlaneID ID `json:"control_plane_id"`
}

// OrganizationInterface is the interface for manipulating Organization and Member.
type OrganizationInterface interface {
	// GetOrganization gets an existing API7 Cloud Organization.
	// The given `orgID` parameter should specify the Organization that you want to get.
	// Currently, the `opts` parameter doesn't matter, users can pass the `nil` value.
	GetOrganization(ctx context.Context, orgID ID, opts *ResourceGetOptions) (*Organization, error)
	// ListMembers returns an iterator for listing Members in the specified Organization with the
	// given list conditions.
	// Users need to specify the Organization, Paging in the `opts`.
	ListMembers(ctx context.Context, opts *ResourceListOptions) (MemberListIterator, error)
	// InviteMember invites a new member to the organization.
	// The given `email` parameter should specify a correct mail address.
	// The given `role` parameter should specify an appropriate member role.
	// Users need to specify the Organization in the `opts`.
	InviteMember(ctx context.Context, email string, role *Role, opts *ResourceCreateOptions) (*Member, error)
	// ReInviteMember re-invites an existing member (which state is MemberStatePending) to the organization.
	// The given `memberID` parameter should specify the existing member.
	// Users need to specify the Organization in the `opts`.
	ReInviteMember(ctx context.Context, memberID ID, opts *ResourceUpdateOptions) (*Member, error)
	// RemoveMember removes an existing member from the organization.
	// The given `memberID` parameter should specify the existing member.
	// Users need to specify the Organization in the `opts`.
	RemoveMember(ctx context.Context, memberID ID, opts *ResourceDeleteOptions) error
	// GetMember gets an existing member from the organization.
	// The given `memberID` parameter should specify the existing member.
	// Users need to specify the Organization in the `opts`.
	GetMember(ctx context.Context, memberID ID, opts *ResourceGetOptions) (*Member, error)
	// UpdateMemberRoles updates the roles for the specified member.
	// The given `memberID` parameter should specify the existing member.
	// The given `roleBindings` parameter specifies new roles for this member.
	// Users need to specify the Organization in the `opts`.
	UpdateMemberRoles(ctx context.Context, memberID ID, roleBindings []RoleBinding, opts *ResourceUpdateOptions) error
	// ListRoles returns an iterator for listing Roles in the specified Organization with the
	// given list conditions.
	// Users need to specify the Organization, Paging in the `opts`.
	ListRoles(ctx context.Context, opts *ResourceListOptions) (RoleListIterator, error)
	// TransferOwnership transfers the organization ownership from yourself to another member.
	// The `toMember` parameter should specify the existing member in the same organization.
	// Users need to specify the Organization, Paging in the `opts`.
	// Note the operation will fail if you're not the owner of this organization.
	// After the transferring, your role will be downgraded to organization admin.
	TransferOwnership(ctx context.Context, toMember ID, opts *ResourceUpdateOptions) error
}

// MemberListIterator is an iterator for listing Members.
type MemberListIterator interface {
	// Next returns the next Member according to the filter conditions.
	Next() (*Member, error)
}

// RoleListIterator is an iterator for listing Roles.
type RoleListIterator interface {
	// Next returns the next Role according to the filter conditions.
	Next() (*Role, error)
}

type organizationImpl struct {
	client httpClient
}

type memberListIterator struct {
	iter listIterator
}

func (iter *memberListIterator) Next() (*Member, error) {
	var member Member
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	if err = json.Unmarshal(rawData, &member); err != nil {
		return nil, err
	}
	return &member, nil
}

type roleListIterator struct {
	iter listIterator
}

func (iter *roleListIterator) Next() (*Role, error) {
	var role Role
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	if err = json.Unmarshal(rawData, &role); err != nil {
		return nil, err
	}
	return &role, nil
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

func (impl *organizationImpl) ListMembers(ctx context.Context, opts *ResourceListOptions) (MemberListIterator, error) {
	iter := listIterator{
		ctx:      ctx,
		resource: "member",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "members"),
		paging:   mergePagination(opts.Pagination),
	}

	return &memberListIterator{
		iter: iter,
	}, nil
}

func (impl *organizationImpl) ListRoles(ctx context.Context, opts *ResourceListOptions) (RoleListIterator, error) {
	iter := listIterator{
		ctx:      ctx,
		resource: "role",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "roles"),
		paging:   mergePagination(opts.Pagination),
	}

	return &roleListIterator{
		iter: iter,
	}, nil
}

func (impl *organizationImpl) InviteMember(ctx context.Context, email string, role *Role, opts *ResourceCreateOptions) (*Member, error) {
	var member Member

	body := struct {
		Email  string `json:"email"`
		RoleID string `json:"role_id"`
	}{
		Email:  email,
		RoleID: role.ID.String(),
	}

	uri := path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "members")
	err := impl.client.sendPostRequest(ctx, uri, "", body, jsonPayloadDecodeFactory(&member))
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (impl *organizationImpl) ReInviteMember(ctx context.Context, memberID ID, opts *ResourceUpdateOptions) (*Member, error) {
	var member Member

	uri := path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "members", memberID.String(), "re_invite")
	err := impl.client.sendPutRequest(ctx, uri, "", nil, jsonPayloadDecodeFactory(&member))
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (impl *organizationImpl) RemoveMember(ctx context.Context, memberID ID, opts *ResourceDeleteOptions) error {
	uri := path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "members", memberID.String())
	err := impl.client.sendDeleteRequest(ctx, uri, "", nil)
	if err != nil {
		return err
	}
	return nil
}

func (impl *organizationImpl) GetMember(ctx context.Context, memberID ID, opts *ResourceGetOptions) (*Member, error) {
	var member Member

	uri := path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "members", memberID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&member))
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (impl *organizationImpl) TransferOwnership(ctx context.Context, targetMemberID ID, opts *ResourceUpdateOptions) error {
	uri := path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "members", targetMemberID.String(), "transfer_ownership")
	err := impl.client.sendPostRequest(ctx, uri, "", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (impl *organizationImpl) UpdateMemberRoles(ctx context.Context, memberID ID, roleBindings []RoleBinding, opts *ResourceUpdateOptions) error {
	uri := path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "members", memberID.String())
	err := impl.client.sendPutRequest(ctx, uri, "", roleBindings, nil)
	if err != nil {
		return err
	}
	return nil
}
