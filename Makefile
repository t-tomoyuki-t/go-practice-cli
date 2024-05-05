.DEFAULT_GOAL := build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go mod tidy
	go build -o ./go-practice-cli main.go
