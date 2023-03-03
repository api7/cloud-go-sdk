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
		ServerAddr:      "https://api.test.api7.cloud",
		Token:           os.Args[1],
		EnableHTTPTrace: true,
	})
	if err != nil {
		panic(err)
	}

	waitingForLog := make(chan struct{})
	go printLog(sdk.TraceChan(), waitingForLog)

	me, err := sdk.Me(context.Background())
	if err != nil {
		panic(err)
	}
	org, err := sdk.GetOrganization(context.Background(), me.OrgIDs[0], nil)
	if err != nil {
		panic(err)
	}
	iter, err := sdk.ListClusters(context.Background(), &cloud.ResourceListOptions{
		Organization: org,
	})
	if err != nil {
		panic(err)
	}
	cluster, err := iter.Next()
	if err != nil {
		panic(err)
	}

	appIter, err := sdk.ListApplications(context.Background(), &cloud.ResourceListOptions{
		Cluster: cluster,
		Pagination: &cloud.Pagination{
			Page:     1,
			PageSize: 100,
		},
	})
	if err != nil {
		panic(err)
	}

	lastAppID := cloud.ID(0)
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
		lastAppID = app.ID
	}

	gatewayInstances, err := sdk.ListAllGatewayInstances(context.Background(), cluster.ID, nil)
	if err != nil {
		panic(err)
	}
	for _, gw := range gatewayInstances {
		fmt.Printf("id:%s, version:%s, ip:%s\n", gw.ID, gw.Version, gw.IP)
	}

	if lastAppID != cloud.ID(0) {
		sdk.SetGlobalClusterID(cluster.ID)
		apis, _ := sdk.ListAPIs(context.Background(), &cloud.ResourceListOptions{
			Application: &cloud.Application{ID: lastAppID},
			Pagination: &cloud.Pagination{
				Page:     1,
				PageSize: 100,
			}})
		api, _ := apis.Next()
		fmt.Println(api)
	}

	waitingForLog <- struct{}{}
}

func printLog(c <-chan *cloud.TraceSeries, done chan struct{}) {
	for {
		select {
		case data := <-c:
			fmt.Println("\033[36m" + cloud.FormatTraceSeries(data) + "\033[0m")
		case <-done:
			return
		}
	}
}
