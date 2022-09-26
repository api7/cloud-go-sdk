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
	"fmt"
	"net/http"
)

type auth struct {
	client *http.Client
}

func newAuth(client *http.Client) Auth {
	return nil
}

func (auth *auth) CreateAccessToken(ctx context.Context, token *AccessToken) (*AccessToken, error) {
	return nil, nil
}

func (auth *auth) DeleteAccessToken(ctx context.Context, token *AccessToken) error {
	uri := fmt.Sprintf("/api/v1/user/access_tokens/%s", token.ID)
	req, err := http.NewRequest("DELETE", uri, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	if err != nil {
		return err
	}
	resp, err := auth.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return DeleteAccessTokenError
	}

	return nil
}
