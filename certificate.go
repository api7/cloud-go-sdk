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

// Certificate is the definition of API7 Cloud Certificate, which also contains
// some management fields.
type Certificate struct {
	CertificateSpec `json:",inline"`

	// ID is the unique identify to mark an object.
	ID ID `json:"id"`
	// ControlPlaneID is id of control plane that current certificate belong with
	ControlPlaneID ID `json:"control_plane_id"`
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
	ServerCertificate CertificateType = "server"
	// ClientCertificate means client-type certificate
	ClientCertificate CertificateType = "client"
)

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
	// CreateCertificate creates an API7 Cloud Certificate in the specified control plane.
	// The given `cert` parameter should specify the desired Certificate specification.
	// Users need to specify the ControlPlane in the `opts`.
	// The returned Certificate will contain the same Certificate specification plus some
	// management fields and default values, the `PrivateKey` field will be empty.
	CreateCertificate(ctx context.Context, cert *Certificate, opts *ResourceCreateOptions) (*Certificate, error)
	// UpdateCertificate updates an existing API7 Cloud Certificate in the specified control plane.
	// The given `cert` parameter should specify the desired Certificate specification.
	// Users need to specify the ControlPlane in the `opts`.
	// The returned Certificate will contain the same Certificate specification plus some
	// management fields and default values, the `PrivateKey` field will be empty.
	UpdateCertificate(ctx context.Context, cert *Certificate, opts *ResourceUpdateOptions) (*Certificate, error)
	// DeleteCertificate deletes an existing API7 Cloud Certificate in the specified control plane.
	// The given `certID` parameter should specify the Certificate that you want to delete.
	// Users need to specify the ControlPlane in the `opts`.
	DeleteCertificate(ctx context.Context, certID ID, opts *ResourceDeleteOptions) error
	// GetCertificate gets an existing API7 Cloud Certificate in the specified control plane.
	// The given `certID` parameter should specify the Certificate that you want to get.
	// Users need to specify the ControlPlane in the `opts`.
	// The `PrivateKey` field will be empty in the returned Certificate.
	GetCertificate(ctx context.Context, certID ID, opts *ResourceGetOptions) (*Certificate, error)
	// ListCertificates returns an iterator for listing Certificates in the specified control plane with the
	// given list conditions.
	// Users need to specify the ControlPlane, Paging and Filter conditions (if necessary)
	// in the `opts`.
	// The `PrivateKey` field will be empty in the returned Certificate.
	ListCertificates(ctx context.Context, opts *ResourceListOptions) (CertificateListIterator, error)
}

// CertificateListIterator is an iterator for listing Certificates.
type CertificateListIterator interface {
	// Next returns the next Certificate according to the filter conditions.
	Next() (*Certificate, error)
}

type certificateImpl struct {
	client httpClient
}
type certificatesListIterator struct {
	iter listIterator
}

func (iter *certificatesListIterator) Next() (*Certificate, error) {
	var cert Certificate
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

func (impl *certificateImpl) CreateCertificate(ctx context.Context, cert *Certificate, opts *ResourceCreateOptions) (*Certificate, error) {
	var createdCert Certificate

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "certificates")
	err := impl.client.sendPostRequest(ctx, uri, "", cert, jsonPayloadDecodeFactory(&createdCert))
	if err != nil {
		return nil, err
	}
	return &createdCert, nil
}

func (impl *certificateImpl) UpdateCertificate(ctx context.Context, cert *Certificate, opts *ResourceUpdateOptions) (*Certificate, error) {
	var updatedCert Certificate

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "certificates", cert.ID.String())
	err := impl.client.sendPutRequest(ctx, uri, "", cert, jsonPayloadDecodeFactory(&updatedCert))
	if err != nil {
		return nil, err
	}
	return &updatedCert, nil
}

func (impl *certificateImpl) DeleteCertificate(ctx context.Context, certID ID, opts *ResourceDeleteOptions) error {
	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "certificates", certID.String())
	return impl.client.sendDeleteRequest(ctx, uri, "", nil)
}

func (impl *certificateImpl) GetCertificate(ctx context.Context, certID ID, opts *ResourceGetOptions) (*Certificate, error) {
	var cert Certificate

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "certificates", certID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&cert))
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
		path:     path.Join(_apiPathPrefix, "controlplanes", opts.ControlPlane.ID.String(), "certificates"),
		paging:   mergePagination(opts.Pagination),
		filter:   opts.Filter,
	}

	return &certificatesListIterator{iter: iter}, nil
}
