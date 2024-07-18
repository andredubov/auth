FROM alpine:3.13

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.21.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

COPY migrations/*.sql ./migrations/
COPY migration_local.sh .
COPY local.env .

RUN chmod +x migration_local.sh

ENTRYPOINT ["bash", "migration_local.sh"]