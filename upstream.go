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
	"fmt"
)

// ServiceRegistryType describes the type of service registry.
type ServiceRegistryType int

const (
	// ServiceRegistryKubernetes indicates the kubernetes-type service registry.
	ServiceRegistryKubernetes = ServiceRegistryType(iota) + 1
)

const (
	// HealthCheckTypeTCP indicates a TCP-type health check.
	HealthCheckTypeTCP = "tcp"
	// HealthCheckTypeHTTP indicates an HTTP-type health check.
	HealthCheckTypeHTTP = "http"
	// HealthCheckTypeHTTPS indicates an HTTPS-type health check.
	HealthCheckTypeHTTPS = "https"
)

// UpstreamAndVersion contains both the upstream definition and the version information.
type UpstreamAndVersion struct {
	Upstream Upstream `json:"upstream"`
	// Version information about this upstream
	Version string `json:"version"`
	// ClientCertID settings the client cert for communicating with the upstream
	// Deprecated: use Upstream.ClientCertID.
	ClientCertID ID `json:"client_cert_id,omitempty"`
}

// Upstream is the definition of the upstream on Application.
type Upstream struct {
	// The scheme to communicate with the upstream
	Scheme string `json:"scheme"`
	// LBType is the load balancing strategy of the upstream
	LBType string `json:"lb_type,omitempty"`
	// HashKey is the hash key used to balance the upstream
	HashKey string `json:"hash_key,omitempty"`
	// ServiceDiscovery is the service discovery of the upstream
	ServiceDiscovery *UpstreamServiceDiscovery `json:"service_discovery,omitempty"`
	// The upstream endpoints
	Targets []UpstreamTarget `json:"targets,omitempty"`
	// Retries is sets the number of retries while passing the request to Upstream using the underlying Nginx mechanism.
	Retries *int `json:"retries,omitempty"`
	// Timeout is sets the timeout for connecting to, and sending and receiving messages to and from the Upstream
	Timeout *UpstreamTimeout `json:"timeout,omitempty"`
	// UpstreamHostMode configures the host header when the request is forwarded to the upstream
	UpstreamHostMode string `json:"upstream_host_mode,omitempty"`
	// UpstreamHost specifies the host of the Upstream request, this is only valid if the upstream_host_mode is set to rewrite
	UpstreamHost string `json:"upstream_host,omitempty"`
	// ClientCertID settings the client cert for communicating with the upstream
	ClientCertID ID `json:"client_cert_id,omitempty"`
	//Checks the data of health check
	Checks *Checks `json:"checks,omitempty"`
}

// Checks the data of health check
type Checks struct {
	Active  *ActiveHealthCheck  `json:"active,omitempty"`
	Passive *PassiveHealthCheck `json:"passive,omitempty"`
}

// ActiveHealthCheck the data of active health check
type ActiveHealthCheck struct {
	Type  string                    `json:"type"`
	HTTP  *HTTPActiveHealthCheck    `json:"http"`
	HTTPS *HTTPSActiveHealthCheck   `json:"https"`
	TCP   *TCPActiveCheckPredicates `json:"tcp"`
}

// HTTPActiveHealthCheck is the configuration of HTTP active health check
type HTTPActiveHealthCheck struct {
	ProbeTimeout     int64                   `json:"probe_timeout,omitempty"`
	ConcurrentProbes int64                   `json:"concurrent_probes,omitempty"`
	HTTPProbePath    string                  `json:"http_probe_path,omitempty"`
	HTTPProbeHost    string                  `json:"http_probe_host,omitempty"`
	ProbeTargetPort  int64                   `json:"probe_target_port,omitempty"`
	HTTPProbeHeaders ProbeHeader             `json:"http_probe_headers,omitempty"`
	Healthy          HTTPHealthyPredicates   `json:"healthy,omitempty"`
	UnHealthy        HTTPUnhealthyPredicates `json:"unhealthy,omitempty"`
}

// ProbeHeader indicates headers that will be taken in probe requests.
type ProbeHeader map[string]string

func (header ProbeHeader) ToStringArray() []string {
	var res []string
	if header == nil {
		return res
	}
	for k, v := range header {
		s := fmt.Sprintf("%v: %v", k, v)
		res = append(res, s)
	}
	return res
}

// HTTPSActiveHealthCheck the data of active health check for https
type HTTPSActiveHealthCheck struct {
	HTTPActiveHealthCheck

	VerifyTargetTlsCertificate bool `json:"verify_target_tls_certificate"`
}

// HTTPHealthyPredicates healthy predicates.
type HTTPHealthyPredicates struct {
	TargetsCheckInterval int64 `json:"targets_check_interval,omitempty"`
	HTTPStatusCodes      []int `json:"http_status_codes,omitempty"`
	Successes            int64 `json:"successes,omitempty"`
}

