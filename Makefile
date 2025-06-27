BINARY_NAME=GoLockFile

BUILD_DIR=build

CMD_DIR=./cmd

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

build-macos:
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-arm64 $(CMD_DIR)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME).exe $(CMD_DIR)

build-all: build build-macos build-windows

clean:
	rm -rf $(BUILD_DIR)

.PHONY: build build-macos build-windows build-all clean
