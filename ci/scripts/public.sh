#!/bin/bash -eux

pushd dp-frontend-area-profiles
  make public-build
  cp build/dp-frontend-area-profiles Dockerfile.concourse ../public
popd