// HTTPUnhealthyPredicates unhealthy predicates.
type HTTPUnhealthyPredicates struct {
	TargetsCheckInterval int64 `json:"targets_check_interval,omitempty"`
	HTTPStatusCodes      []int `json:"http_status_codes,omitempty"`
	HTTPFailures         int64 `json:"http_failures,omitempty"`
	Timeouts             int64 `json:"timeouts,omitempty"`
}

// TCPActiveCheckPredicates predicates for the TCP probe active health check
type TCPActiveCheckPredicates struct {
	ProbeTimeout     int64                   `json:"probe_timeout,omitempty"`
	ConcurrentProbes int64                   `json:"concurrent_probes,omitempty"`
	ProbeTargetPort  int64                   `json:"probe_target_port,omitempty"`
	Healthy          *TCPHealthyPredicates   `json:"healthy,omitempty"`
	UnHealthy        *TCPUnhealthyPredicates `json:"unhealthy,omitempty"`
}

// TCPHealthyPredicates the healthy case data of tcp health check.
type TCPHealthyPredicates struct {
	TargetsCheckInterval int64 `json:"targets_check_interval,omitempty"`
	Successes            int64 `json:"successes,omitempty"`
}

// TCPUnhealthyPredicates the unhealthy case data of tcp health check.
type TCPUnhealthyPredicates struct {
	TargetsCheckInterval int64 `json:"targets_check_interval,omitempty"`
	TcpFailures          int64 `json:"tcp_failures,omitempty"`
	Timeouts             int64 `json:"timeouts,omitempty"`
}

// PassiveHealthCheck the data of passive health check
type PassiveHealthCheck struct {
	Type  string                     `json:"type"`
	HTTP  *HTTPPassiveHealthCheck    `json:"http"`
	HTTPS *HTTPPassiveHealthCheck    `json:"https"`
	TCP   *TCPPassiveCheckPredicates `json:"tcp"`
}

// HTTPPassiveHealthCheck is the configuration of HTTP passive health check
type HTTPPassiveHealthCheck struct {
	Healthy   HTTPHealthyPredicatesForPassive   `json:"healthy,omitempty"`
	UnHealthy HTTPUnhealthyPredicatesForPassive `json:"unhealthy,omitempty"`
}

// HTTPHealthyPredicatesForPassive healthy predicates for passive health check.
type HTTPHealthyPredicatesForPassive struct {
	HTTPStatusCodes []int `json:"http_status_codes,omitempty"`
}

// HTTPUnhealthyPredicatesForPassive unhealthy predicates for passive health check.
type HTTPUnhealthyPredicatesForPassive struct {
	HTTPStatusCodes []int `json:"http_status_codes,omitempty"`
	HTTPFailures    int64 `json:"http_failures,omitempty"`
	Timeouts        int64 `json:"timeouts,omitempty"`
}

// TCPPassiveCheckPredicates predicates for the TCP probe passive health check
type TCPPassiveCheckPredicates struct {
	UnHealthy *TCPUnhealthyPredicatesForPassive `json:"unhealthy,omitempty"`
}

// TCPUnhealthyPredicatesForPassive the unhealthy case data of passive tcp health check.
type TCPUnhealthyPredicatesForPassive struct {
	TcpFailures int64 `json:"tcp_failures,omitempty"`
	Timeouts    int64 `json:"timeouts,omitempty"`
}

// UpstreamServiceDiscovery is the service discovery of the upstream.
type UpstreamServiceDiscovery struct {
	// ServiceRegistry is the type of service registry
	ServiceRegistry ServiceRegistryType `json:"service_registry"`
	// ServiceRegistryID is the id of service registry
	ServiceRegistryID ID `json:"service_registry_id"`
	// KubernetesService is the kubernetes service discovery of the upstream
	KubernetesService KubernetesUpstreamServiceDiscovery `json:"kubernetes_service"`
}

// KubernetesUpstreamServiceDiscovery is the kubernetes service discovery of the upstream.
type KubernetesUpstreamServiceDiscovery struct {
	// Namespace is the namespace of the kubernetes endpoint
	Namespace string `json:"namespace"`
	// Name is the name of the kubernetes endpoint
	Name string `json:"name"`
	// Port is the target port of the kubernetes endpoint
	Port string `json:"port"`
}

// UpstreamTarget is the definition for an upstream endpoint.
type UpstreamTarget struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}

// UpstreamTimeout is the timeout for connecting to, and sending and receiving messages to and from the Upstream, value in seconds.
type UpstreamTimeout struct {
	Connect int `json:"connect,omitempty"`
	Send    int `json:"send,omitempty"`
	Read    int `json:"read,omitempty"`
}
