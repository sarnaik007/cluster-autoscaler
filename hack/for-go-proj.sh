#! /bin/bash

# Copyright 2014 The Kubernetes Authors.
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
set -o pipefail
set -o nounset

CONTRIB_ROOT="$(dirname ${BASH_SOURCE})/.."
PROJECT_NAMES=(addon-resizer cluster-autoscaler vertical-pod-autoscaler)

if [[ $# -ne 1 ]]; then
  echo "missing subcommand: [build|install|test]"
  exit 1
fi

CMD="${1}"

case "${CMD}" in
  "build")
    ;;
  "install")
    ;;
  "test")
    ;;
  *)
    echo "invalid subcommand: ${CMD}"
    exit 1
    ;;
esac

for project_name in ${PROJECT_NAMES[*]}; do
  (
    project=${CONTRIB_ROOT}/${project_name}
    echo "${CMD}ing ${project}"
    cd "${project}"
    case "${CMD}" in
      "test")
        go mod vendor
        go test -race $(go list ./... | grep -v /vendor/ | grep -v vertical-pod-autoscaler/e2e)
        ;;
      *)
        godep go "${CMD}" ./...
        ;;
    esac
  )
done;

if [ "${CMD}" = "build" ] || [ "${CMD}" == "test" ]; then
  cd ${CONTRIB_ROOT}/vertical-pod-autoscaler/e2e
  go test -mod vendor -run=None ./...
fi
