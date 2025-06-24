.PHONY: build test clean install lint fmt vet help

BINARY_NAME=e2s
BUILD_DIR=build
PLATFORMS=linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run tests
	go test -v ./...

fmt: ## Format code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: fmt vet ## Run linting tools

build: ## Build binary for current platform
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

build-all: ## Build binaries for all platforms
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		OS=$$(echo $$platform | cut -d'/' -f1); \
		ARCH=$$(echo $$platform | cut -d'/' -f2); \
		echo "Building for $$OS/$$ARCH..."; \
		if [ "$$OS" = "windows" ]; then \
			GOOS=$$OS GOARCH=$$ARCH go build -o $(BUILD_DIR)/$(BINARY_NAME)-$$OS-$$ARCH.exe .; \
		else \
			GOOS=$$OS GOARCH=$$ARCH go build -o $(BUILD_DIR)/$(BINARY_NAME)-$$OS-$$ARCH .; \
		fi; \
	done

install: ## Install binary to GOPATH/bin
	go install .

clean: ## Clean build artifacts
	rm -rf $(BUILD_DIR)

release-prep: clean test lint build-all ## Prepare for release (clean, test, lint, build all)

.DEFAULT_GOAL := help