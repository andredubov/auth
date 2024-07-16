
LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	./bin/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/joho/godotenv
	go get -u github.com/jackc/pgx/v4

generate:
	make generate-auth-api

generate-auth-api:
	mkdir -p ./pkg/auth/v1
	protoc --proto_path=./api/auth/v1 --go_out=./pkg/auth/v1 \
	--go_opt=paths=source_relative --plugin=protoc-gen-go=./bin/protoc-gen-go \
	--go-grpc_out=./pkg/auth/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=./bin/protoc-gen-go-grpc \
	./api/auth/v1/auth.proto