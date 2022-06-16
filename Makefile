.PHONY: generate
generate:
	go generate ./...

.PHONY: wire
wire:
	wire

.PHONY: build
build:
	go build -o main

.PHONY: run
run:
	./main
