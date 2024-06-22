MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
ROOT := $(dir $(MAKEFILE_PATH))

BINARY_NAME=arc-cleaner
VERSION=""
COMMIT := $(shell git rev-parse --verify HEAD)

PKG := ./...

GOPATH ?= $(HOME)/go
GOBIN ?= $(GOPATH)/bin
PATH := $(GOBIN):$(PATH)

LD_FLAGS = -ldflags " \
	-X github.com/wielewout/arc-cleaner/cmd.version=$(VERSION) \
	-X github.com/wielewout/arc-cleaner/cmd.commit=$(COMMIT) \
	"

.PHONY: all
all: help

.PHONY: help
help: ## Show help menu
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} \
		/^[a-zA-Z_-]+:.*?## / { \
			printf "    %-24s%s\n", $$1, $$2 \
		}' $(MAKEFILE_LIST)
	@echo ''

.PHONY: build
build: ## Build application
	@mkdir -p $(ROOT)/bin
	@go build $(LD_FLAGS) -o $(ROOT)/bin/$(BINARY_NAME) main.go

.PHONY: run
run: build ## Run application
	@$(ROOT)/bin/$(BINARY_NAME)

.PHONY: test
test: ## Run all tests
	@go test -v $(PKG)

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	@go test -covermode=count -coverprofile=profile.cov $(PKG) || true
	@go tool cover -func profile.cov

.PHONY: test-coverage-report
test-coverage-report: test-coverage ## Run tests with coverage and open report
	@go tool cover --html profile.cov

.PHONY: lint
lint: ## Lint code
	@golangci-lint run $(PKG)

.PHONY: clean
clean: ## Clean up artifacts
	@rm -rf $(ROOT)/bin $(ROOT)/*.cov
