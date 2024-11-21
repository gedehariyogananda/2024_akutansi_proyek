WIRE_DIR=./Routes/Di

.PHONY: wire
wire:
	go generate $(WIRE_DIR)


start:
	@go run main.go