#!/usr/bin/env bash
set \
  -o errexit \
  -o nounset \
  -o pipefail

docker build \
  -t github.com/fdogov/contracts/buf \
  -f build/buf/Dockerfile \
  build/buf
