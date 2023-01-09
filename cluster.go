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

// ClusterStage is used to depict different control plane lifecycles.
type ClusterStage int

func (cs ClusterStage) String() string {
	switch cs {
	case ClusterPending:
		return "pending"
	case ClusterCreating:
		return "creating"
	case ClusterNormal:
		return "normal"
	case ClusterCreateFailed:
		return "create failed"
	case ClusterDeleting:
		return "deleting"
	case ClusterDeleted:
		return "deleted"
	default:
		return "unknown"
	}
}

const (
	// ClusterPending means a control plane is not created yet.
	ClusterPending = ClusterStage(iota + 1)
	// ClusterCreating means a control plane is being created.
	ClusterCreating
	// ClusterNormal means a control plane was created, and now it's normal.
	ClusterNormal
	// ClusterCreateFailed means a control plane was not created successfully.
	ClusterCreateFailed
	// ClusterDeleting means a control plane is being deleted.
	ClusterDeleting
	// ClusterDeleted means a control plane was deleted.
	ClusterDeleted
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
	// the heartbeat probes sent from the instance are received periodically,
	// at the same while, the configuration delivery (currently it's ETCD
	// connections) is also normal.
	GatewayInstanceHealthy = GatewayInstanceStatus("Healthy")
	// GatewayInstanceOnlyHeartbeats indicates the instance sends heartbeat probes
	// periodically but the configuration cannot be delivered to the instance.
	GatewayInstanceOnlyHeartbeats = GatewayInstanceStatus("Only Heartbeats")
	// GatewayInstanceLostConnection indicate the instance lose heartbeat in short time(between InstanceLostConnectionThresholdDuration and InstanceOfflineThresholdDuration)
	GatewayInstanceLostConnection = GatewayInstanceStatus("Lost Connection")
	// GatewayInstanceOffline indicates the instance loses heartbeat for long time(out-of the InstanceLiveThresholdDuration)
	GatewayInstanceOffline = GatewayInstanceStatus("Offline")
)

