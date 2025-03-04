APP_DOMAIN=./internal/app/domain
BUSINESS_DOMAIN=./internal/business/domain

SHELL := /bin/bash

gen-e2e-got:
	@read -p "Please enter the domain name you want to test: " rawName; \
	name=$$(echo "$$rawName" | awk '{print tolower(substr($$0,1,1)) tolower(substr($$0,2))}'); \
	if [ ! -d "$(APP_DOMAIN)/$$name""api" ]; then \
		echo "Error: The domain '$$name' does not exist. Aborting."; \
		exit 1; \
	else \
		httpyac ./internal/app/domain/$$name""api/z_CRUD_e2e_test.http -all ;\
		for file in ./internal/app/domain/$$name""api/test/got/*.json; do \
		  	jq . $$file | sponge $$file; \
		done; \
	fi

e2e-test:
	@read -p "Please enter the domain name you want to test: " rawName; \
	name=$$(echo "$$rawName" | awk '{print tolower(substr($$0,1,1)) tolower(substr($$0,2))}'); \
	if [ ! -d "$(APP_DOMAIN)/$$name""api" ]; then \
		echo "Error: The domain '$$name' does not exist. Aborting."; \
		exit 1; \
	else \
	  	httpyac ./internal/app/domain/$$name""api/z_CRUD_e2e_test.http -all ;\
		for file in ./internal/app/domain/$$name""api/test/got/*.json; do \
		  	jq . $$file | sponge $$file; \
			wantfile=$$(echo $$file | sed "s/got/want/g"); \
			echo "Comparing $$file with $$wantfile"; \
			diff <(jq --indent 4 'walk(if type == "object" then with_entries(select(.key | test("At$$") | not)) else . end)' $$file) \
			     <(jq --indent 4 'walk(if type == "object" then with_entries(select(.key | test("At$$") | not)) else . end)' $$wantfile) \
			     || echo "Differences found in $$file"; \
		done; \
	fi


e2e-test2:
	@read -p "Please enter the domain name you want to test: " rawName; \
	name=$$(echo "$$rawName" | awk '{print tolower(substr($$0,1,1)) tolower(substr($$0,2))}'); \
	if [ ! -d "$(APP_DOMAIN)/$$name""api" ]; then \
	echo "Error: The domain '$$name' does not exist. Aborting."; \
	exit 1; \
	else \
	httpyac ./internal/app/domain/$$name""api/z_CRUD_e2e_test.http -all ;\
	all_passed=true; \
	for file in ./internal/app/domain/$$name""api/test/got/*.json; do \
  	jq . $$file | sponge $$file; \
  	wantfile=$$(echo $$file | sed "s/got/want/g"); \
	echo "Comparing $$file with $$wantfile"; \
	if ! diff <(jq --indent 4 'walk(if type == "object" then with_entries(select(.key | test("At$$") | not)) else . end)' $$file) \
			<(jq --indent 4 'walk(if type == "object" then with_entries(select(.key | test("At$$") | not)) else . end)' $$wantfile); then \
	echo "Differences found in $$file"; \
	all_passed=false; \
	fi; \
	done; \
	if [ "$$all_passed" = true ]; then \
	echo "✅ OK"; \
	fi; \
	fi

fe-install:
	@echo "installing frontend dependencies...";\
	pnpm i;

fe-start:
	@echo "starting frontend...";\
	pnpm run start;

fe-i: fe-install

fe-s: fe-start

rpc-fail:
	@echo "fix github file too large...";\
	git config --global http.postBuffer 524288000 ;\
	echo "reset limit as 524288000.";

github-fail:
	@echo "===================================================================================================";\
	echo "1. please check terminal, if filesize issue, use make rpc-fail for github file too large...";\
	echo "===================================================================================================";