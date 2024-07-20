FROM golang:1.22.3-alpine AS builder

COPY . /github.com/andredubov/auth
WORKDIR /github.com/andredubov/auth

RUN go mod download && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/auth ./cmd/auth/main.go

FROM alpine:3.20

WORKDIR /root/
COPY --from=builder /github.com/andredubov/auth/bin/auth .

CMD ["./auth"]