#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]:-$0}"; )" &> /dev/null && pwd 2> /dev/null; )";

pushd ${DIR}/../
bash vendor/k8s.io/code-generator/generate-groups.sh all \
  example.com/m/pkg/generated \
  example.com/m/pkg/apis \
  foo:v1 \
  --go-header-file ${DIR}/boilerplate.go.txt \
  -v5 \
