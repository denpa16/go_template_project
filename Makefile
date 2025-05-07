LOCAL_BIN:=$(CURDIR)/bin

ENV_NAME="go_template_project"

include .env

.PHONY: .build
.build:
	go build -o bin/go_template_project ./cmd/go_template_project/
	go build -o bin/go_template_project_migrate ./cmd/go_template_project_migrate/

# Скачиваем зависимости
.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@latest
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest

run:
	go run ./cmd/go_template_project/

.PHONY: .sqlc-generate
.sqlc-generate:
	./bin/sqlc -f ./sqlc/sqlc.json generate

test:
	go test ./tests/... -count=1

.PHONY: .swag-generate
.swag-generate:
	./bin/swag init -o ./api -g ./cmd/go_template_project/main.go --parseDependency  --parseInternal --parseDepth 1
	./bin/swag fmt -d ./internal/app

.PHONY: .goose-generate
.goose-generate:
	cd migrations && ../bin/goose create $(c) sql

.PHONY: .goose-up
.goose-up:
	go run ./cmd/vp_migrate/ up

.PHONY: .goose-redo
.goose-redo:
	go run ./cmd/vp_migrate/ redo

.PHONY: .goose-down
.goose-down:
	go run ./cmd/vp_migrate/ down

.PHONY: .goose-reset
.goose-reset:
	go run ./cmd/vp_migrate/ reset
