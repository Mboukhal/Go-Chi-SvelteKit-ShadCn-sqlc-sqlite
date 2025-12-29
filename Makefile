.PHONY: server ui dev test gen install clean build_ui build_server build_prod build start docker_start docker_clean start_services stop_services clean_services db db_push db_down migrate

# Include .env file
include .env
export


all: dev

# -------- Development --------
server:
	## Start Air for hot reload
	APP_ENV=development air

ui:
	## Start Bun development server for UI
	cd ./cmd/ui && bun run dev && cd -

dev: gen
	## Start development server with hot reload
	@$(MAKE) -j ui server

test:
	## Run tests with coverage
	find ./apps -name '*_test.go' | xargs -n1 -I{} sh -c 'dir=$$(dirname "{}"); go test "$$dir" -v'

# 	rm -rf internal/db
gen:
	## Generate database code with sqlc
	sqlc generate

# -------- Installations --------
install:
	## Install server and ui dependencies
	go mod tidy
	cd ./cmd/ui && bun install && cd -

# -------- Builds --------
clean:
	## Clean up build artifacts
	rm -rf dist
	rm -rf cmd/ui/build/*

build_ui:
	## Build the UI for production
	cd ./cmd/ui && bun run build && cd -

build_server:
	## Build the server for production
	APP_ENV=production go build -o dist/server ./cmd/main.go

build: install build_ui build_server
	## Build the production server and ui

build_prod:
	## Build docker image for production
	docker build -t factorybase:latest .

start:
	## Start the production server
	APP_ENV=production ./dist/server


# -------- Database --------
db:
	## Create a new database
	goose -dir internal/adapter/db/schema create $(ARGS) sql

push:
	## Push the database to the latest version
	goose -dir internal/adapter/db/schema sqlite3 $${DATABASE_URL} up

down:
	## Migrate the database down by one version
	goose -dir internal/adapter/db/schema sqlite3 $${DATABASE_URL} down 1

migrate:
	## Migrate the database up to the latest version
	goose -dir internal/adapter/db/schema sqlite3 $${DATABASE_URL} $(ARGS)
