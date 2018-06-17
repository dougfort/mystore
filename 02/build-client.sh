#!/bin/bash

set -euxo pipefail

MYSTOREPATH="$GOPATH/src/github.com/dougfort/mystore/02"

pushd $MYSTOREPATH/mystore-client
go install --race
popd