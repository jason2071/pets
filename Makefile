.PHONY: run build tidy test fmt compose-up compose-down compose-force migrate-up migrate-down migrate-force migrate-create

DB_URL ?= postgres://postgres:postgres@localhost:5432/pets?sslmode=disable
MIGRATE = migrate -path migrations -database "$(DB_URL)"

compose-up:
	docker compose up -d

compose-down:
	docker compose down

compose-force:
	docker compose down -v
	docker compose up -d --build --force-recreate

# Production schema management (run with AUTO_MIGRATE=false).
migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

# Reset dirty state: make migrate-force version=1
migrate-force:
	$(MIGRATE) force $(version)

# Create new migration pair: make migrate-create name=add_owner
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api

tidy:
	go mod tidy

test:
	go test ./...

fmt:
	go fmt ./...
