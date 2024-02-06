LOCAL_BIN=$(CURDIR)/bin

# Установка библиотек
install-bin:
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@v1.29.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1

# Сборка
build:
	make clean
	make install-bin
	make generate
	make bin

# Запуск бинарника
run:
	bin/grpc_server

# Генерация кода
generate:
	bin/buf generate

# Создание бинарника
bin:
	go build -o bin/grpc_server cmd/grpc_server/main.go

clean:
	rm -rf $(LOCAL_BIN)
	rm -rf pkg/microservice
	rm -rf docs