// Cluster contains the control plane specification and management fields.
type Cluster struct {
	ClusterSpec

	// ID is the unique identify of this control plane.
	ID ID `json:"id,inline"`
	// Name is the control plane name.
	Name string `json:"name"`
	// CreatedAt is the creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the last modified time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ClusterSpec is the specification of control plane.
type ClusterSpec struct {
	// OrganizationID refers to an Organization object, which
	// indicates the belonged organization for this control plane.
	OrganizationID ID `json:"org_id"`
	// RegionID refers to a Region object, which indicates the
	// region that the Cloud Plane resides.
	RegionID ID `json:"region_id"`
	// Status indicates the control plane status, candidate values are:
	// * ClusterBuildInProgress: the control plane is being created.
	// * ClusterCreating means a control plane is being created.
	// * ClusterNormal: the control plane is built, and can be used normally.
	// * ClusterCreateFailed means a control plane was not created successfully.
	// * ClusterDeleting means a control plane is being deleted.
	// * ClusterDeleted means a control plane was deleted.
	Status ClusterStage `json:"status"`
	// Domain is the domain assigned by APISEVEN Cloud and has correct
	// records so that DP instances can access APISEVEN Cloud by it.
	Domain string `json:"domain"`
	// ConfigPayload is the customized  gateway config for specific control plane
	ConfigPayload string `json:"-"`
	// Settings is the settings for the control plane.
	Settings ClusterSettings `json:"settings"`
	// Plugins settings on Control Plane level
	Plugins Plugins `json:"policies,omitempty"`
	// ConfigVersion is the version for the control plane.
	ConfigVersion int `json:"config_version"`
}

// ClusterSettings is control plane settings
type ClusterSettings struct {
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
	// URLHandlingOptions is the url handling options using in  gateway
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

// TLSBundle contains a pair of certificate, private key,
// and the issuing certificate.
type TLSBundle struct {
	Certificate   string `json:"certificate"`
	PrivateKey    string `json:"private_key"`
	CACertificate string `json:"ca_certificate"`
}

// GatewayInstancePayload contains basic information for a gateway instance.
type GatewayInstancePayload struct {
	// ID is the unique identity for the APISEVEN instance.
	ID string `json:"id"`
	// Hostname is the name for the VM or container that the APISEVEN
	// instance resides.
	Hostname string `json:"hostname"`
	// IP is the IP address of the VM or container that the APISEVEN
	// instance resides.
	IP string `json:"ip"`
	// Domain is the domain assigned by APISEVEN Cloud for the owner
	// (organization) of the APISEVEN instance.
	Domain string `json:"domain"`
	// APICalls is the number of HTTP requests counted in the reporting period
	APICalls uint64 `json:"api_calls"`
	// Version is the version of the  gateway
	Version string `json:"version"`
	// EtcdReachable indicates whether the instance can reach the etcd.
	EtcdReachable bool `json:"etcd_reachable"`
	// ConfigVersion is the version of the config currently in use on the  gateway
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

// ClusterInterface is the interface for manipulating Control Plane.
type ClusterInterface interface {
	// GetCluster gets an existing API7 Cloud Cluster.
	// The given `clusterID` parameter should specify the Cluster that you want to get.
	// Users need to specify the Organization.ID in the `opts`.
	GetCluster(ctx context.Context, clusterID ID, opts *ResourceGetOptions) (*Cluster, error)
	// UpdateClusterSettings updates the ClusterSettings for the specified Cluster.
	// The given `clusterID` parameter should specify the Cluster that you want to update.
	// The given `settings` parameter should specify the new settings you want to apply.
	// Users need to specify the Organization.ID in the `opts`.
	UpdateClusterSettings(ctx context.Context, clusterID ID, settings *ClusterSettings, opts *ResourceUpdateOptions) error
	// UpdateClusterPlugins updates the plugins bound on the specified Cluster.
	// The given `clusterID` parameter should specify the Cluster that you want to update.
	// The given `plugins` parameter should specify the new plugins you want to bind.
	// Users need to specify the Organization.ID in the `opts`.
	UpdateClusterPlugins(ctx context.Context, clusterID ID, plugins Plugins, opts *ResourceUpdateOptions) error
	// ListClusters returns an iterator for listing Control Planes in the specified Organization with the
	// given list conditions.
	// Users need to specify the Organization, Paging, and Filter conditions (if necessary)
	// in the `opts`.
	ListClusters(ctx context.Context, opts *ResourceListOptions) (ClusterListIterator, error)
	// GenerateGatewaySideCertificate generates the tls bundle for gateway instances to communicate with
	// the specified Control Plane on API7 Cloud.
	// The `clusterID` parameter specifies the Control Plane ID.
	// Note currently users don't need to pass the `opts` parameter. Just pass `nil` is OK.
	GenerateGatewaySideCertificate(ctx context.Context, clusterID ID, opts *ResourceCreateOptions) (*TLSBundle, error)
	// ListAllGatewayInstances returns all the gateway instances (ever) connected to the given Control Plane.
	// Note currently users don't need to pass the `opts` parameter. Just pass `nil` is OK.
	ListAllGatewayInstances(ctx context.Context, clusterID ID, opts *ResourceListOptions) ([]GatewayInstance, error)
	// ListAllAPILabels lists all labels for API.
	// The `clusterID` parameter specifies the Control Plane ID.
	// Note currently users don't need to pass the `opts` parameter. Just pass `nil` is OK.
	// The returned label slice will be `nil` if there is no any labels for API.
	ListAllAPILabels(ctx context.Context, clusterID ID, opts *ResourceListOptions) ([]string, error)
	// ListAllApplicationLabels lists all labels for Application.
	// The `clusterID` parameter specifies the Control Plane ID.
	// Note currently users don't need to pass the `opts` parameter. Just pass `nil` is OK.
	// The returned label slice will be `nil` if there is no any labels for Application.
	ListAllApplicationLabels(ctx context.Context, clusterID ID, opts *ResourceListOptions) ([]string, error)
	// ListAllCertificateLabels lists all labels for Certificate.
	// The `clusterID` parameter specifies the Control Plane ID.
	// Note currently users don't need to pass the `opts` parameter. Just pass `nil` is OK.
	// The returned label slice will be `nil` if there is no any labels for Certificate.
	ListAllCertificateLabels(ctx context.Context, clusterID ID, opts *ResourceListOptions) ([]string, error)
	// ListAllConsumerLabels lists all labels for Consumer.
	// The `clusterID` parameter specifies the Control Plane ID.
	// Note currently users don't need to pass the `opts` parameter. Just pass `nil` is OK.
	// The returned label slice will be `nil` if there is no any labels for Consumer.
	ListAllConsumerLabels(ctx context.Context, clusterID ID, opts *ResourceListOptions) ([]string, error)
}

// ClusterListIterator is an iterator for listing Control Planes.
type ClusterListIterator interface {
	// Next returns the next Control Plane according to the filter conditions.
	Next() (*Cluster, error)
}

type clusterImpl struct {
	client httpClient
}

type clusterListIterator struct {
	iter listIterator
}

func (iter *clusterListIterator) Next() (*Cluster, error) {
	var cluster Cluster
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	if err = json.Unmarshal(rawData, &cluster); err != nil {
		return nil, err
	}
	return &cluster, nil
}

func newCluster(cli httpClient) ClusterInterface {
	return &clusterImpl{
		client: cli,
	}
}

func (impl *clusterImpl) GetCluster(ctx context.Context, clusterID ID, opts *ResourceGetOptions) (*Cluster, error) {
	var cluster Cluster

	uri := path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "clusters", clusterID.String())
	if err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&cluster)); err != nil {
		return nil, err
	}
	return &cluster, nil
}

