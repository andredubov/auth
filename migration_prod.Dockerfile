FROM alpine:3.13

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root/

ADD migrations ./migrations
COPY migration_prod.sh .
COPY prod.env .

RUN chmod +x migration_prod.sh

ENTRYPOINT ["bash", "migration_prod.sh"]