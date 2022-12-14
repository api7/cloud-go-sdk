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
)

// AuthInterface is the interface for the authentication process with API7 Cloud.
type AuthInterface interface {
	// CreateAccessToken creates a new access token. It returns a new AccessToken object which
	// fills the Token field.
	CreateAccessToken(ctx context.Context, token *AccessToken) (*AccessToken, error)
	// DeleteAccessToken deletes an access token.
	DeleteAccessToken(ctx context.Context, token *AccessToken) error
}

type auth struct {
	client httpClient
}

func newAuth(client httpClient) AuthInterface {
	return &auth{client: client}
}

func (auth *auth) CreateAccessToken(ctx context.Context, token *AccessToken) (*AccessToken, error) {
	return nil, nil
}

func (auth *auth) DeleteAccessToken(ctx context.Context, token *AccessToken) error {
	return nil
}
