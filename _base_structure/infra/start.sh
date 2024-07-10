#!/bin/sh

go mod tidy
go mod vendor
(./main 2>/dev/null && echo "running './main'") || (echo "running 'go run main.go'" && go run main.go)
