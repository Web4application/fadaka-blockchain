#! /usr/bin/make -f

# Project variables
# Dynamically pull the version from git tags or commit hash
VERSION := $(shell git describe --tags 2>/dev/null || git describe --all)
BUILD := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
PROJECT_NAME := fadaka-blockchain
BIN_DIR := bin
ASSETS_BIN := $(BIN_DIR)/assets

# Use linker flags to provide version/build settings to the Go binary
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

# Make is verbose in Linux. Make it silent for cleaner output.
MAKEFLAGS += --silent

# Go files discovery
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

.PHONY: all build test fmt lint lint-install check fix update-auto add-token clean

# Default target
all: fmt lint test build

build:
	@echo "  >  Building $(PROJECT_NAME) to $(ASSETS_BIN)"
	@mkdir -p $(BIN_DIR)
	go build $(LDFLAGS) -o $(ASSETS_BIN) ./cmd

test:
	@echo "  >  Running unit tests"
	go test -cover -race -coverprofile=coverage.txt -covermode=atomic -v ./...

fmt:
	@echo "  >  Formatting all Go files"
	gofmt -w ${GOFMT_FILES}

lint-install:
	@if [ ! -f $(BIN_DIR)/golangci-lint ]; then \
		echo "  >  Installing golangci-lint"; \
		curl -sSfL https://raw.githubusercontent.com | sh -s -- -b $(BIN_DIR) v1.50.1; \
	fi

lint: lint-install
	@echo "  >  Running golangci-lint"
	$(BIN_DIR)/golangci-lint run --timeout=2m

# Asset Management commands
check: build
	$(ASSETS_BIN) check

fix: build
	$(ASSETS_BIN) fix

update-auto: build
	$(ASSETS_BIN) update-auto

# Helper commands for tokens
# Usage: make add-token asset_id=YOUR_ID
add-token: build
	$(ASSETS_BIN) add-token $(asset_id)

add-tokenlist: build
	$(ASSETS_BIN) add-tokenlist $(asset_id)

add-tokenlist-extended: build
	$(ASSETS_BIN) add-tokenlist-extended $(asset_id)

clean:
	@echo "  >  Cleaning build artifacts"
	rm -rf $(BIN_DIR)
	rm -f coverage.txt
