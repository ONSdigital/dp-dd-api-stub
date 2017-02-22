#!/bin/bash -eux

export GOPATH=$(pwd)/go

pushd $GOPATH/src/github.com/ONSdigital/dp-dd-api-stub
  make generate
  go test -tags 'production' ./...
popd
