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

	"github.com/pkg/errors"
)

// Interface is the entrypoint of the Cloud Go SDK.
type Interface interface {
	SetGlobalClusterID(id ID)
	TraceInterface
	UserInterface
	AuthInterface
	ApplicationInterface
	APIInterface
	ClusterInterface
	OrganizationInterface
	RegionInterface
	CanaryReleaseInterface
	CertificateInterface
	ConsumerInterface
	LogCollectionInterface
	ServiceDiscoveryInterface
}

// AccessToken is the token used by API7 Cloud to authenticate clients.
type AccessToken struct {
	ID     string    `json:"id"`
	Notes  string    `json:"notes"`
	Expire time.Time `json:"expire"`
	// Token field will only be shown when you create an access token.
	Token string `json:"token"`
}

type impl struct {
	httpCli httpClient
	TraceInterface
	UserInterface
	AuthInterface
	ApplicationInterface
	APIInterface
	ClusterInterface
	OrganizationInterface
	RegionInterface
	CanaryReleaseInterface
	CertificateInterface
	ConsumerInterface
	LogCollectionInterface
	ServiceDiscoveryInterface
}

func (i *impl) SetGlobalClusterID(id ID) {
	i.httpCli.setClusterID(id)
}

var (
	_apiPathPrefix       = "/api/v1"
	ClusterHeaderName    = "X-API7-Cloud-Gateway-Cluster-ID"
	ErrClusterIDNotExist = errors.New("cluster id not exist")
)

// NewInterface creates an Interface object.
func NewInterface(opts *Options) (Interface, error) {
	var (
		token *AccessToken
		err   error
	)

	opts.merge(DefaultOptions)

	if opts.Token != "" {
		token = &AccessToken{
			Token: opts.Token,
		}
	} else {
		token, err = configureTokenFromFile(opts.TokenPath)
	}

	if err != nil {
		return nil, errors.Wrap(err, "new interface")
	}

	idGenerator, err := NewIDGenerator()
	if err != nil {
		return nil, errors.Wrap(err, "new interface")
	}

	trace := newTracer()
	cli, err := constructHTTPClient(&httpClientConstructOptions{
		configOptions: opts,
		token:         token,
		tracer:        trace,
		idGenerator:   idGenerator,
	})
	if err != nil {
		return nil, errors.Wrap(err, "new interface")
	}

	return &impl{
		httpCli:                   cli,
		TraceInterface:            trace,
		UserInterface:             newUser(cli),
		AuthInterface:             newAuth(cli),
		ApplicationInterface:      newApplication(cli),
		APIInterface:              newAPI(cli),
		ClusterInterface:          newCluster(cli),
		OrganizationInterface:     newOrganization(cli),
		RegionInterface:           newRegion(cli),
		CanaryReleaseInterface:    newCanaryRelease(cli),
		CertificateInterface:      newCertificate(cli),
		ConsumerInterface:         newConsumer(cli),
		LogCollectionInterface:    newLogCollection(cli),
		ServiceDiscoveryInterface: newServiceDiscovery(cli),
	}, err
}
