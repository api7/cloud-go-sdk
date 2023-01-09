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
	"encoding/json"
	"errors"
	"strconv"

	"github.com/sony/sonyflake"
)

// ID is the type of the id field used for any entities
type ID uint64

// String indicates how to convert ID to a string.
func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// MarshalJSON is the way to encode ID to JSON string.
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(id), 10))
}

// UnmarshalJSON is the way to decode ID from JSON string.
func (id *ID) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch v := value.(type) {
	case string:
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		*id = ID(u)
	default:
		panic("unknown type")
	}
	return nil
}

// IDGenerator is an interface for generating IDs.
type IDGenerator interface {
	// NextID generates an ID.
	NextID() ID
}

type snowflake sonyflake.Sonyflake

func (s *snowflake) NextID() ID {
	uid, err := (*sonyflake.Sonyflake)(s).NextID()
	if err != nil {
		panic("get sony flake uid failed:" + err.Error())
	}
	return ID(uid)
}

// NewIDGenerator returns an IDGenerator object.
func NewIDGenerator() (IDGenerator, error) {
	ips, err := getLocalIPs()
	if err != nil {
		panic(err)
	}
	sf := (*snowflake)(sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (u uint16, e error) {
			return sumIPs(ips), nil
		},
	}))
	if sf == nil {
		return nil, errors.New("failed to new snoyflake object")
	}
	return sf, nil
}

const (
	// Any means any status
	Any = EntityStatus(-1)
	// Uninitialized represents the entity has been saved to the db, but the associated resource has not yet been ready
	Uninitialized = EntityStatus(0)
	// Normal indicates that the entity and associated resources are ready
	Normal = EntityStatus(50)
	// Deleted indicates the entity has been deleted
	Deleted = EntityStatus(100)
)

const (
	// APITypeRest indicates this is a Rest API.
	APITypeRest = "Rest"
	// APITypeWebSocket indicates this is a Websocket API.
	APITypeWebSocket = "WebSocket"

	// CanaryReleaseTypePercent means using percent to do canary release
	CanaryReleaseTypePercent = "percent"
	// CanaryReleaseTypeRule means using rule match to do canary release
	CanaryReleaseTypeRule = "rule"
)

const (
	// PathPrefixMatch means the requests' URL path leads with the API path will match this API;
	PathPrefixMatch = "Prefix"
	// PathExactMatch means the requests' URL path has to be same to the API path.
	PathExactMatch = "Exact"
)

const (
	// ProtocolHTTP indicates the HTTP protocol.
	ProtocolHTTP = "HTTP"
	// ProtocolHTTPS indicates the HTTPS protocol.
	ProtocolHTTPS = "HTTPS"
)

const (
	// ActiveStatus indicates an object is active, and this object
	// will be seen by gateway instances.
	ActiveStatus = iota
	// InactiveStatus indicates an object is inactive, and this object
	// won't be seen by gateway instances.
	InactiveStatus
)

const (
	// CanaryReleaseStatePaused indicates the pause state of CanaryRelease.
	CanaryReleaseStatePaused = "paused"
	// CanaryReleaseStateInProgress indicates the in_progress state of CanaryRelease.
	CanaryReleaseStateInProgress = "in_progress"
	// CanaryReleaseStateFinished indicates the finish state of CanaryRelease.
	CanaryReleaseStateFinished = "finished"
)

// EntityStatus is common status definition for any kind of entity:
// * Uninitialized represents the entity has been saved to the db, but the associated resource has not yet been ready.
// * Normal indicates that the entity and associated resources are ready.
// * Deleted indicates the entity has been deleted.
type EntityStatus int

// ResourceCreateOptions contains some options for creating an API7 Cloud resource.
type ResourceCreateOptions struct {
	// Organization indicates where the resources are.
	// This field should be specified when users want to create resources.
	// in the organization. e.g., when inviting a member, the
	// Organization.ID should be specified.
	Organization *Organization
	// Cluster indicates where the resource belongs.
	// This field should be specified when users want to create resources
	// in the cluster. e.g., when creating Application, the
	// Cluster.ID should be specified.
	Cluster *Cluster
	// Application indicates which Application should this resource belong.
	// This field should be specified when users want to update sub-resources
	// in the Application. e.g., when creating API, CanaryRelease, the
	// Application.ID should be specified.
	Application *Application
}

// ResourceUpdateOptions contains some options for updating an API7 Cloud resource.
type ResourceUpdateOptions struct {
	// Organization indicates where the resources are.
	// This field should be specified when users want to update resources.
	// in the organization. e.g., when re-inviting a member, the
	// Organization.ID should be specified.
	Organization *Organization
	// Cluster indicates where the resource belongs.
	// This field should be specified when users want to update resources
	// in the cluster. e.g., when updating Application, the
	// Cluster.ID should be specified.
	Cluster *Cluster
	// Application indicates which Application should this resource belong.
	// This field should be specified when users want to update sub-resources
	// in the Application. e.g., when updating API, the
	// Application.ID should be specified.
	Application *Application
}

