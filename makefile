APP_NAME = vehicleapi
BUILD_PATH="./cmd/vehicleapi"
BIN_DIR = bin
OUTPUT = $(BIN_DIR)/$(APP_NAME)
GOOS = linux
# GOARCH = amd64
GOARCH = arm64
CC = aarch64-linux-gnu-gcc
# CC = gcc
build:
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) CC=$(CC) \
	go build -x -v -o $(OUTPUT)-$(GOARCH) $(BUILD_PATH)
