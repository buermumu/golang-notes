#!/bin/bash

export GOPATH=`pwd`
export PATH=$PATH:$GOPATH/bin

echo "GOPATH:${GOPATH}";

gofmt -w src
go install main

echo "install complete";

