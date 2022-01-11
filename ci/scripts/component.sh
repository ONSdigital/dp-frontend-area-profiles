#!/bin/bash -eux

wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list
apt-get update && apt-get -y install google-chrome-stable
pushd dp-frontend-area-profiles
  make test-component
popd
