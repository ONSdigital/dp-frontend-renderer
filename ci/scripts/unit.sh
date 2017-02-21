#!/bin/bash -eux

export GOPATH=$(pwd)/go

pushd $GOPATH/src/github.com/ONSdigital/dp-frontend-renderer
  go test -tags 'production' ./...
popd
