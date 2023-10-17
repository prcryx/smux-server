include .env
export

BINARY_NAME := smux-server
.DEFAULT_GOAL := run
OS := $(OS_NAME)
ARCH := $(ARCH_NAME)

.PHONY: all build clean

build:
	@echo "building..."
	@GOARCH=$(ARCH) GOOS=$(OS) go build -o build/$(BINARY_NAME)-$(OS) cmd/app/main.go

clean:
	@go clean
	@rm -rf ./build/*

run: build
	@echo "running..."
	@./build/$(BINARY_NAME)-$(OS)