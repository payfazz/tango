FROM debian

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

RUN set -eux \
 && mkdir /logs \
 && ln -sf /dev/stdout /logs/out.log \
 && ln -sf /dev/stderr /logs/err.log

COPY app .

CMD exec ./app 1>>/logs/out.log 2>>/logs/err.log
