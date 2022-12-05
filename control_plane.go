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
	"github.com/pkg/errors"
	"path"
	"time"
)

// ControlPlaneStage is used to depict different control plane lifecycles.
type ControlPlaneStage int

func (cs ControlPlaneStage) String() string {
	switch cs {
	case ControlPlanePending:
		return "pending"
	case ControlPlaneCreating:
		return "creating"
	case ControlPlaneNormal:
		return "normal"
	case ControlPlaneCreateFailed:
		return "create failed"
	case ControlPlaneDeleting:
		return "deleting"
	case ControlPlaneDeleted:
		return "deleted"
	default:
		return "unknown"
	}
}

const (
	// ControlPlanePending means a control plane is not created yet.
	ControlPlanePending = ControlPlaneStage(iota + 1)
	// ControlPlaneCreating means a control plane is being created.
	ControlPlaneCreating
	// ControlPlaneNormal means a control plane was created, and now it's normal.
	ControlPlaneNormal
	// ControlPlaneCreateFailed means a control plane was not created successfully.
	ControlPlaneCreateFailed
	// ControlPlaneDeleting means a control plane is being deleted.
	ControlPlaneDeleting
	// ControlPlaneDeleted means a control plane was deleted.
	ControlPlaneDeleted
)

const (
	// DeleteURITailSlash means delete the tail slash of the request uri before matching
	DeleteURITailSlash = "Delete Tail Slash"
)

const (
	// RewriteServerHeader means rewrite the Server header in the response.
	RewriteServerHeader = "Rewrite"
	// HideVersionToken means hide the APISIX version info in the Server header.
	HideVersionToken = "Hide Version Token"
)

const (
	// RealIPPositionHeader indicates the real ip is in an HTTP header.
	RealIPPositionHeader = "header"
	// RealIPPositionQuery indicates the real ip is in a query string.
	RealIPPositionQuery = "query"
	// RealIPPositionCookie indicates the real ip is in cookie.
	RealIPPositionCookie = "cookie"
)

// GatewayInstanceStatus is the status of an gateway instance.
type GatewayInstanceStatus string

const (
	// GatewayInstanceHealthy indicates the instance is healthy. Note Healthy means
	// the heartbeat probes sent from instance are received periodically,
	// at the same while, the configuration delivery (currently it's ETCD
	// connections) is also normal.
	GatewayInstanceHealthy = GatewayInstanceStatus("Healthy")
	// GatewayInstanceOnlyHeartbeats indicates the instance sends heartbeat probes
	// periodically but the configuration cannot be delivered to the instance.
	GatewayInstanceOnlyHeartbeats = GatewayInstanceStatus("Only Heartbeats")
	// GatewayInstanceLostConnection indicate the instance lose heartbeat in short time(between InstanceLostConnectionThresholdDuration and InstanceOfflineThresholdDuration)
	GatewayInstanceLostConnection = GatewayInstanceStatus("Lost Connection")
	// GatewayInstanceOffline indicate the instance lose heartbeat in long time(out-of the InstanceLiveThresholdDuration)
	GatewayInstanceOffline = GatewayInstanceStatus("Offline")
)

