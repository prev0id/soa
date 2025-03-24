.PHONY: all build up down clean migrate-up

all: build up

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down -v

migrate-up:
	goose -dir ./user/migrations postgres "host=localhost port=5432 user=user password=password dbname=users_db sslmode=disable" up