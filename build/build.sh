#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/JiaoDean/crd-controller
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`

rm -rf build/crd-controller

export GOARCH="amd64"
export GOOS="linux"

GIT_HASH=`git rev-parse --short HEAD || echo "HEAD"`
GIT_BRANCH=`git symbolic-ref --short -q HEAD`
BUILD_TIME=`date +%FT%T%z`
GIT_TAG=`git describe --tag --long | awk -F '-' '{print $1}'`

cd ${GOPATH}/src;${GOPATH}/src/k8s.io/code-generator/generate-groups.sh  all github.com/JiaoDean/crd-controller/pkg/client github.com/JiaoDean/crd-controller/pkg/apis crd:v1beta1 --output-base "${GOPATH}/src";cd -;
CGO_ENABLED=0 go build -ldflags "-X main.BRANCH=${GIT_BRANCH} -X main.REVISION=${GIT_HASH} -X main.TAG=${GIT_TAG} -X main.BUILDTIME=${BUILD_TIME}" -o ./build/crd-controller

cd ${GOPATH}/src/github.com/JiaoDean/crd-controller/build/

if [[ "$1" == "" ]]; then
  docker build -t=registry:crd-controller-${GIT_HASH} ./
  docker push registry:crd-controller-${GIT_HASH}
fi
