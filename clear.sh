#!/bin/bash

set -euxo pipefail

MYSTOREPATH="$GOPATH/src/github.com/dougfort/mystore"

rm $MYSTOREPATH/02/protobuf/*.go
rm $MYSTOREPATH/03/protobuf/*.go

rm $GOPATH/bin/mystore*