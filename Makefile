include .env.example

GOOSE_DBSTRING=${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Local

run-db:
	docker-compose up -d mysql

install-goose:
	@which goose > /dev/null || go install github.com/pressly/goose/v3/cmd/goose@v3.18.0

goose-add-migration: install-goose
ifeq ($(name),)
	@echo "Run this command with the migration file name."
	@echo "Usage:"
	@echo "	$$ make goose-add-migration name=<name>"
else
	goose -dir ./migrations create ${name} go
endif

goose-build:
	cd ./migrations && go build -o ./goose-migrations *.go

goose-up: goose-build
	cd ./migrations && ./goose-migrations -dir ./ "${GOOSE_DBSTRING}&multiStatements=true" up

goose-down: goose-build
	cd ./migrations && ./goose-migrations -dir ./ "${GOOSE_DBSTRING}&multiStatements=true" down