up:
	docker compose up -d

up-b:
	docker compose up -d --build

down:
	docker compose down

prepare:
	docker compose exec app bash -c 'go run ./pkg/script/migration.go'