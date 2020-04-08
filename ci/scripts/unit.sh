#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-frontend-renderer
  make test
popd
