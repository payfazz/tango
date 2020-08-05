#!/bin/bash

git config --global url."https://$1:$2@github.com".insteadOf "https://github.com"
go env -w GOPRIVATE=github.com/payfazz

./test/scripts/test.sh