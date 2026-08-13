db:
	docker compose exec db psql -U admin -d ezflow

migrate:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir sql/migrations postgres "postgres://admin:1234@localhost:5433/ezflow?sslmode=disable" up

build-app:
	docker compose build app

run-app:
	docker compose run app

logs-app:
	docker compose logs -f app

logs-db:
	docker compose logs -f db
