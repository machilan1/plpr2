# ==============================================================================
# Convenience commands

dev-reupall: dev-down dev-up dev-logs

dev-revisecode: dev-update dev-logs

## dev-enterdb: enter the database container
dev-enterdb:
	@echo 'Entering the database container...'
	docker exec -it ${PROJECT_NAME}-db-1 psql -U ${DB_USER} -d ${DB_NAME}

## newfile: make new file, path=$1,name=$2
.PHONY: newfile
newfile:
	@echo 'creating new file: ${path}${name}...'
	mkdir -p $(patsubst %\,%,$(path))
	touch $(patsubst %\,%,$(path))/${name}

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

## new-domain: make new domain, name=$1
new-domain:
	@read -p "Please enter the domain name you want to add: " name;\
	if [ -d "./app/domain/$$name" ] || [ -d "./business/domain/$$name" ] || [ -d "./business/domain/$$name/stores/$$name""db" ]; then \
		echo "Error: The domain '$$name' already exists. Aborting."; \
		exit 1; \
	else \
		mkdir ./app/domain/$$name""api ;\
		mkdir -p "./business/domain/$$name/stores/""$$name""db" ;\
		echo "package $$name""api" > ./app/domain/$$name""api/model.go;\
		echo "package $$name""api" > ./app/domain/$$name""api/route.go;\
		echo "package $$name""api" > ./app/domain/$$name""api/$$name""api.go;\
		echo "package $$name" > ./business/domain/$$name/model.go;\
		echo "package $$name" > ./business/domain/$$name/name.go;\
		echo "package $$name""db" > ./business/domain/$$name/stores/$$name""db/model.go;\
	fi

dev-ru: dev-reupall
dev-rv: dev-revisecode


t-ser:
	@echo 'Starting the test server...'
	go test -count=1 -c -o testserver ./cmd/api;\
	SEED_DISABLED=false ./testserver -test.v -test.run TestServer;

#dev-resetdb:
#	docker exec -it ${PROJECT_NAME}-db-1 psql -U ${DB_USER} -d ${DB_NAME} -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
#	docker compose -f ./build/docker-compose.yaml up -d --build --no-deps migrate