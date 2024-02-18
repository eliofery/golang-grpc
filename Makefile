include .env

# Settings
.DEFAULT_GOAL=help
.PHONY: bin help

# Environments
LOCAL_BIN=$(CURDIR)/bin
LOCAL_MIGRATION_DIR=$(CURDIR)/migrations/sql
LOCAL_MIGRATION_DSN="postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@127.0.0.1:$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable"

ifeq ($(OS), Windows_NT)
	SHELL=powershell.exe
	SHELL_VERSION=$(shell (Get-Host | Select-Object Version | Format-Table -HideTableHeaders | Out-String).Trim())
	OS=$(shell "{0} {1}" -f "windows", (Get-ComputerInfo -Property OsVersion, OsArchitecture | Format-Table -HideTableHeaders | Out-String).Trim())
	PACKAGE=$(shell (Get-Content go.mod -head 1).Split(" ")[1])
	HELP_CMD=Select-String "^[a-zA-Z_-]+:.*?\#\# .*$$" "./Makefile" | Foreach-Object { $$_data = $$_.matches -split ":.*?\#\# "; $$obj = New-Object PSCustomObject; Add-Member -InputObject $$obj -NotePropertyName ('Command') -NotePropertyValue $$_data[0]; Add-Member -InputObject $$obj -NotePropertyName ('Description') -NotePropertyValue $$_data[1]; $$obj } | Format-Table -HideTableHeaders @{Expression={ $$e = [char]27; "$$e[36m$$($$_.Command)$${e}[0m" }}, Description
else
	SHELL=bash
	SHELL_VERSION=$(shell echo $$BASH_VERSION)
	UNAME=$(shell uname -s)
	VERSION_AND_ARCH=$(shell uname -rm)
	ifeq ($(UNAME),Darwin)
		OS=macos ${VERSION_AND_ARCH}
	else ifeq ($(UNAME),Linux)
		OS=$(shell cat /proc/version)
	else
		$(error OS not supported by this Makefile)
	endif
	PACKAGE=$(shell awk 'NR==1 {print $$2}' go.mod)
	HELP_CMD=grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' './Makefile' | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
endif

HELP_PROJECT=$(LOCAL_BIN)/grpc-server --help

download-bin: ## Download library binaries
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@v1.29.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.18.0
	GOBIN=$(LOCAL_BIN) go install github.com/cortesi/modd/cmd/modd@v0.8

build: ## Build project
	make clean
	make install-bin
	make generate
	make bin

run: ## Run project
	$(LOCAL_BIN)/grpc-server

run-dev: ## Run project in development mode
	$(LOCAL_BIN)/modd

run-with-migration: ## Run project with migration
	$(LOCAL_BIN)/grpc-server -migration

lint: ## Checking code with a linter
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.yaml
	$(LOCAL_BIN)/buf lint

generate: ## Generating code from protobuf files
	$(LOCAL_BIN)/buf generate

bin: ## Building a binary project file for Linux
	GOOS=linux GOARCH=amd64 go build -o bin/grpc-server cmd/grpc_server/main.go

migration-create: ## Create migration
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) create example sql

migration-status: ## Dump the migration status for the current DB
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) status -v

migration-up: ## Migrate the DB to the most recent version available
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) up -v

migration-down: ## Roll back the version by 1
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) down -v

bump: ## Update packages version
	go get -u ./...

clean: ## Clear the project
	rm -rf $(LOCAL_BIN)
	rm -rf pkg/api
	rm -rf docs/api

about: ## Display info related to the build
	@echo "OS: $(OS)"
	@echo "Shell: $(SHELL) $(SHELL_VERSION)"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: $(shell awk 'NR==1 {print $$2}' go.mod)"
	@echo "Openssl version: $(shell openssl version)"

help: ## Show this help
	@echo "Usage: make [target]"
	@echo ""
	@${HELP_CMD}
	@echo ""
	@${HELP_PROJECT}
