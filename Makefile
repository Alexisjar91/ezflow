db:
	docker compose exec db psql -U admin -d ezflow

migrate:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir sql/migrations postgres "postgres://admin:1234@localhost:5433/ezflow?sslmode=disable" up

migrate-down:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir sql/migrations postgres "postgres://admin:1234@localhost:5433/ezflow?sslmode=disable" down

migrate-reset:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir sql/migrations postgres "postgres://admin:1234@localhost:5433/ezflow?sslmode=disable" reset

migrate-create:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir sql/migrations create $(name) sql

db-reset:
	docker compose down -v && docker compose up db -d

build-app:
	docker compose build app

run-app:
	docker compose run app

logs-app:
	docker compose logs -f app

logs-db:
	docker compose logs -f db