// ResourceDeleteOptions contains some options for deleting an API7 Cloud resource.
type ResourceDeleteOptions struct {
	// Organization indicates where the resources are.
	// This field should be specified when users want to update resources.
	// in the organization. e.g., when deleting a member, the
	// Organization.ID should be specified.
	Organization *Organization
	// Cluster indicates where the resource is.
	// This field should be specified when users want to delete resources
	// in the cluster. e.g., when deleting Application, the
	// Cluster.ID should be specified.
	Cluster *Cluster
	// Application indicates which Application should this resource belong.
	// This field should be specified when users want to delete sub-resources
	// in the Application. e.g., when deleting API, the
	// Application.ID should be specified.
	Application *Application
}

// ResourceGetOptions contains some options for getting an API7 Cloud resource.
type ResourceGetOptions struct {
	// Organization indicates where the resources are.
	// This field should be specified when users want to list resources.
	// in the organization. e.g., when getting a member, the
	// Organization.ID should be specified.
	Organization *Organization
	// Cluster indicates where the resource is.
	// This field should be specified when users want to get a resource.
	// in the cluster. e.g., when getting Application, the
	// Cluster.ID should be specified.
	Cluster *Cluster
	// Application indicates which Application should this resource belong.
	// This field should be specified when users want to fetch sub-resources
	// in the Application. e.g., when fetching API, the
	// Application.ID should be specified.
	Application *Application
}

// ResourceListOptions contains some options for listing the same kind of API7 Cloud resources.
type ResourceListOptions struct {
	// Organization indicates where the resources are.
	// This field should be specified when users want to list resources.
	// in the organization. e.g., when iterating Cluster, the
	// Organization.ID should be specified.
	Organization *Organization
	// Cluster indicates where the resources are.
	// This field should be specified when users want to list resources.
	// in the cluster. e.g., when iterating Application, the
	// Cluster.ID should be specified.
	Cluster *Cluster
	// Application indicates which Application should this resource belong.
	// This field should be specified when users want to list sub-resources
	// in the Application. e.g., when listing API, the
	// Application.ID should be specified.
	Application *Application
	// Pagination indicates the start page and the page size for listing resources.
	Pagination *Pagination
	// Filter indicates conditions to filter out resources.
	Filter *Filter
}

// ExpressionLogicalRelationship is the logical relationship between expressions.
type ExpressionLogicalRelationship string

const (
	// MatchAll meaning all the expressions should be matched.
	MatchAll ExpressionLogicalRelationship = "All"
	// MatchAny meaning any of the expressions should be matched.
	MatchAny ExpressionLogicalRelationship = "Any"
)

// ExpressionSubject is the subject category of the expression.
type ExpressionSubject string

const (
	// HeaderSubject indicates the expression subject is from a HTTP request header.
	HeaderSubject ExpressionSubject = "header"
	// QuerySubject indicates the expression subject is from the HTTP query string.
	QuerySubject ExpressionSubject = "query"
	// CookieSubject indicates the expression subject is from Cookie header.
	CookieSubject ExpressionSubject = "cookie"
	// PathSubject indicates the expression subject is from the URI path.
	PathSubject ExpressionSubject = "path"
	// VariableSubject indicates the expression subject is a Nginx or APISIX variable.
	VariableSubject ExpressionSubject = "variable"
)

// ExpressionOperator is the operator of the expression.
type ExpressionOperator string

const (
	// EqualOperator indicates the expression operator is "equal"
	EqualOperator ExpressionOperator = "equal"
	// NotEqualOperator indicates the expression operator is "not_equal"
	NotEqualOperator ExpressionOperator = "not_equal"
	// RegexMatchOperator indicates the expression operator is "regex_match"
	RegexMatchOperator ExpressionOperator = "regex_match"
	// RegexNotMatchOperator indicates the expression operator is "regex_not_match"
	RegexNotMatchOperator ExpressionOperator = "regex_not_match"
	// PresentOperator indicates the expression operator is "present"
	PresentOperator ExpressionOperator = "present"
	// NotPresentOperator indicates the expression operator is "not_present"
	NotPresentOperator ExpressionOperator = "not_present"
	// LargerEqualOperator indicates the expression operator is "larger_equal"
	LargerEqualOperator ExpressionOperator = "larger_equal"
	// LargerThanOperator indicates the expression operator is "larger_than"
	LargerThanOperator ExpressionOperator = "larger_than"
	// LessEqualOperator indicates the expression operator is "less_equal"
	LessEqualOperator ExpressionOperator = "less_equal"
	// LessThanOperator indicates the expression operator is "less_than"
	LessThanOperator ExpressionOperator = "less_than"
)

// Expression is the route match expressions.
type Expression struct {
	// Subject is the subject category of the expression.
	Subject ExpressionSubject `json:"subject,omitempty"`
	// Name is the subject name of the expression.
	Name string `json:"name,omitempty"`
	// Operator is the operator of the expression.
	Operator ExpressionOperator `json:"operator,omitempty"`
	// Value is the value that the expression should be matched.
	Value string `json:"value,omitempty"`
}
