#!/bin/bash

set -euxo pipefail

MYSTOREPATH="$GOPATH/src/github.com/dougfort/mystore/03"

protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --proto_path=$MYSTOREPATH/protobuf \
    --plugin=$GOPATH/bin/protoc-gen-grpc-gateway \
    --grpc-gateway_out=logtostderr=true:$MYSTOREPATH/protobuf \
    $MYSTOREPATH/protobuf/mystore.proto
