g-file:
	@read -p "Please enter the domain name you want to add: " rawName;\
	read -p "Please enter the domain name for replacing abbr form: " rawAbbr;\
	read -p "Please enter the domain name for replacing plural form: " rawPlural;\
	name=$$(echo "$$rawName" | awk '{print tolower(substr($$0,1,1)) tolower(substr($$0,2))}');\
	u_name=$$(echo "$$rawName" | awk '{print toupper(substr($$0,1,1)) tolower(substr($$0,2))}');\
	abbr=$$(echo "$$rawAbbr" | awk '{print tolower(substr($$0,1,1)) tolower(substr($$0,2))}');\
	u_abbr=$$(echo "$$rawAbbr" | awk '{print toupper(substr($$0,1,1)) tolower(substr($$0,2))}');\
	plural=$$(echo "$$rawPlural" | awk '{print tolower(substr($$0,1,1)) tolower(substr($$0,2))}');\
	u_plural=$$(echo "$$rawPlural" | awk '{print toupper(substr($$0,1,1)) tolower(substr($$0,2))}');\
	if [ -d "./app/domain/$$name""api" ] || [ -d "./business/domain/$$name" ]; then \
		echo "Error: The domain '$$name' already exists. Aborting.";\
		exit 1;\
	else \
		mkdir ./app/domain/$$name""api ;\
		mkdir -p "./business/domain/$$name/stores/""$$name""db" ;\
		touch ./app/domain/$$name""api/model.go;\
		touch ./app/domain/$$name""api/route.go;\
		touch ./app/domain/$$name""api/mid.go;\
		touch ./app/domain/$$name""api/$$name""api.go;\
		touch ./app/domain/$$name""api/filter.go;\
		touch ./app/domain/$$name""api/order.go;\
		touch ./business/domain/$$name/model.go;\
		touch ./business/domain/$$name/$$name.go;\
		touch ./business/domain/$$name/order.go;\
		touch ./business/domain/$$name/filter.go;\
		touch ./business/domain/$$name/stores/$$name""db/model.go;\
		touch ./business/domain/$$name/stores/$$name""db/order.go;\
		touch ./business/domain/$$name/stores/$$name""db/filter.go;\
		touch ./business/domain/$$name/stores/$$name""db/$$name""db.go;\
	fi;\
	old_domain="xxx";\
	old_u_domain="XXX";\
	old_abbr="yyy";\
	old_u_abbr="YYY";\
	old_plural="zzz";\
	old_u_plural="ZZZ";\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./app/domain/xxxapi/model.go" > ./app/domain/$$name""api/model.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./app/domain/xxxapi/route.go" > ./app/domain/$$name""api/route.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./app/domain/xxxapi/mid.go" > ./app/domain/$$name""api/mid.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./app/domain/xxxapi/xxxapi.go" > ./app/domain/$$name""api/$$name""api.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./business/domain/xxx/model.go" > ./business/domain/$$name/model.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./business/domain/xxx/xxx.go" > ./business/domain/$$name/$$name.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./business/domain/xxx/stores/xxxdb/model.go" > ./business/domain/$$name/stores/$$name""db/model.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./business/domain/xxx/stores/xxxdb/xxxdb.go" > ./business/domain/$$name/stores/$$name""db/$$name""db.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./app/domain/xxxapi/filter.go" > ./app/domain/$$name""api/filter.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./app/domain/xxxapi/order.go" > ./app/domain/$$name""api/order.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./business/domain/xxx/filter.go" > ./business/domain/$$name/filter.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./business/domain/xxx/order.go" > ./business/domain/$$name/order.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./business/domain/xxx/stores/xxxdb/filter.go" > ./business/domain/$$name/stores/$$name""db/filter.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		"./business/domain/xxx/stores/xxxdb/order.go" > ./business/domain/$$name/stores/$$name""db/order.go;
