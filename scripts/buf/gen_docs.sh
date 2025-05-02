#!/usr/bin/env bash
set \
  -o errexit \
  -o nounset \
  -o pipefail

gen_docs() {
  for path in "${paths[@]}"; do
    buf generate --path "$path" --template buf.gen.docs.yaml
  done
}

gen_docs
