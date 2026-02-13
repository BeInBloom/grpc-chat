LOCAL_BIN := $(CURDIR)/bin
BUF := $(LOCAL_BIN)/buf

BUF_VERSION := v1.65.0
PROTOC_GEN_GO_VERSION := latest
PROTOC_GEN_GO_GRPC_VERSION := latest

PROTO_DIR := proto
SERVER_OUT_DIR := gen/go

export PATH := $(LOCAL_BIN):$(PATH)
export GOBIN := $(LOCAL_BIN)

OS := $(strip $(shell uname -s))
ARCH := $(strip $(shell uname -m))

.PHONY: all bin-deps generate lint lint-go lint-fix fmt clean \
	test test-auth test-chat \
	mockgen tidy build vet \
	run-auth run-chat

all: generate

bin-deps: $(BUF)
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

$(BUF):
	@mkdir -p $(LOCAL_BIN)
	@curl -sSL "https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$(OS)-$(ARCH)" -o $(BUF)
	@chmod +x $(BUF)

generate: bin-deps
	@mkdir -p $(SERVER_OUT_DIR)
	@$(BUF) generate proto --template buf.gen.yaml

lint: lint-go
	@$(BUF) lint proto

lint-go:
	@golangci-lint run ./gen/go/... ./pkg/... ./services/auth/... ./services/chat/...

lint-fix:
	@golangci-lint run ./gen/go/... ./pkg/... ./services/auth/... ./services/chat/... --fix

fmt:
	@goimports -w services/auth services/chat pkg gen/go

mockgen:
	@go generate ./services/auth/... ./services/chat/...

tidy:
	@go mod tidy -C services/auth
	@go mod tidy -C services/chat
	@go mod tidy -C pkg
	@go mod tidy -C gen/go
	@go work sync

build:
	@go build -o bin/auth ./services/auth/cmd/...
	@go build -o bin/chat ./services/chat/cmd/...

run-auth:
	@go run ./services/auth/cmd/...

run-chat:
	@go run ./services/chat/cmd/...

vet:
	@go vet ./pkg/... ./services/auth/... ./services/chat/...

clean:
	rm -rf $(LOCAL_BIN)
	rm -rf $(SERVER_OUT_DIR)

test: test-auth test-chat

test-auth:
	go -C services/auth test ./... -v

test-chat:
	go -C services/chat test ./... -v