g-listing:
	@read -p "Please enter the domain name you want to add: " rawName;\
	read -p "Please enter the domain name for replacing abbr form: " rawAbbr;\
	read -p "Please enter the domain name for replacing plural form: " rawPlural;\
	name=$$(echo "$$rawName" | awk '{print tolower(substr($$0,1,1)) tolower(substr($$0,2))}');\
	u_name=$$(echo "$$rawName" | awk '{print toupper(substr($$0,1,1)) tolower(substr($$0,2))}');\
	abbr=$$(echo "$$rawAbbr" | awk '{print tolower(substr($$0,1,1)) tolower(substr($$0,2))}');\
	u_abbr=$$(echo "$$rawAbbr" | awk '{print toupper(substr($$0,1,1)) tolower(substr($$0,2))}');\
	plural=$$(echo "$$rawPlural" | awk '{print tolower(substr($$0,1,1)) tolower(substr($$0,2))}');\
	u_plural=$$(echo "$$rawPlural" | awk '{print toupper(substr($$0,1,1)) tolower(substr($$0,2))}');\
	if [ -d "./app/domain/$$name""api" ] || [ -d "./business/domain/$$name" ]; then \
		echo "Error: The domain '$$name' already exists. Aborting.";\
		exit 1;\
	else \
		mkdir ./app/domain/$$name""api ;\
		mkdir -p "./business/domain/$$name/stores/""$$name""db" ;\
		touch ./app/domain/$$name""api/model.go;\
		touch ./app/domain/$$name""api/route.go;\
		touch ./app/domain/$$name""api/$$name""api.go;\
		touch ./business/domain/$$name/model.go;\
		touch ./business/domain/$$name/$$name.go;\
		touch ./business/domain/$$name/stores/$$name""db/model.go;\
		touch ./business/domain/$$name/stores/$$name""db/$$name""db.go;\
	fi;\
	old_domain="xxx";\
	old_domain2="xx2x";\
	old_u_domain="XXX";\
	old_u_domain2="XX2X";\
	old_abbr="yyy";\
	old_u_abbr="YYY";\
	old_plural="zzz";\
	old_u_plural="ZZZ";\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		-e "s/$$old_domain2/$$name/g" \
		-e "s/$$old_u_domain2/$$u_name/g" \
		"./app/domain/xx2xapi/model.go" > ./app/domain/$$name""api/model.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		-e "s/$$old_domain2/$$name/g" \
		-e "s/$$old_u_domain2/$$u_name/g" \
		"./app/domain/xx2xapi/route.go" > ./app/domain/$$name""api/route.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		-e "s/$$old_domain2/$$name/g" \
		-e "s/$$old_u_domain2/$$u_name/g" \
		"./app/domain/xx2xapi/xx2xapi.go" > ./app/domain/$$name""api/$$name""api.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		-e "s/$$old_domain2/$$name/g" \
		-e "s/$$old_u_domain2/$$u_name/g" \
		"./business/domain/xx2x/model.go" > ./business/domain/$$name/model.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		-e "s/$$old_domain2/$$name/g" \
		-e "s/$$old_u_domain2/$$u_name/g" \
		"./business/domain/xx2x/xx2x.go" > ./business/domain/$$name/$$name.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		-e "s/$$old_domain2/$$name/g" \
		-e "s/$$old_u_domain2/$$u_name/g" \
		"./business/domain/xx2x/stores/xx2xdb/model.go" > ./business/domain/$$name/stores/$$name""db/model.go;\
	sed -e "s/$$old_domain/$$name/g" \
		-e "s/$$old_u_domain/$$u_name/g" \
		-e "s/$$old_abbr/$$abbr/g" \
		-e "s/$$old_u_abbr/$$u_abbr/g" \
		-e "s/$$old_plural/$$plural/g" \
		-e "s/$$old_u_plural/$$u_plural/g" \
		-e "s/$$old_domain2/$$name/g" \
		-e "s/$$old_u_domain2/$$u_name/g" \
		"./business/domain/xx2x/stores/xx2xdb/xx2xdb.go" > ./business/domain/$$name/stores/$$name""db/$$name""db.go;
