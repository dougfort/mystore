#!/bin/bash

set -euxo pipefail

MYSTOREPATH="$GOPATH/src/github.com/dougfort/mystore/03"

pushd $MYSTOREPATH/gateway-proxy-server
go install --race
popd