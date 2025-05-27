APP_NAME = vehicleapi
BUILD_PATH="./cmd/vehicleapi"
OUTPUT = $(APP_NAME)
GOOS = linux
GOARCH = arm64
CC = aarch64-linux-gnu-gcc

build:
	CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) CC=$(CC) \
	go build -x -v -o $(OUTPUT)-$(GOARCH) $(BUILD_PATH)