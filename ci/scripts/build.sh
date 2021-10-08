#!/bin/bash -eux

pushd dp-frontend-area-profiles-controller
  make build
  cp build/dp-frontend-area-profiles-controller Dockerfile.concourse ../build
popd
