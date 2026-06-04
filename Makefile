.PHONY: run build tidy test fmt

run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api

tidy:
	go mod tidy

test:
	go test ./...

fmt:
	go fmt ./...