func (impl *clusterImpl) UpdateClusterSettings(ctx context.Context, clusterID ID, settings *ClusterSettings, opts *ResourceUpdateOptions) error {
	uri := path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "clusters", clusterID.String(), "config")
	if err := impl.client.sendPatchRequest(ctx, uri, "", settings, nil); err != nil {
		return err
	}
	return nil
}

func (impl *clusterImpl) UpdateClusterPlugins(ctx context.Context, clusterID ID, plugins Plugins, opts *ResourceUpdateOptions) error {
	uri := path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "clusters", clusterID.String(), "plugins")
	if err := impl.client.sendPatchRequest(ctx, uri, "", plugins, nil); err != nil {
		return err
	}
	return nil
}

func (impl *clusterImpl) ListClusters(ctx context.Context, opts *ResourceListOptions) (ClusterListIterator, error) {
	iter := listIterator{
		ctx:      ctx,
		resource: "control plane",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "orgs", opts.Organization.ID.String(), "clusters"),
		paging:   mergePagination(opts.Pagination),
		filter:   opts.Filter,
	}

	return &clusterListIterator{iter: iter}, nil
}

func (impl *clusterImpl) GenerateGatewaySideCertificate(ctx context.Context, clusterID ID, _ *ResourceCreateOptions) (*TLSBundle, error) {
	var bundle TLSBundle

	uri := path.Join(_apiPathPrefix, "clusters", clusterID.String(), "dp_certificate")
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&bundle))
	if err != nil {
		return nil, err
	}
	return &bundle, nil
}

func (impl *clusterImpl) ListAllGatewayInstances(ctx context.Context, clusterID ID, _ *ResourceListOptions) ([]GatewayInstance, error) {
	var (
		lr        listResponse
		instances []GatewayInstance
	)
	uri := path.Join(_apiPathPrefix, "clusters", clusterID.String(), "instances")
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

func (impl *clusterImpl) ListAllAPILabels(ctx context.Context, clusterID ID, _ *ResourceListOptions) ([]string, error) {
	return impl.listAllLabels(ctx, clusterID, "api")
}

func (impl *clusterImpl) ListAllApplicationLabels(ctx context.Context, clusterID ID, _ *ResourceListOptions) ([]string, error) {
	return impl.listAllLabels(ctx, clusterID, "application")
}

func (impl *clusterImpl) ListAllConsumerLabels(ctx context.Context, clusterID ID, _ *ResourceListOptions) ([]string, error) {
	return impl.listAllLabels(ctx, clusterID, "consumer")
}

func (impl *clusterImpl) ListAllCertificateLabels(ctx context.Context, clusterID ID, _ *ResourceListOptions) ([]string, error) {
	return impl.listAllLabels(ctx, clusterID, "certificate")
}

func (impl *clusterImpl) listAllLabels(ctx context.Context, clusterID ID, resource string) ([]string, error) {
	var labels []string

	uri := path.Join(_apiPathPrefix, "clusters", clusterID.String(), "labels", resource)
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&labels))
	if err != nil {
		return nil, err
	}
	return labels, nil
}
