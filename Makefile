PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
PROJECT_NAME = "subscriptions"

$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)

GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint
RELEASE_STR = $(shell git rev-parse --short HEAD)

.PHONY: build
build: .build-linux .build-macos

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(PROJECT_NAME)-linux-amd64 .

.PHONY: build-macos
build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/$(PROJECT_NAME)-darwin-amd64 .

.PHONY: .install-linter
.install-linter:
	@ [ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.55.2

.PHONY: lint
lint: .install-linter
	$(GOLANGCI_LINT) run ./... --config=./.golangci.yml

.PHONY: test
test:
	go test ./...

.PHONY: modules
modules:
	go mod tidy && go mod vendor

.PHONY: run
run:
	go run main.go run