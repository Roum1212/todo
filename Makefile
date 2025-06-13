include .env

LOCAL_BIN := $(CURDIR)/bin
PATH := $(PATH):$(LOCAL_BIN)

APP_VERSION ?= `git describe --tags || echo "v0.0.1"`

BUF_VERSION := v1.50.1
GOLANGCI_LINT_VERSION := v2.1.2
PROTOC_GEN_GO_VERSION := v1.36.6
PROTOC_GEN_GO_GRPC_VERSION := v1.5.1
PROTOC_GEN_GRPC_GATEWAY_VERSION := v2.26.3

.buf-install:
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)

.golangci-lint-install:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

.proto-install: .buf-install
	GOBIN=$(LOCAL_BIN) go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@$(PROTOC_GEN_GRPC_GATEWAY_VERSION) \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@$(PROTOC_GEN_GRPC_GATEWAY_VERSION)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)

.vendor-proto: .vendor-proto-remove .vendor-proto-install/google/api .vendor-proto-install/protoc-gen-openapiv2/options

.vendor-proto-install/google/api:
	git clone \
		-b master \
		--depth=1 \
		--filter=tree:0 https://github.com/googleapis/googleapis vendor-proto/googleapis \
		-n \
		--single-branch && \
 	cd vendor-proto/googleapis && \
	git sparse-checkout set --no-cone google/api && \
	git checkout
	mkdir -p vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

.vendor-proto-install/protoc-gen-openapiv2/options:
	git clone \
		-b main \
		--depth=1 \
		--filter=tree:0 https://github.com/grpc-ecosystem/grpc-gateway vendor-proto/grpc-ecosystem \
		-n \
		--single-branch && \
 	cd vendor-proto/grpc-ecosystem && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p vendor-proto/protoc-gen-openapiv2
	mv vendor-proto/grpc-ecosystem/protoc-gen-openapiv2/options vendor-proto/protoc-gen-openapiv2
	rm -rf vendor-proto/grpc-ecosystem

.vendor-proto-remove:
	rm -rf vendor-proto

.PHONY: go-fmt
go-fmt: .golangci-lint-install
	$(LOCAL_BIN)/golangci-lint fmt -c .golangci.yml

.PHONY: go-lint
go-lint: .golangci-lint-install
	$(LOCAL_BIN)/golangci-lint run -c .golangci.yml ./...

.PHONY: proto-generate
proto-generate: .proto-install .vendor-proto
	$(LOCAL_BIN)/buf generate --config buf.yml --template buf.gen.yml

.PHONY: proto-lint
proto-lint: .buf-install
	$(LOCAL_BIN)/buf lint
