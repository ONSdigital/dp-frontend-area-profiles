#!/bin/bash -eux

pushd dp-frontend-area-profiles
  make build
  cp build/dp-frontend-area-profiles Dockerfile.concourse ../build
popd
