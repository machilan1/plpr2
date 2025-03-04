ENV_FILE := .env
-include $(ENV_FILE)

sinclude ./scripts/foundation/build.mk
sinclude ./scripts/foundation/rebase.mk
sinclude ./scripts/foundation/undertesting.mk
sinclude ./scripts/project/project.mk

# ==============================================================================
# Define environment variables

# get current directory without the full path
CURRENT_DIR := $(notdir $(patsubst %/,%,$(CURDIR)))
export COMPOSE_PROJECT_NAME := $(CURRENT_DIR)

# ==============================================================================
# Install dependencies

dev-brew:
	brew update
	brew list pre-commit || brew install pre-commit
	pre-commit install

dev-gotooling: dev-gotooling-dev dev-gotooling-ci

dev-gotooling-dev:
	go install github.com/divan/expvarmon@latest
	go install github.com/air-verse/air@latest
	go install github.com/open-policy-agent/opa@latest
	go install mvdan.cc/gofumpt@latest
	go install golang.org/x/tools/cmd/goimports@latest

dev-gotooling-ci:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all

# ==============================================================================

# lint uses the same linter as CI and tries to report the same results running
# locally. There is a chance that CI detects linter errors that are not found
# locally, but it should be rare.
lint:
	golangci-lint run --config .golangci.yaml

vuln-check:
	govulncheck ./...

fmt:
	go fmt ./...
	goimports -l -w cmd internal
	gofumpt -l -w cmd internal

# diff-check runs git-diff and fails if there are any changes.
.PHONY: diff-check
diff-check:
	@FINDINGS="$$(git status -s -uall)" ; \
		if [ -n "$${FINDINGS}" ]; then \
			echo "Changed files:\n\n" ; \
			echo "$${FINDINGS}\n\n" ; \
			echo "Diffs:\n\n" ; \
			git diff ; \
			git diff --cached ; \
			exit 1 ; \
		fi

.PHONY: generate
generate:
	@go generate ./...

.PHONY: generate-check
generate-check: generate diff-check

# ==============================================================================
# Testing

test:
	go test -count=1 -shuffle=on -timeout=5m ./...

test-acc:
	go test -count=1 -shuffle=on -timeout=10m -race ./... -coverprofile=coverage.out

test-coverage:
	go tool cover -func=./coverage.out

# ==============================================================================
# Running from within docker

dev-up:
	docker compose -f ./build/docker-compose.yaml up -d --build --remove-orphans

dev-down:
	docker compose down -v

# ==============================================================================
# Build containers

dev-build: dev-build-api dev-build-migrate dev-build-seed

dev-build-api:
	docker compose -f ./build/docker-compose.yaml build api

dev-build-migrate:
	docker compose -f ./build/docker-compose.yaml build migrate

dev-build-seed:
	docker compose -f ./build/docker-compose.yaml build seed

# ==============================================================================
# Build & restart containers

dev-update: dev-update-api

dev-update-api:
	docker compose -f ./build/docker-compose.yaml up -d --build --no-deps api

# ==============================================================================
# Logs

dev-logs:
	docker compose logs api -f --no-log-prefix --tail=100 | go run cmd/logfmt/main.go

# ==============================================================================
# Build binaries

.PHONY: build-go
build-go:
	go build -o ./bin/api ./cmd/api
	go build -o ./bin/migrate ./cmd/migrate
	go build -o ./bin/seed ./cmd/seed

# ==============================================================================
# Code generation

codegen:
	@read -p "Please enter the domain name you want to add: " name; \
	if [ -d "./internal/app/domain/$$name" ] || [ -d "./internal/business/domain/$$name" ] || [ -d "./internal/business/domain/$$name/stores/$$name""db" ]; then \
		echo "Error: The domain '$$name' already exists. Aborting."; \
		exit 1; \
	else \
		read -p "Please enter the abbreviation name you want to add: " abbr; \
		read -p "Please enter the plural you want to add: " plur; \
		read -p "Please choose whether to use pagination (Y/n): " pag; \
		go run ./cmd/codegen $$name $$abbr $$plur $$pag; \
	fi

codegen2:
	@read -p "Please enter the domain name you want to add: " name; \
	if [ -d "./internal/app/domain/$$name" ] || [ -d "./internal/business/domain/$$name" ] || [ -d "./internal/business/domain/$$name/stores/$$name""db" ]; then \
		echo "Error: The domain '$$name' already exists. Aborting."; \
		exit 1; \
	else \
		read -p "Please enter the abbreviation name you want to add: " abbr; \
		read -p "Please enter the plural you want to add: " plur; \
		go run ./cmd/codegen $$name $$abbr $$plur q n; \
	fi

# ======================================================================================================================

dev-migrate-up:
	docker compose -f ./build/docker-compose.yaml run --build --no-deps migrate | go run cmd/logfmt/main.go

dev-migrate-down:
	@read -p "Please enter the number of migrations you want to rollback: " num; \
	docker compose -f ./build/docker-compose.yaml run --build --no-deps -e MIGRATION_DOWN=true -e MIGRATION_VERSION=$$num migrate | go run cmd/logfmt/main.go

dev-migrate-down-all:
	docker compose -f ./build/docker-compose.yaml run --build --no-deps -e MIGRATION_DOWN=true -e MIGRATION_VERSION=0 migrate | go run cmd/logfmt/main.go

dev-run-seed:
	@read -p "Please enter the version of the seed you want to run (leave empty to run all): " version; \
	if [ -z "$$version" ]; then \
		docker compose -f ./build/docker-compose.yaml run --build --no-deps seed | go run cmd/logfmt/main.go; \
	else \
		docker compose -f ./build/docker-compose.yaml run --build --no-deps -e SEED_VERSION=$$version seed | go run cmd/logfmt/main.go; \
	fi

dev-run-seed-all:
	docker compose -f ./build/docker-compose.yaml run --build --no-deps seed | go run cmd/logfmt/main.go;

dev-resetdb:
	-@$(MAKE) dev-migrate-down-all
	-@$(MAKE) dev-migrate-up
	-@$(MAKE) dev-run-seed-all