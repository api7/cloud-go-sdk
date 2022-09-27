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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type auth struct {
	client *http.Client
}

type accessTokenPayload struct {
	Expire int64  `json:"expire"`
	Notes  string `json:"notes"`
}

type response struct {
	Payload struct {
		ID          string `json:"id"`
		AccessToken string `json:"access_token"`
		Notes       string `json:"notes"`
	} `json:"payload"`
}

func newAuth(client *http.Client) Auth {

	return &auth{
		client: client,
	}
}

func (auth *auth) CreateAccessToken(ctx context.Context, token *AccessToken) (*AccessToken, error) {

	requestPayload := accessTokenPayload{
		Expire: token.Expire.Unix(),
		Notes:  token.Notes,
	}

	json_data, err := json.Marshal(requestPayload)
	fmt.Println(bytes.NewBuffer(json_data))
	resp, err := auth.client.Post(globalAPIBaseUrl+"/user/access_tokens", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	var result response
	fmt.Println("result", &result)
	json.Unmarshal(body, &result)

	var returnData = &AccessToken{
		ID:    result.Payload.ID,
		Token: result.Payload.AccessToken,
		Notes: result.Payload.Notes,
	}
	return returnData, nil
}

func (auth *auth) DeleteAccessToken(ctx context.Context, token *AccessToken) error {
	return nil
}
