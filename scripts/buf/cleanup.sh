#!/usr/bin/env bash

cleanup() {
  # when moving contracts to other folders, delete the previous auto-generated code
  rm -rf gen/go
  rm -rf gen/python
  rm -rf third_party/OpenAPI/client
  rm -rf gen/typescript/src
  rm -rf gen/docs
}

cleanup
