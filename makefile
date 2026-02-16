APP ?= vehicleapi
BIN_DIR := bin
BUILD_PATH= .

.PHONY: all build arm64 darwin-arm64 armv7 clean
all: build

# Native build for current machine
build:
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=1 go build -trimpath -o $(BIN_DIR)/$(APP) $(BUILD_PATH)

# Linux ARM64 (aarch64)
arm64:
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 \
	go build -trimpath -o $(BIN_DIR)/$(APP)-arm64 .

# Linux ARMv7 (32-bit)
armv7:
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 \
	go build -trimpath -o $(BIN_DIR)/$(APP)-armv7 .

# macOS ARM64 (Apple Silicon)
darwin-arm64:
	mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 \
	go build -trimpath -o $(BIN_DIR)/$(APP)-darwin-arm64 .

clean:
	rm -rf $(BIN_DIR)


# APP_NAME = vehicleapi
# BUILD_PATH="./cmd/vehicleapi"
# BIN_DIR = bin
# OUTPUT = $(BIN_DIR)/$(APP_NAME)
# GOOS = linux
# # GOARCH = amd64
# GOARCH = arm64
# CC = aarch64-linux-gnu-gcc
# # CC = gcc
# build:
# 	mkdir -p $(BIN_DIR)
# 	CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) CC=$(CC) \
# 	go build -x -v -o $(OUTPUT)-$(GOARCH) $(BUILD_PATH)
