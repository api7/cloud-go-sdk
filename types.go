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
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Interface interface {
	Auth

	// ConfigureTokenFromFile configures an access token from the filesystem.
	ConfigureTokenFromFile(tokenPath string) error
}

// Auth is the interface for the authentication process with API7 Cloud.
type Auth interface {
	// CreateAccessToken creates a new access token. It returns a new AccessToken object which
	// fills the Token field.
	CreateAccessToken(ctx context.Context, token *AccessToken) (*AccessToken, error)
	// DeleteAccessToken deletes an access token.
	DeleteAccessToken(ctx context.Context, token *AccessToken) error
}

// AccessToken is the token used by API7 Cloud to authenticate clients.
type AccessToken struct {
	ID     string    `json:"id"`
	Notes  string    `json:"notes"`
	Expire time.Time `json:"expire"`
	// Token field will only be shown when you create an access token.
	Token string `json:"token"`
}

type client struct {
	Auth

	token *AccessToken
}

// NewInterface creates an Interface object.
func NewInterface() Interface {
	httpCli := http.DefaultClient

	return &client{
		Auth: newAuth(httpCli),
	}
}

func (client *client) ConfigureTokenFromFile(tokenPath string) error {
	var content struct {
		User struct {
			AccessToken string `yaml:"access_token"`
		} `yaml:"user"`
	}

	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(data, &content); err != nil {
		return errors.Wrap(err, "invalid token file")
	}

	if content.User.AccessToken == "" {
		return ErrEmptyToken
	}

	client.token = &AccessToken{
		Token: content.User.AccessToken,
	}
	return nil
}
