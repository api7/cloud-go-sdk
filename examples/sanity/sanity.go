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

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/api7/cloud-go-sdk"
)

func main() {
	sdk, err := cloud.NewInterface(&cloud.Options{
		ServerAddr: "https://api.test.api7.cloud",
		Token:      os.Args[1],
	})
	if err != nil {
		panic(err)
	}

	me, err := sdk.Me(context.Background())
	if err != nil {
		panic(err)
	}
	org, err := sdk.GetOrganization(context.Background(), me.OrgIDs[0], nil)
	if err != nil {
		panic(err)
	}
	iter, err := sdk.ListControlPlanes(context.Background(), &cloud.ResourceListOptions{
		Organization: org,
	})
	if err != nil {
		panic(err)
	}
	cp, err := iter.Next()
	if err != nil {
		panic(err)
	}

	appIter, err := sdk.ListApplications(context.Background(), &cloud.ResourceListOptions{
		ControlPlane: cp,
		Pagination: &cloud.Pagination{
			Page:     1,
			PageSize: 100,
		},
	})
	if err != nil {
		panic(err)
	}

	var i int
	for {
		app, err := appIter.Next()
		if err != nil {
			panic(err)
		}
		if app == nil {
			break
		}
		i++
		fmt.Printf("got application #%d:\n%+v\n", i, app)
	}
}
