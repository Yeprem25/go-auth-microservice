include .env
export

COMPOSE    = docker compose
MIGRATIONS = api-gateway/migrations

.PHONY: db-up db-down migrate-new migrate-up migrate-down run

db-up:
	$(COMPOSE) up -d

db-down:
	$(COMPOSE) down

migrate-new:
	atlas migrate diff $(name) --env local

migrate-up:
	migrate -path $(MIGRATIONS) -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS) -database "$(DATABASE_URL)" down 1

run:
	go run ./api-gateway/cmd