APP ?= vehicleapi
BIN_DIR := bin
BUILD_PATH := .
export CGO_CFLAGS := -std=gnu11 -Wno-error=discarded-qualifiers

.PHONY: all build linux-amd64 arm64 armv7 darwin-arm64 clean
all: build

build:
	# go clean -cache
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=1 \
	go build -trimpath -o $(BIN_DIR)/$(APP) $(BUILD_PATH)

# Linux AMD64
amd64:
	# go clean -cache
	mkdir -p $(BIN_DIR)
	rm -f $(BIN_DIR)/$(APP)-arm64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=gcc \
	go build -trimpath -x -o $(BIN_DIR)/$(APP)-amd64 $(BUILD_PATH)

# Linux ARM64 (aarch64)
arm64:
	# go clean -cache
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc \
	go build -trimpath -o $(BIN_DIR)/$(APP)-arm64 $(BUILD_PATH)

# Linux ARMv7 (32-bit)
# Requires an arm-linux-gnueabihf cross-compiler installed.
armv7:
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CC=arm-linux-gnueabihf-gcc \
	go build -trimpath -o $(BIN_DIR)/$(APP)-armv7 $(BUILD_PATH)

# macOS ARM64 (Apple Silicon)
# Cross-compiling CGO to darwin requires an osxcross toolchain; this won't work with plain gcc on Linux.
# for osxcross: set CC=o64-clang
darwin-arm64:
	mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 CC=o64-clang \
	go build -trimpath -o $(BIN_DIR)/$(APP)-darwin-arm64 $(BUILD_PATH)

clean:
	rm -rf $(BIN_DIR)
