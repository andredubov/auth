.PHONY:
.SILENT:

LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.21.1
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.3.14

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/joho/godotenv
	go get -u github.com/jackc/pgx/v4
	go get -u github.com/Masterminds/squirrel
	go get -u github.com/andredubov/golibs

generate:
	make generate-auth-api

genearate-mocks:
	go generate ./...

test:	
	go clean -testcache
	go test ./... -covermode count -coverpkg=github.com/andredubov/auth/internal/service/...,github.com/andredubov/auth/internal/api/... -count 5

test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode count -coverpkg=github.com/andredubov/auth/internal/service/...,github.com/andredubov/auth/internal/api/... -count 5
	grep -v 'mocks\|config' coverage.tmp.out  > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out;
	go tool cover -func=./coverage.out | grep "total";
	grep -sqFx "/coverage.out" .gitignore || echo "/coverage.out" >> .gitignore

generate-auth-api:
	mkdir -p ./pkg/auth/v1
	protoc --proto_path=./api/auth/v1 --go_out=./pkg/auth/v1 \
	--go_opt=paths=source_relative --plugin=protoc-gen-go=./bin/protoc-gen-go \
	--go-grpc_out=./pkg/auth/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=./bin/protoc-gen-go-grpc \
	./api/auth/v1/auth.proto

local-docker-compose-up:
	docker compose --env-file ./config/.env stop auth-migrator
	docker compose --env-file ./config/.env stop auth-server
	docker compose --env-file ./config/.env rm -f auth-migrator
	docker compose --env-file ./config/.env rm -f auth-server
	docker compose --env-file ./config/.env build auth-migrator 
	docker compose --env-file ./config/.env build auth-server
	docker compose --env-file ./config/.env up --force-recreate -d

build:
	go build -o ./bin/auth ./cmd/auth/main.go

run: build
	./bin/auth -config-path ./config/.env