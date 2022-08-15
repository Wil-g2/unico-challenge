GOPATH ?= $(HOME)/go

run-api:
	go run main.go

run-cmd:
	go run cmd/main.go

deps:
	go mod tidy
	go mod download

format:
	go fmt ./...

test-cov:
	go test -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html

generate-swag:
	swag init -g main.go