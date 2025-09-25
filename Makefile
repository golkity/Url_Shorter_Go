LOG_DIR    = ../logs

COMPOSE_FILE := infrastructure/docker-compose.yml

swag:
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/auth/main.go -o api/docs

test:
	go test ./...

build:
	go build -v ./cmd/auth

docker:
	docker build -t $(IMAGE):latest .

db-cl:
	rm -rf postgres-data
	rm -rf redis-data

dc-up:
	docker-compose -f $(COMPOSE_FILE) up --build -d

dc-down:
	docker-compose -f $(COMPOSE_FILE) down

migration:
	docker exec -i infrastructure-postgres-1 psql \
          -U root \
          -d postgres \
          < ./db/migrations/0001_init.sql

dev: export LOG_LEVEL = debug
dev: export LOG_FILE  = $(LOG_DIR)/dev.log
dev: swag dc-up

stage: export LOG_LEVEL = warn
stage: swag dc-up

prod: export LOG_LEVEL = error
prod: swag dc-up