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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigureTokenFromFile(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		fileContent   string
		expectedToken string
		expectedError string
	}{
		{
			name: "valid token",
			fileContent: `
user:
  access_token: This is my token
`,
			expectedToken: "This is my token",
		},
		{
			name: "empty token",
			fileContent: `
user:
  access_token: ""
`,
			expectedError: ErrEmptyToken.Error(),
		},
		{
			name:          "invalid token file",
			fileContent:   "bad data",
			expectedError: "invalid token file",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			file, err := os.CreateTemp("./", "credentials-*")
			assert.Nil(t, err, "create temp file")

			defer func() {
				err = os.Remove(file.Name())
				assert.Nilf(t, err, "delete temp file: %s", file.Name())
			}()

			_, err = file.WriteString(tc.fileContent)
			assert.Nil(t, err, "write file: %s, content: %s", file.Name(), tc.fileContent)

			iface := NewInterface()
			err = iface.ConfigureTokenFromFile(file.Name())
			if tc.expectedError != "" {
				assert.Contains(t, err.Error(), tc.expectedError, "check the token configuration error")
			} else {
				assert.Nil(t, err, "check if the toke configuration is OK")
				cli := iface.(*client)
				assert.Equal(t, tc.expectedToken, cli.token.Token, "check token value")
			}
		})
	}
}
