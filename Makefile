SHELL := /bin/bash
COMPOSE := docker compose

.PHONY: help db-up db-down dev api stub web build test contract smoke seed psql up down logs fmt tidy clean

help:
	@grep -E '^[a-z-]+:.*?## ' $(MAKEFILE_LIST) | sed 's/:.*## /\t/' | expand -t22

db-up: ## start Postgres only (the dev loop runs Go natively)
	$(COMPOSE) up -d db
	@until $(COMPOSE) exec -T db pg_isready -U app -d market >/dev/null 2>&1; do sleep 1; done
	@echo "postgres ready on 5455"

db-down: ## stop Postgres
	$(COMPOSE) stop db

dev: db-up ## run stub + api + vite locally, all in the foreground
	@trap 'kill 0' EXIT; \
	go run ./apps/stubgw & \
	go run ./apps/api & \
	(cd apps/web && npm run dev) & \
	wait

api: ## run the backend alone
	go run ./apps/api

stub: ## run the stub gateway alone
	go run ./apps/stubgw

web: ## run the vite dev server alone
	cd apps/web && npm run dev

build: ## build both binaries with the web app embedded
	cd apps/web && npm ci && npm run build
	mkdir -p bin
	go build -o bin/api ./apps/api
	go build -o bin/stubgw ./apps/stubgw

test: ## run every Go test
	go test ./...

contract: ## run only the doc-driven contract tests
	DOCS_DIR=./1688-api-docs go test ./internal/ali/... ./internal/stub/... -run 'Modelled|RequestParams|Registry|Samples|ErrorCodes|Topics|Documented' -v

smoke: ## signed round trip against the stub
	go run ./apps/api -smoke

seed: ## import the stub catalogue into our database
	curl -fsS -XPOST localhost:8787/api/admin/import \
	  -H "Authorization: Bearer $${ADMIN_TOKEN:-dev}" \
	  -H 'Content-Type: application/json' \
	  -d '{"all":true}' | jq .

psql: ## interactive psql (no local client needed)
	$(COMPOSE) exec db psql -U app -d market

up: ## everything in containers
	$(COMPOSE) up -d --build

down: ## stop everything
	$(COMPOSE) down

logs: ## tail container logs
	$(COMPOSE) logs -f --tail=100

fmt: ## gofmt the tree
	gofmt -w ./apps ./internal

tidy: ## resolve Go dependencies
	go mod tidy

clean: ## remove build output and local stub state
	rm -rf bin apps/web/dist .stub-state.json
