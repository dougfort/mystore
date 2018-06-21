#!/bin/bash

set -euxo pipefail

MYSTOREPATH="$GOPATH/src/github.com/dougfort/mystore/01"

pushd $MYSTOREPATH/grpc-server
go install --race
popd