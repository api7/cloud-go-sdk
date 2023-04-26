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

// Certificate is the definition of API7 Cloud Certificate, which also contains
// some management fields.
type Certificate struct {
	CertificateSpec `json:",inline"`

	// ID is the unique identify to mark an object.
	ID ID `json:"id"`
	// ClusterID is id of cluster that current certificate belong with
	ClusterID ID `json:"cluster_id"`
	// Status is status of certificate
	Status EntityStatus `json:"status"`
	// CreatedAt is the object creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the last modified time of this object.
	UpdatedAt time.Time `json:"updated_at"`
}

// CertificateType is the type of log collection
type CertificateType string

const (
	// ServerCertificate means server-type certificate
	ServerCertificate CertificateType = "Server"
	// ClientCertificate means client-type certificate
	ClientCertificate CertificateType = "Client"
)

// CertificateMetadata contains the metadata of an user uploaded certificate.
type CertificateMetadata struct {
	// SNIs is service name indicates of certificate
	SNIs []string `json:"snis"`
	// NotBefore is valid after this time
	NotBefore time.Time `json:"not_before"`
	// NotAfter is invalid after this time
	NotAfter time.Time `json:"not_after"`
	// Subject is subject of certificate, contains fields like country, organization, common name...
	Subject string `json:"subject"`
	// Issuer is issuer of certificate
	Issuer string `json:"issuer"`
	// SerialNumber is serial number of certificate
	SerialNumber string `json:"serial_number"`
	// SignatureAlgorithm is signature algorithm of certificate
	SignatureAlgorithm string `json:"signature_algorithm"`
	// Extensions is extensions of certificate
	Extensions map[string]string `json:"extensions,omitempty"`
}

// CertificateDetails contains the details of the user uploaded certificate.
type CertificateDetails struct {
	// Extensions is extensions of certificate
	Extensions map[string]string `json:"extensions,omitempty"`
	// Issuer is issuer of certificate
	Issuer string `json:"issuer"`
	// NotBefore is valid after this time
	NotBefore time.Time `json:"not_before"`
	// NotAfter is invalid after this time
	NotAfter time.Time `json:"not_after"`
	// SNIs is service name indicates of certificate
	SNIs []string `json:"snis"`
	// SerialNumber is serial number of certificate
	SerialNumber string `json:"serial_number"`
	// Subject is subject of certificate, contains fields like country, organization, common name...
	Subject string `json:"subject"`
	// SignatureAlgorithm is signature algorithm of certificate
	SignatureAlgorithm string `json:"signature_algorithm"`
	// ClusterID is id of cluster that current certificate belong with
	ClusterID ID `json:"cluster_id"`
	// CreatedAt is the object creation time.
	CreatedAt time.Time `json:"created_at"`
	// ID is the unique identify to mark an object.
	ID ID `json:"id"`
	// CACertificate is CA certificate to verify client certificate
	CACertificate *CertificateMetadata `json:"ca_certificate,omitempty"`
	// Status is status of certificate
	Status EntityStatus `json:"status"`
	// UpdatedAt is the last modified time of this object.
	UpdatedAt time.Time `json:"updated_at"`
	// Labels are used for resource classification and indexing
	Labels []string `json:"labels,omitempty"`
	// Type is certificate type
	Type string `json:"type"`
}

// CertificateSpec is the specification of the Certificate
type CertificateSpec struct {
	// The certificate in PEM format.
	Certificate string `json:"certificate"`
	// Private key is the private key for the corresponding certificate in PEM format.
	PrivateKey string `json:"private_key"`
	// CACertificate is the client ca certificate in PEM format used to verify client certificate.
	CACertificate string `json:"ca_certificate,omitempty"`
	// Labels are used for resource classification and indexing.
	Labels []string `json:"labels,omitempty"`
	// Type is the certificate type, optional values can be:
	//   * ServerCertificate, server-type certificate.
	//   * ClientCertificate, client-type certificate.
	Type CertificateType `json:"type,omitempty"`
}

