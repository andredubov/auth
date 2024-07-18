FROM alpine:3.13

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.21.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

COPY migrations ./migrations
COPY migration_prod.sh .
COPY prod.env .

RUN chmod +x migration_prod.sh

ENTRYPOINT ["bash", "migration_prod.sh"]