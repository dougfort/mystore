#!/bin/bash

set -euxo pipefail

MYSTOREPATH="$GOPATH/src/github.com/dougfort/mystore/02"

protoc --proto_path=$MYSTOREPATH/protobuf \
    --plugin=$GOPATH/bin/protoc-gen-go \
    --go_out=plugins=grpc:$MYSTOREPATH/protobuf \
    $MYSTOREPATH/protobuf/mystore.proto