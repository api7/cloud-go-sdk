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

// ServiceRegistryType describes the type of service registry.
type ServiceRegistryType int

const (
	// ServiceRegistryKubernetes indicates the kubernetes-type service registry.
	ServiceRegistryKubernetes = ServiceRegistryType(iota) + 1
)

// ServiceRegistrySpec is the service registry specification.
type ServiceRegistrySpec struct {
	// Name is the service registry name.
	Name string `json:"name"`
	// Enable indicates whether the service registry is enabled.
	Enabled bool `json:"enabled"`
	// Type is the service registry type.
	Type ServiceRegistryType `json:"type"`
	// Kubernetes is the kubernetes service registry.
	// It's valid only if Type is ServiceRegistryKubernetes.
	Kubernetes *KubernetesServiceRegistry `json:"kubernetes,omitempty"`
}

// KubernetesServiceRegistry is the Kubernetes registry.
type KubernetesServiceRegistry struct {
	// APIServer is the kubernetes api server config
	APIServer KubernetesAPIServer `json:"api_server"`
	// ServiceAccountTokenFile is the path of the service account token file
	ServiceAccountTokenFile string `json:"service_account_token_file,omitempty"`
	// ServiceAccountTokenValue is the service account token value
	ServiceAccountTokenValue string `json:"service_account_token_value,omitempty"`
	// NamespaceSelector is the namespace selector of kubernetes service discovery
	NamespaceSelector *KubernetesNamespaceSelector `json:"namespace_selector,omitempty"`
	// EndpointsLabelSelectors is the endpoints label selectors of kubernetes service discovery
	EndpointsLabelSelectors []KubernetesEndpointsLabelSelector `json:"endpoints_label_selectors,omitempty"`
}

// KubernetesAPIServer is configuration for the Kubernetes API server.
type KubernetesAPIServer struct {
	// Scheme is the scheme of http server
	Scheme string `json:"scheme,omitempty"`
	// Host is the host of http server
	Host string `json:"host,omitempty"`
	// Port is the port of http server
	Port int `json:"port,omitempty"`
}

// KubernetesNamespaceSelector is the namespace selector of kubernetes service discovery
type KubernetesNamespaceSelector struct {
	// Operator is the operator of the selector
	Operator string `json:"operator,omitempty"`
	// Patterns is the patterns of the selector
	Patterns []string `json:"patterns,omitempty"`
}

// KubernetesEndpointsLabelSelector is the label selector of kubernetes service discovery
type KubernetesEndpointsLabelSelector struct {
	// Key is the key of the label selector
	Key string `json:"key,omitempty"`
	// Operator is the operator of the selector
	Operator string `json:"operator,omitempty"`
	// Value is the value of the label selector
	Value string `json:"value,omitempty"`
}

// ServiceRegistry attaches some management field to ServiceRegistry.
type ServiceRegistry struct {
	ServiceRegistrySpec `json:",inline"`

	// ID is the service registry id.
	ID ID `json:"id" gorm:"column:id"`
	// ClusterID is id of control plane that current service registry belong with.
	ClusterID ID `json:"cluster_id"`
	// Status is status of service registry.
	Status EntityStatus `json:"status"`
	// CreatedAt is the object creation time
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the last modified time of this object
	UpdatedAt time.Time `json:"updated_at"`
}

