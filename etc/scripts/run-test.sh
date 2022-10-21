#!/usr/bin/env sh

docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.49.0 golangci-lint run --fix ./... && \
go test -coverprofile=coverage.out -v -p 1 ./...
