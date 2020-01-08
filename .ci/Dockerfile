FROM golang:1.13 AS builder

WORKDIR /tango

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./
RUN go build -o /app cmd/tango/main.go

FROM debian

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

RUN set -eux \
 && mkdir /logs \
 && ln -sf /dev/stdout /logs/out.log \
 && ln -sf /dev/stderr /logs/err.log

COPY --from=builder /app .
CMD exec ./app 1>>/logs/out.log 2>>/logs/err.log
