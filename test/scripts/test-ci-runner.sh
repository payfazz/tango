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

push docker run -d --net container:$random \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=postgres \
    -e POSTGRES_DB=tango \
    postgres
push docker run -d --net container:$random \
    redis --requirepass redis

docker run --net container:$random --rm -v $REPO:/app -w /app -i golang:1.13 \
    /bin/bash ./test/scripts/test-ci-setup.sh $1 $2 || exit $?
