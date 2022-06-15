.PHONY: generate
generate:
	go generate ./...

.PHONY: wire
wire:
	wire
