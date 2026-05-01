.PHONY: dev build run test lint tidy

dev:
	air

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./...

lint:
	golangci-lint run

tidy:
	go mod tidy
