EXECUTABLE=templater
BASE_BUILD_DIR=bin
DARWIN_AMD64=$(BASE_BUILD_DIR)/darwin_amd64/$(EXECUTABLE)
DARWIN_ARM64=$(BASE_BUILD_DIR)/darwin_arm64/$(EXECUTABLE)
LINUX_AMD64=$(BASE_BUILD_DIR)/linux_amd64/$(EXECUTABLE)
LINUX_ARM=$(BASE_BUILD_DIR)/linux_arm/$(EXECUTABLE)
LINUX_ARM64=$(BASE_BUILD_DIR)/linux_arm64/$(EXECUTABLE)
WINDOWS_AMD64=$(BASE_BUILD_DIR)/windows_amd64/$(EXECUTABLE).exe
VERSION=$(shell git describe --tags --always --long --dirty)

all: test build-all

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Testing
test:
	go test ./...

# Building
build-all: build-mac build-linux build-windows

# MacOS Builds
build-mac: build-darwin-amd64 build-darwin-arm64

build-darwin-amd64:
	env GOOS=darwin GOARCH=amd64 go build -v -o $(DARWIN_AMD64)
	chmod +x $(DARWIN_AMD64)

build-darwin-arm64:
	env GOOS=darwin GOARCH=arm64 go build -v -o $(DARWIN_ARM64)
	chmod +x $(DARWIN_ARM64)

# Linux Builds
build-linux: build-linux-amd64 build-linux-arm build-linux-arm64

build-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX_AMD64)
	chmod +x $(LINUX_AMD64)

build-linux-arm:
	env GOOS=linux GOARCH=arm go build -v -o $(LINUX_ARM)
	chmod +x $(LINUX_ARM)

build-linux-arm64:
	env GOOS=linux GOARCH=arm64 go build -v -o $(LINUX_ARM64)
	chmod +x $(LINUX_ARM64)

# Windows Builds
build-windows: build-windows-amd64

build-windows-amd64:
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS_AMD64)
	chmod +x $(WINDOWS_AMD64)
