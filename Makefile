test:
	go test -race -cover ./...
.PHONY: test

test-verbose:
	go test -v -race -cove ./...
.PHONY: test-verbose

run:
	go run cmd/gameoflife/main.go
.PHONY: run

build:
	go build -o bin/gameoflife cmd/gameoflife/main.go
.PHONY: build
