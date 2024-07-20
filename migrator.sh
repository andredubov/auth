#!/bin/bash

MIGRATION_DSN="host=${POSTGRES_HOST} port=${POSTGRES_PORT} dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=${POSTGRES_SSL_MODE}"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "$MIGRATION_DSN" up -v