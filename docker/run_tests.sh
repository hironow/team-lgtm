#!/bin/bash
set -ex

if [[ "$#" -eq 0 ]]; then
  TEST_DIR='/workspace'
else
  TEST_DIR=${1}
fi

cd ${TEST_DIR}

if [[ -f go.mod ]]; then
    GO111MODULE=on
    go mod download
fi

go test ./...