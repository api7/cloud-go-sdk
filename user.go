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

// User defines user information for API7 Cloud.
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	OrgIDs    []ID   `json:"org_ids"`
	// TODO change the type when we need this field.
	Members json.RawMessage `json:"members"`
	// TODO change the type when we need this field.
	ProductTour json.RawMessage `json:"product_tour"`
	Connection  string          `json:"connection"`
	AvatarURL   string          `json:"avatar_url"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// UserInterface is the interface for the user-related process on API7 Cloud.
type UserInterface interface {
	// Me returns the current user's information.
	Me(ctx context.Context) (*User, error)
}

type userImpl struct {
	client httpClient
	store  StoreInterface
}

func newUser(cli httpClient, store StoreInterface) UserInterface {
	return &userImpl{
		client: cli,
		store:  store,
	}
}

func (impl *userImpl) Me(ctx context.Context) (*User, error) {
	var user User

	apiPath := path.Join(_apiPathPrefix, "/user/me")
	if err := impl.client.sendGetRequest(ctx, apiPath, "", jsonPayloadDecodeFactory(&user), map[string]string{}); err != nil {
		return nil, errors.Wrap(err, apiPath)
	}
	return &user, nil
}
