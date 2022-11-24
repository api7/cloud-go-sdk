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
	"errors"
	"io/ioutil"
	"path"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCertificate(t *testing.T) {
	t.Parallel()
	key, err := ioutil.ReadFile("./testdata/test.key")
	if err != nil {
		assert.Nil(t, err, "read test key error")
		return
	}
	cert, err := ioutil.ReadFile("./testdata/test.pem")
	if err != nil {
		assert.Nil(t, err, "read test cert error")
		return
	}
	testCases := []struct {
		name          string
		pendingCert   *Certificate
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "create successfully",
			pendingCert: &Certificate{
				CertificateSpec: CertificateSpec{
					Certificate: string(cert),
					PrivateKey:  string(key),
					Labels:      nil,
					Type:        "Client",
				},
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/certificates"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingCert: &Certificate{
				CertificateSpec: CertificateSpec{
					Certificate: "invalid cert",
					PrivateKey:  "invalid  key",
					Labels:      nil,
					Type:        "Client",
				},
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPostRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/certificates"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newCertificate(cli).CreateCertificate(context.Background(), tc.pendingCert, &ResourceCreateOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check certificates create error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestUpdateCertificate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		pendingCert   *Certificate
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "update successfully",
			pendingCert: &Certificate{
				CertificateSpec: CertificateSpec{
					Certificate: "invalid cert",
					PrivateKey:  "invalid  key",
					Labels:      nil,
					Type:        "Client",
				},
				ID: 12,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/certificates/12"), "", gomock.Any(), gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			pendingCert: &Certificate{
				CertificateSpec: CertificateSpec{
					Certificate: "invalid cert",
					PrivateKey:  "invalid  key",
					Labels:      nil,
					Type:        "Client",
				},
				ID: 12,
			},
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendPutRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/certificates/12"), "", gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			_, err := newCertificate(cli).UpdateCertificate(context.Background(), tc.pendingCert, &ResourceUpdateOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check certificates update error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestDeleteCertificate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "delete successfully",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/certificates/12"), "", nil).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendDeleteRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/certificates/12"), "", nil).Return(errors.New("mock error"))
				return cli

			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			err := newCertificate(cli).DeleteCertificate(context.Background(), 12, &ResourceDeleteOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check certificates delete error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestGetCertificate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError string
		mockFunc      func(t *testing.T) httpClient
	}{
		{
			name: "get successfully",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/certificates/12"), "", gomock.Any()).Return(nil)
				return cli

			},
			expectedError: "",
		},
		{
			name: "mock error",
			mockFunc: func(t *testing.T) httpClient {
				ctrl := gomock.NewController(t)
				cli := NewMockhttpClient(ctrl)
				cli.EXPECT().sendGetRequest(gomock.Any(), path.Join(_apiPathPrefix, "/controlplanes/1/certificates/12"), "", gomock.Any()).Return(errors.New("mock error"))
				return cli
			},
			expectedError: "mock error",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cli := tc.mockFunc(t)
			// ignore the certificates check since currently we don't mock it, and the certificate is always a zero value.
			_, err := newCertificate(cli).GetCertificate(context.Background(), 12, &ResourceGetOptions{
				ControlPlane: &ControlPlane{
					ID: 1,
				},
			})
			if tc.expectedError == "" {
				assert.Nil(t, err, "check certificates get error")
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "check the error details")
			}
		})
	}
}

func TestListCertificates(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		iterator *certificatesListIterator
	}{
		{
			name: "create iterator successfully",
			iterator: &certificatesListIterator{
				iter: listIterator{
					resource: "certificates",
					path:     "/api/v1/controlplanes/123/certificates",
					paging: Pagination{
						Page:     14,
						PageSize: 25,
					},
					eof: false,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// ignore the certificates check since currently we don't mock it, and the certificate is always a zero value.
			raw, err := newCertificate(nil).ListCertificates(context.Background(), &ResourceListOptions{
				ControlPlane: &ControlPlane{
					ID: 123,
				},
				Pagination: &Pagination{
					Page:     14,
					PageSize: 25,
				},
			})
			assert.Nil(t, err, "check list certificates error")
			iter := raw.(*certificatesListIterator)
			assert.Equal(t, tc.iterator.iter.resource, iter.iter.resource, "check resource")
			assert.Equal(t, tc.iterator.iter.path, iter.iter.path, "check path")
			assert.Equal(t, tc.iterator.iter.paging.Page, iter.iter.paging.Page, "check page")
			assert.Equal(t, tc.iterator.iter.paging.PageSize, iter.iter.paging.PageSize, "check page size")
		})
	}
}
