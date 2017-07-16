#!/bin/bash

export GOPATH=`pwd`
export PATH=$PATH:$GOPATH/bin

echo "GOPATH:${GOPATH}";

go get github.com/tidwall/gjson
