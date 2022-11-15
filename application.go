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
	// Policy settings on Application level
	Policies Policies `json:"policies,omitempty"`
	// Upstream settings for the Application
	Upstreams []UpstreamAndVersion `json:"upstreams"`
	// DefaultUpstreamVersion settings for the upstream that should be used
	DefaultUpstreamVersion string `json:"default_upstream_version,omitempty"`
	// Active is status of application
	Active int `json:"active"`
}
