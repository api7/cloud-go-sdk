#!/usr/bin/env bash
# Copyright 2022 API7.ai, Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

aux_files=(
  github.com/api7/cloud-go-sdk=user.go
  github.com/api7/cloud-go-sdk=http_trace.go
  github.com/api7/cloud-go-sdk=auth.go
  github.com/api7/cloud-go-sdk=application.go
  github.com/api7/cloud-go-sdk=api.go
  github.com/api7/cloud-go-sdk=consumer.go
  github.com/api7/cloud-go-sdk=certificate.go
  github.com/api7/cloud-go-sdk=cluster.go
  github.com/api7/cloud-go-sdk=organization.go
  github.com/api7/cloud-go-sdk=region.go
  github.com/api7/cloud-go-sdk=canary_release.go
  github.com/api7/cloud-go-sdk=log_collection.go
  github.com/api7/cloud-go-sdk=service_discovery.go
  github.com/api7/cloud-go-sdk=store.go
)

elems=${aux_files[*]}

mockgen -write_package_comment=false \
  -source=./http.go \
  -self_package=github.com/api7/cloud-go-sdk \
  -package=cloud > ./http_mock.go

mockgen -write_package_comment=false \
  -source=./types.go \
  -aux_files ${elems// /,} \
  -package=cloud \
  -self_package=github.com/api7/cloud-go-sdk \
  > ./types_mock.go
