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
	CanaryReleaseStatePaused = "pause"
	// CanaryReleaseStateInProgress indicates the in_progress state of CanaryRelease.
	CanaryReleaseStateInProgress = "in_progress"
	// CanaryReleaseStateFinished indicates the finish state of CanaryRelease.
	CanaryReleaseStateFinished = "finish"
)

// EntityStatus is common status definition for any kind of entity:
// * Uninitialized represents the entity has been saved to the db, but the associated resource has not yet been ready.
// * Normal indicates that the entity and associated resources are ready.
// * Deleted indicates the entity has been deleted.
type EntityStatus int

// ResourceCreateOptions contains some options for creating an API7 Cloud resource.
type ResourceCreateOptions struct {
	// ControlPlane indicates where the resource belongs.
	// This field should be specified when users want to create resources
	// in the control plane. e.g., when creating Application, the
	// ControlPlane.ID should be specified.
	ControlPlane *ControlPlane
	// Application indicates which Application should this resource belong.
	// This field should be specified when users want to update sub-resources
	// in the Application. e.g., when creating API, CanaryRelease, the
	// Application.ID should be specified.
	Application *Application
}

// ResourceUpdateOptions contains some options for updating an API7 Cloud resource.
type ResourceUpdateOptions struct {
	// ControlPlane indicates where the resource belongs.
	// This field should be specified when users want to update resources
	// in the control plane. e.g., when updating Application, the
	// ControlPlane.ID should be specified.
	ControlPlane *ControlPlane
	// Application indicates which Application should this resource belong.
	// This field should be specified when users want to update sub-resources
	// in the Application. e.g., when updating API, the
	// Application.ID should be specified.
	Application *Application
}

// ResourceDeleteOptions contains some options for deleting an API7 Cloud resource.
type ResourceDeleteOptions struct {
	// ControlPlane indicates where the resource is.
	// This field should be specified when users want to delete resources
	// in the control plane. e.g., when deleting Application, the
	// ControlPlane.ID should be specified.
	ControlPlane *ControlPlane
	// Application indicates which Application should this resource belong.
	// This field should be specified when users want to delete sub-resources
	// in the Application. e.g., when deleting API, the
	// Application.ID should be specified.
	Application *Application
}

// ResourceGetOptions contains some options for getting an API7 Cloud resource.
type ResourceGetOptions struct {
	// ControlPlane indicates where the resource is.
	// This field should be specified when users want to get a resource.
	// in the control plane. e.g., when getting Application, the
	// ControlPlane.ID should be specified.
	ControlPlane *ControlPlane
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
	// in the organization. e.g., when iterating ControlPlane, the
	// Organization.ID should be specified.
	Organization *Organization
	// ControlPlane indicates where the resources are.
	// This field should be specified when users want to list resources.
	// in the control plane. e.g., when iterating Application, the
	// ControlPlane.ID should be specified.
	ControlPlane *ControlPlane
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
