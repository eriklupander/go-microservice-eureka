#!/usr/bin/env bash

export GOARCH=amd64
export GOOS=linux
go build -o bin/goeureka src/github.com/eriklupander/goeureka/*.go
docker build -t vendor .
export GOARCH=amd64
export GOOS=darwin