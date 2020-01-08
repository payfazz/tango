#!/bin/sh

set -ex

go test -cover -coverprofile=c.out ./internal/domain/*
go tool cover -html=c.out -o coverage.html
