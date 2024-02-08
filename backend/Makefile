LOCAL_BIN=$(CURDIR)/bin

# Скачивание библиотек
download-bin:
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@v1.29.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2

# Сборка
build:
	make clean
	make install-bin
	make generate
	make bin

# Запуск бинарника
run:
	bin/grpc-server

# Проверка кода
lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.yaml

# Генерация кода
generate:
	$(LOCAL_BIN)/buf generate

# Создание бинарника проекта
.PHONY: bin
bin:
	GOOS=linux GOARCH=amd64 go build -o bin/grpc-server cmd/grpc_server/main.go

clean:
	rm -rf $(LOCAL_BIN)
	rm -rf pkg/microservice
	rm -rf docs
