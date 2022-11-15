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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDMarshalJSON(t *testing.T) {
	data, err := ID(13).MarshalJSON()
	assert.Nil(t, err, "check marshaling error")
	assert.Equal(t, "\"13\"", string(data), "check marshaling result")
}

func TestIDUnmarshalJSON(t *testing.T) {
	var id ID
	err := (&id).UnmarshalJSON([]byte("\"13\""))
	assert.Nil(t, err, "check unmarshalling error")
	assert.Equal(t, 13, int(id), "check unmarshalling result")

	err = (&id).UnmarshalJSON([]byte("\"13a\""))
	assert.NotNil(t, err, "check unmarshalling error")
}
