#!/bin/bash -eux

export BINPATH=$(pwd)/bin
export GOPATH=$(pwd)/go

pushd $GOPATH/src/github.com/ONSdigital/dp-frontend-renderer
  go build -tags 'production' -o $BINPATH/dp-frontend-renderer
popd
