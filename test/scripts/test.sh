#!/bin/sh

set -ex

export ENV=testing

go test -cover -p 1 -coverprofile=c.out ./internal/domain/*
go tool cover -html=c.out -o coverage.html
