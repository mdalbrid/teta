.PHONY: help

PROJECTNAME=$(shell basename "$(PWD)")

build:
	go build -v ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test -v ./...

migrate_up:
	migrate -path ./internal/infra/migrations -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./internal/infra/migrations -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' down

docker_all:
	docker-compose build go-app
	docker-compose up go-app

docker_build:
	docker-compose build go-app

docker_up:
	docker-compose up -d go-app
	docker-compose ps

docker_down:
	docker-compose down
	docker-compose ps

docker_rm:
	#docker volume rm $$(docker volume ls -q)
	sudo rm -rf $(PWD)/.database

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/    /'