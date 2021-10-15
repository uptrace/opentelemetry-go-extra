#!/bin/bash

set -e

PACKAGE_DIRS=$(find . -mindepth 2 -type f -name 'go.mod' -exec dirname {} \; \
  | sed 's/^\.\///' \
  | sort)

for dir in $PACKAGE_DIRS
do
    (
        echo "testing ${dir}..."
        cd $dir
        go test ./...
	     go test ./... -short -race
	     go test ./... -run=NONE -bench=.
	     go vet ./...
        golangci-lint run
    )
done
