build:
	@go build -o bin/blocksie

run: build
	@./bin/docker

test:
	@go test -v ./...
