# --- Конфигурация ---
LOCAL_BIN := $(CURDIR)/bin
BUF := $(LOCAL_BIN)/buf

# Версии инструментов (можно легко обновлять)
# Проверьте актуальную на github.com/bufbuild/buf/releases
BUF_VERSION := v1.49.0 
PROTOC_GEN_GO_VERSION := latest
PROTOC_GEN_GO_GRPC_VERSION := latest

# Пути
PROTO_DIR := proto
SERVER_OUT_DIR := server/pkg/pb

# Добавляем локальный bin в PATH, чтобы buf видел плагины (protoc-gen-go)
export PATH := $(LOCAL_BIN):$(PATH)
# Указываем GOBIN, чтобы go install клал бинарники к нам в папку
export GOBIN := $(LOCAL_BIN)

# OS/Arch для скачивания buf (Linux/x86_64 для Arch)
OS := $(shell uname -s)
ARCH := $(shell uname -m)

.PHONY: all bin-deps generate lint clean

all: generate

# 1. Установка всех инструментов в ./bin
bin-deps: $(BUF)
	@echo "Installing Go plugins to $(LOCAL_BIN)..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

# Правило для скачивания buf, если его нет
$(BUF):
	@echo "Downloading Buf $(BUF_VERSION)..."
	@mkdir -p $(LOCAL_BIN)
	@curl -sSL "https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$(OS)-$(ARCH)" -o $(BUF)
	@chmod +x $(BUF)

# 2. Линтинг
lint: bin-deps
	@echo "Linting proto files..."
	@$(BUF) lint

# 3. Генерация кода
generate: bin-deps lint
	@echo "Generating Go code..."
	@mkdir -p $(SERVER_OUT_DIR)
	@$(BUF) generate
	@echo "Done. Code generated in $(SERVER_OUT_DIR)"

# 4. Очистка
clean:
	rm -rf $(LOCAL_BIN)
	rm -rf $(SERVER_OUT_DIR)/*.go
