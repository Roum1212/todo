#!make
include .env

LOCAL_BIN := $(CURDIR)/bin
PATH := $(PATH):$(LOCAL_BIN)

GOLANGCI_LINT_VERSION := v2.1.2

.golangci-lint-install:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

.PHONY: go-fmt
go-fmt: .golangci-lint-install
	$(LOCAL_BIN)/golangci-lint fmt -c .golangci.yml

.PHONY: go-lint
go-lint: .golangci-lint-install
	$(LOCAL_BIN)/golangci-lint run -c .golangci.yml ./internal/...