// ServiceDiscoveryInterface is the interface for manipulating API7 Cloud service discovery features.
type ServiceDiscoveryInterface interface {
	// CreateServiceRegistry creates an API7 Cloud ServiceRegistry in the specified control plane.
	// The given `registry` parameter should specify the desired ServiceRegistry specification.
	// Users need to specify the Cluster in the `opts`.
	// The returned ServiceRegistry will contain the same ServiceRegistry specification plus some
	// management fields and default values.
	CreateServiceRegistry(ctx context.Context, registry *ServiceRegistry, opts *ResourceCreateOptions) (*ServiceRegistry, error)
	// UpdateServiceRegistry updates an existing API7 Cloud ServiceRegistry in the specified control plane.
	// The given `registry` parameter should specify the ServiceRegistry that you want to update.
	// Users need to specify the Cluster in the `opts`.
	// The returned ServiceRegistry will contain the same ServiceRegistry specification plus some
	// management fields and default values.
	UpdateServiceRegistry(ctx context.Context, registry *ServiceRegistry, opts *ResourceUpdateOptions) (*ServiceRegistry, error)
	// DeleteServiceRegistry deletes an existing API7 Cloud ServiceRegistry in the specified control plane.
	// The given `appID` parameter should specify the Application that you want to delete.
	// Users need to specify the Cluster in the `opts`.
	DeleteServiceRegistry(ctx context.Context, registryID ID, opts *ResourceDeleteOptions) error
	// GetServiceRegistry gets an existing API7 Cloud ServiceRegistry in the specified control plane.
	// The given `registryID` parameter should specify the ServiceRegistry that you want to get.
	// Users need to specify the Cluster in the `opts`.
	GetServiceRegistry(ctx context.Context, registryID ID, opts *ResourceGetOptions) (*ServiceRegistry, error)
	// ListServiceRegistries returns an iterator for listing service registries in the specified control plane
	// with the given list conditions.
	// Users need to specify the Cluster, Paging and Filter conditions (if necessary) in the `opts`.
	ListServiceRegistries(ctx context.Context, opts *ResourceListOptions) (ServiceRegistryListIterator, error)
}

// ServiceRegistryListIterator is an iterator for listing service registries.
type ServiceRegistryListIterator interface {
	// Next returns the next ServiceRegistry according to the filter conditions.
	Next() (*ServiceRegistry, error)
}

type serviceRegistryImpl struct {
	client httpClient
}

type serviceRegistryListIterator struct {
	iter listIterator
}

func (iter *serviceRegistryListIterator) Next() (*ServiceRegistry, error) {
	var registry ServiceRegistry
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	if err = json.Unmarshal(rawData, &registry); err != nil {
		return nil, err
	}
	return &registry, nil
}

func newServiceDiscovery(cli httpClient) ServiceDiscoveryInterface {
	return &serviceRegistryImpl{
		client: cli,
	}
}

func (impl *serviceRegistryImpl) CreateServiceRegistry(ctx context.Context, registry *ServiceRegistry, opts *ResourceCreateOptions) (*ServiceRegistry, error) {
	var createdRegistry ServiceRegistry

	cpID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", cpID.String(), "service_registries")
	err := impl.client.sendPostRequest(ctx, uri, "", registry, jsonPayloadDecodeFactory(&createdRegistry))
	if err != nil {
		return nil, err
	}
	return &createdRegistry, nil
}

func (impl *serviceRegistryImpl) UpdateServiceRegistry(ctx context.Context, registry *ServiceRegistry, opts *ResourceUpdateOptions) (*ServiceRegistry, error) {
	var updatedRegistry ServiceRegistry

	cpID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", cpID.String(), "service_registries", registry.ID.String())
	err := impl.client.sendPutRequest(ctx, uri, "", registry, jsonPayloadDecodeFactory(&updatedRegistry))
	if err != nil {
		return nil, err
	}
	return &updatedRegistry, nil
}

func (impl *serviceRegistryImpl) DeleteServiceRegistry(ctx context.Context, registryID ID, opts *ResourceDeleteOptions) error {
	cpID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", cpID.String(), "service_registries", registryID.String())
	return impl.client.sendDeleteRequest(ctx, uri, "", nil)
}

func (impl *serviceRegistryImpl) GetServiceRegistry(ctx context.Context, registryID ID, opts *ResourceGetOptions) (*ServiceRegistry, error) {
	var registry ServiceRegistry

	cpID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", cpID.String(), "service_registries", registryID.String())

	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&registry))
	if err != nil {
		return nil, err
	}
	return &registry, nil
}

func (impl *serviceRegistryImpl) ListServiceRegistries(ctx context.Context, opts *ResourceListOptions) (ServiceRegistryListIterator, error) {
	iter := listIterator{
		ctx:      ctx,
		resource: "service_registry",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "clusters", opts.Cluster.ID.String(), "service_registries"),
		paging:   mergePagination(opts.Pagination),
		filter:   opts.Filter,
	}

	return &serviceRegistryListIterator{iter: iter}, nil
}
