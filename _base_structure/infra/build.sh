#!/bin/sh

mkdir tmp || true
go mod tidy
go mod vendor
go build -o main