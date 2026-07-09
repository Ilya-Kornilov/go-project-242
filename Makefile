.PHONY: build clean test lint lint-fix

# Директория для бинарников
BIN_DIR := bin
BINARY_NAME := hexlet-path-size
CMD_PATH := ./cmd/hexlet-path-size

build:
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(BINARY_NAME) $(CMD_PATH)

clean:
	@rm -rf $(BIN_DIR)

test:
	@go test -v ./...

lint:
	@golangci-lint run ./...

lint-fix:
	@golangci-lint run --fix ./...