// CertificateInterface is the interface for manipulating Certificates.
type CertificateInterface interface {
	// CreateCertificate creates an API7 Cloud Certificate in the specified cluster.
	// The given `cert` parameter should specify the desired Certificate specification.
	// Users need to specify the Cluster in the `opts`.
	// The returned Certificate will contain the same Certificate specification plus some
	// management fields and default values, the `PrivateKey` field will be empty.
	CreateCertificate(ctx context.Context, cert *Certificate, opts *ResourceCreateOptions) (*CertificateDetails, error)
	// UpdateCertificate updates an existing API7 Cloud Certificate in the specified cluster.
	// The given `cert` parameter should specify the desired Certificate specification.
	// Users need to specify the Cluster in the `opts`.
	// The returned Certificate will contain the same Certificate specification plus some
	// management fields and default values, the `PrivateKey` field will be empty.
	UpdateCertificate(ctx context.Context, cert *Certificate, opts *ResourceUpdateOptions) (*CertificateDetails, error)
	// DeleteCertificate deletes an existing API7 Cloud Certificate in the specified cluster.
	// The given `certID` parameter should specify the Certificate that you want to delete.
	// Users need to specify the Cluster in the `opts`.
	DeleteCertificate(ctx context.Context, certID ID, opts *ResourceDeleteOptions) error
	// GetCertificate gets an existing API7 Cloud Certificate in the specified cluster.
	// The given `certID` parameter should specify the Certificate that you want to get.
	// Users need to specify the Cluster in the `opts`.
	// The `PrivateKey` field will be empty in the returned Certificate.
	GetCertificate(ctx context.Context, certID ID, opts *ResourceGetOptions) (*CertificateDetails, error)
	// ListCertificates returns an iterator for listing Certificates in the specified cluster with the
	// given list conditions.
	// Users need to specify the Cluster, Paging and Filter conditions (if necessary)
	// in the `opts`.
	// The `PrivateKey` field will be empty in the returned Certificate.
	ListCertificates(ctx context.Context, opts *ResourceListOptions) (CertificateListIterator, error)
	// DebugCertificateResources returns the corresponding translated APISIX resources for this Certificate.
	// The given `certID` parameter should specify the Certificate that you want to operate.
	// Users need to specify the Cluster.ID in the `opts`.
	// Note, the private key won't be returned due to the security concerns.
	DebugCertificateResources(ctx context.Context, appID ID, opts *ResourceGetOptions) (string, error)
}

// CertificateListIterator is an iterator for listing Certificates.
type CertificateListIterator interface {
	// Next returns the next Certificate according to the filter conditions.
	Next() (*CertificateDetails, error)
}

type certificateImpl struct {
	client httpClient
}
type certificatesListIterator struct {
	iter listIterator
}

func (iter *certificatesListIterator) Next() (*CertificateDetails, error) {
	var cert CertificateDetails
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	return &cert, nil
}

func newCertificate(cli httpClient) CertificateInterface {
	return &certificateImpl{
		client: cli,
	}
}

func (impl *certificateImpl) CreateCertificate(ctx context.Context, cert *Certificate, opts *ResourceCreateOptions) (*CertificateDetails, error) {
	var createdCert CertificateDetails

	clusterID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", clusterID.String(), "certificates")
	err := impl.client.sendPostRequest(ctx, uri, "", cert, jsonPayloadDecodeFactory(&createdCert), appendHeader(mapClusterIdFromOpts(opts)))
	if err != nil {
		return nil, err
	}
	return &createdCert, nil
}

func (impl *certificateImpl) UpdateCertificate(ctx context.Context, cert *Certificate, opts *ResourceUpdateOptions) (*CertificateDetails, error) {
	var updatedCert CertificateDetails

	clusterID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", clusterID.String(), "certificates", cert.ID.String())
	err := impl.client.sendPutRequest(ctx, uri, "", cert, jsonPayloadDecodeFactory(&updatedCert), appendHeader(mapClusterIdFromOpts(opts)))
	if err != nil {
		return nil, err
	}
	return &updatedCert, nil
}

func (impl *certificateImpl) DeleteCertificate(ctx context.Context, certID ID, opts *ResourceDeleteOptions) error {
	clusterID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", clusterID.String(), "certificates", certID.String())
	return impl.client.sendDeleteRequest(ctx, uri, "", nil, appendHeader(mapClusterIdFromOpts(opts)))
}

func (impl *certificateImpl) GetCertificate(ctx context.Context, certID ID, opts *ResourceGetOptions) (*CertificateDetails, error) {
	var cert CertificateDetails

	clusterID := opts.Cluster.ID
	uri := path.Join(_apiPathPrefix, "clusters", clusterID.String(), "certificates", certID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&cert), appendHeader(mapClusterIdFromOpts(opts)))
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func (impl *certificateImpl) ListCertificates(ctx context.Context, opts *ResourceListOptions) (CertificateListIterator, error) {
	iter := listIterator{
		ctx:      ctx,
		resource: "certificates",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "clusters", opts.Cluster.ID.String(), "certificates"),
		paging:   mergePagination(opts.Pagination),
		filter:   opts.Filter,
		headers:  appendHeader(mapClusterIdFromOpts(opts)),
	}

	return &certificatesListIterator{iter: iter}, nil
}

func (impl *certificateImpl) DebugCertificateResources(ctx context.Context, certID ID, opts *ResourceGetOptions) (string, error) {
	var rawData json.RawMessage
	uri := path.Join(_apiPathPrefix, "debug", "config", "clusters", opts.Cluster.ID.String(), "certificate", certID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&rawData), appendHeader(mapClusterIdFromOpts(opts)))
	if err != nil {
		return "", err
	}
	return formatJSONData(rawData)
}
