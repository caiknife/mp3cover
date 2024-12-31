.PHONY: build test

build:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/mp3cover ./mp3cover

test:
	go test ./...
