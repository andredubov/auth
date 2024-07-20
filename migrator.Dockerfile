FROM alpine:3.13

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.21.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

COPY migrations/*.sql ./migrations/
COPY migrator.sh .

RUN chmod +x migrator.sh

ENTRYPOINT ["bash", "migrator.sh"]