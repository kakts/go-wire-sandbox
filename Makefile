.PHONY: generate
generate: # wireのコードも生成される。
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

.PHONY: test
test:
	go test ./...

.PHONY: testv
testv:
	go test -v ./...
