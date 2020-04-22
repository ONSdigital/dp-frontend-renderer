#!/bin/bash -eux

export cwd=$(pwd)

pushd $cwd/dp-frontend-renderer
  make audit
popd