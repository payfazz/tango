#!/bin/sh

set -ex

go run cmd/test/main.go

go test -cover -coverprofile=c.out ./internal/domain/*
go tool cover -html=c.out -o coverage.html
