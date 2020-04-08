#!/bin/bash -eux

export BINPATH=$(pwd)/build
cwd=$(pwd)

pushd $cwd/dp-frontend-renderer
  make build
  cp Dockerfile.concourse $BINPATH/
popd
