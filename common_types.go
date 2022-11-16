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
	"encoding/json"
	"strconv"
)

// ID is the type of the id field used for any entities
type ID uint64

// String indicates how to convert ID to a string.
func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// MarshalJSON is the way to encode ID to JSON string.
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(id), 10))
}

// UnmarshalJSON is the way to decode ID from JSON string.
func (id *ID) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch v := value.(type) {
	case string:
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		*id = ID(u)
	default:
		panic("unknown type")
	}
	return nil
}

const (
	// Any means any status
	Any = EntityStatus(-1)
	// Uninitialized represents the entity has been saved to the db, but the associated resource has not yet been ready
	Uninitialized = EntityStatus(0)
	// Normal indicates that the entity and associated resources are ready
	Normal = EntityStatus(50)
	// Deleted indicates the entity has been deleted
	Deleted = EntityStatus(100)
)

const (
	// ProtocolHTTP indicates the HTTP protocol.
	ProtocolHTTP = "HTTP"
	// ProtocolHTTPS indicates the HTTPS protocol.
	ProtocolHTTPS = "HTTPS"
)

const (
	// ActiveStatus indicates an object is active, and this object
	// will be seen by gateway instances.
	ActiveStatus = iota
	// InactiveStatus indicates an object is inactive, and this object
	// won't be seen by gateway instances.
	InactiveStatus
)

// EntityStatus is common status definition for any kind of entity:
// * Uninitialized represents the entity has been saved to the db, but the associated resource has not yet been ready.
// * Normal indicates that the entity and associated resources are ready.
// * Deleted indicates the entity has been deleted.
type EntityStatus int
