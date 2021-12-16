.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
	golangci-lint run
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build main.go
.PHONY:build

run: build
	go run main.go
.PHONY:run