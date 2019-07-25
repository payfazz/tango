#!/bin/bash

# by https://github.com/anakaiti

DIR=$(dirname "$0")
REPO=$(cd $DIR/../../; pwd)
random=$(date +%s)

ps=""
push() { ps="$ps $($@)";}
cleanup() { docker rm -f ${ps} &> /dev/null; }
trap cleanup EXIT

push docker run -d --name $random alpine sleep 9999

# Please change -e environment to match your test environment

push docker run -d --net container:$random \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=postgres \
    -e POSTGRES_DB=tango-test \
    postgres
push docker run -d --net container:$random \
    redis --requirepass redis

image=$(docker build -q --target builder -f "$REPO/.ci/Dockerfile" "$REPO")
docker run --net container:$random --rm -i $image < $DIR/test.sh || exit $?