// ControlPlane contains the control plane specification and management fields.
type ControlPlane struct {
	ControlPlaneSpec

	// ID is the unique identify of this control plane.
	ID ID `json:"id,inline"`
	// Name is the control plane name.
	Name string `json:"name"`
	// CreatedAt is the creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the last modified time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ControlPlaneSpec is the specification of control plane.
type ControlPlaneSpec struct {
	// OrganizationID refers to an Organization object, which
	// indicates the belonged organization for this control plane.
	OrganizationID ID `json:"org_id"`
	// RegionID refers to a Region object, which indicates the
	// region that the Cloud Plane resides.
	RegionID ID `json:"region_id"`
	// Status indicates the control plane status, candidate values are:
	// * ControlPlaneBuildInProgress: the control plane is being created.
	// * ControlPlaneCreating means a control plane is being created.
	// * ControlPlaneNormal: the control plane is built, and can be used normally.
	// * ControlPlaneCreateFailed means a control plane was not created successfully.
	// * ControlPlaneDeleting means a control plane is being deleted.
	// * ControlPlaneDeleted means a control plane was deleted.
	Status ControlPlaneStage `json:"status"`
	// Domain is the domain assigned by APISEVEN Cloud and has correct
	// records so that DP instances can access APISEVEN Cloud by it.
	Domain string `json:"domain"`
	// ConfigPayload is the customized data plane config for specific control plane
	ConfigPayload string `json:"-"`
	// Settings is the settings for the control plane.
	Settings ControlPlaneSettings `json:"settings"`
	// Plugins settings on Control Plane level
	Plugins Plugins `json:"policies,omitempty"`
	// ConfigVersion is the version for the control plane.
	ConfigVersion int `json:"config_version"`
}

// ControlPlaneSettings is control plane settings
type ControlPlaneSettings struct {
	// ClientSettings is the client settings config that used in apisix
	ClientSettings ClientSettings `json:"client_settings"`
	// ObservabilitySettings is the observability settings config that used in apisix
	ObservabilitySettings ObservabilitySettings `json:"observability_settings"`
	// APIProxySettings is the api proxy settings config that used in apisix
	APIProxySettings APIProxySettings `json:"api_proxy_settings"`
}

// APIProxySettings is the api proxy settings config
type APIProxySettings struct {
	// EnableRequestBuffering indicates whether to enable request buffering
	EnableRequestBuffering bool `json:"enable_request_buffering"`
	// ServerHeaderCustomization is the server header customization settings
	ServerHeaderCustomization *ServerHeaderCustomization `json:"server_header_customization,omitempty"`
	// URLHandlingOptions is the url handling options using in data plane
	// Optional values are:
	// * DeleteURITailSlash
	URLHandlingOptions []string `json:"url_handling_options"`
}

// ServerHeaderCustomization is the server header customization settings
type ServerHeaderCustomization struct {
	// Mode is the mode of the customization
	// Optional values can be:
	// * RewriteServerHeader, rewrite the server header, value is specified by `NewServerHeader`.
	// * HideServerToken, still use APISIX as the server header, but hide the version token.
	Mode string `json:"mode,omitempty"`
	// NewServerHeader is the new server header
	NewServerHeader string `json:"new_server_header,omitempty"`
}

// ClientSettings is the client settings config
type ClientSettings struct {
	// ClientRealIP is the client real ip config that used in apisix
	ClientRealIP ClientRealIPConfig `json:"client_real_ip"`
	// MaximumRequestBodySize is the maximum request body size that used in apisix, 0 means no limit
	MaximumRequestBodySize uint64 `json:"maximum_request_body_size"`
}

// ClientRealIPConfig is the client real ip config
type ClientRealIPConfig struct {
	// ReplaceFrom is the client ip replace from config
	ReplaceFrom ClientIPReplaceFrom `json:"replace_from"`
	// TrustedAddresses is the client ip trusted addresses
	TrustedAddresses []string `json:"trusted_addresses,omitempty"`
	// RecursiveSearch indicates whether the client ip is searched recursively
	RecursiveSearch bool `json:"recursive_search"`
	// Enable indicates whether real ip is enabled
	Enabled bool `json:"enabled"`
}

// ClientIPReplaceFrom is the client ip replace from config
type ClientIPReplaceFrom struct {
	// Position is the position that the client ip should be got from
	// Optional values are:
	// * RealIPPositionHeader, indicates the real ip is in an HTTP header, and the header name is specified by `Name` field.
	// * RealIPPositionQuery, indicates the real ip is in the query string, and the query name is specified by `Name` field.
	// * RealIPPositionCookie, indicates the real ip is in the Cookie, and the field name is specified by `Name` field.
	Position string `json:"position,omitempty"`
	// Name is the name of the variable that the client ip should be got from
	Name string `json:"name,omitempty"`
}

// ObservabilitySettings is the observability settings config
type ObservabilitySettings struct {
	Metrics MetricsConfig `json:"metrics,omitempty"`
	// ShowUpstreamStatusInResponseHeader indicates whether to show all upstream status
	// in `X-APISIX-Upstream-Status` header.
	// This header will be shown only when the status code is `5xx` when this field is diable.
	ShowUpstreamStatusInResponseHeader bool `json:"show_upstream_status_in_response_header"`
	// AccessLogRotate is the access log rotate settings config
	AccessLogRotate AccessLogRotateSettings `json:"access_log_rotate"`
}

// MetricsConfig contains configurations related to metrics.
type MetricsConfig struct {
	// Enable indicates whether gateway instances' metrics should be collected to API7 Cloud.
	Enabled bool `json:"enabled"`
}

// AccessLogRotateSettings is the access log rotate settings config
type AccessLogRotateSettings struct {
	// Enabled indicates whether access log rotation is enabled.
	Enabled bool `json:"enabled"`
	// Interval is time in seconds specifying how often to rotate the logs.
	Interval uint64 `json:"interval,omitempty"`
	// MaximumKeptLogEntries is the maximum number of log entries to keep.
	MaximumKeptLogEntries uint64 `json:"maximum_kept_log_entries,omitempty"`
	// EnableCompression indicates whether to compress the log files.
	EnableCompression bool `json:"enable_compression"`
}

// GatewayInstancePayload contains basic information for a gateway instance.
type GatewayInstancePayload struct {
	// ID is the unique identify for the APISEVEN instance.
	ID string `json:"id"`
	// Hostname is the name for the VM or container that the APISEVEN
	// instance resides.
	Hostname string `json:"hostname"`
	// IP is the IP address of the VM or container that the APISEVEN
	// instance resides.
	IP string `json:"ip"`
	// Domain is the domain that assigned by APISEVEN Cloud for the owner
	// (organization) of the APISEVEN instance.
	Domain string `json:"domain"`
	// APICalls is number of HTTP requests counting in report period
	APICalls uint64 `json:"api_calls"`
	// Version is the version of data plane
	Version string `json:"version"`
	// EtcdReachable indicates whether the instance can reach the etcd.
	EtcdReachable bool `json:"etcd_reachable"`
	// ConfigVersion is the version of the config currently in use on the data plane
	ConfigVersion uint64 `json:"config_version"`
}

// GatewayInstance shows the gateway instance (Apache APISIX) status.
type GatewayInstance struct {
	GatewayInstancePayload `json:",inline"`
	// LastSeenTime is the last time that Cloud seen this instance.
	// An instance should be marked as offline once the elapsed time is over
	// 30s since the last seen time.
	LastSeenTime time.Time `json:"last_seen_time"`
	// RegisterTime is the first time that Cloud seen this instance.
	RegisterTime time.Time `json:"register_time"`
	// Status is the instance status.
	Status GatewayInstanceStatus `json:"status"`
}

// ControlPlaneInterface is the interface for manipulating Control Plane.
type ControlPlaneInterface interface {
	// ListControlPlanes returns an iterator for listing Control Planes in the specified Organization with the
	// given list conditions.
	// Users need to specify the Organization, Paging conditions in the `opts`.
	ListControlPlanes(ctx context.Context, opts *ResourceListOptions) (ControlPlaneListIterator, error)
	// ListAllGatewayInstances returns all the gateway instances (ever) connected to the given Control Plane.
	// Note currently users don't need to pass the `opts` parameter. Just pass `nil` is OK.
	ListAllGatewayInstances(ctx context.Context, cpID ID, opts *ResourceListOptions) ([]GatewayInstance, error)
}

// ControlPlaneListIterator is an iterator for listing Control Planes.
type ControlPlaneListIterator interface {
	// Next returns the next Control Plane according to the filter conditions.
	Next() (*ControlPlane, error)
}

type controlPlaneImpl struct {
	client httpClient
}

type controlPlaneListIterator struct {
	iter listIterator
}

func (iter *controlPlaneListIterator) Next() (*ControlPlane, error) {
	var cp ControlPlane
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	if err = json.Unmarshal(rawData, &cp); err != nil {
		return nil, err
	}
	return &cp, nil
}

func newControlPlane(cli httpClient) ControlPlaneInterface {
	return &controlPlaneImpl{
		client: cli,
	}
}

func (impl *controlPlaneImpl) ListControlPlanes(ctx context.Context, opts *ResourceListOptions) (ControlPlaneListIterator, error) {
	iter := listIterator{
		ctx:      ctx,
		resource: "control plane",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "controlplanes"),
		paging:   mergePagination(opts.Pagination),
	}

	return &controlPlaneListIterator{iter: iter}, nil
}

func (impl *controlPlaneImpl) ListAllGatewayInstances(ctx context.Context, cpID ID, _ *ResourceListOptions) ([]GatewayInstance, error) {
	var (
		lr        listResponse
		instances []GatewayInstance
	)
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "instances")
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&lr))
	if err != nil {
		return nil, err
	}

	for i, raw := range lr.List {
		var instance GatewayInstance
		if err = json.Unmarshal(raw, &instance); err != nil {
			return nil, errors.Wrapf(err, "unmarshal gateway instance #%d", i)
		}
		instances = append(instances, instance)
	}

	return instances, nil
}
