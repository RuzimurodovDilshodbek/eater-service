PWD=$(shell pwd)
SERVICE=eater-svc
MIGRATION_PATH=${PWD}/src/infrastructure/migrations
PROTOS_PATH=$(PWD)/src/infrastructure/protos
pathI:
	cd $(PWD)/src/infrastructure/protos/eater
server:
	go run main.go

migration:
	migrate create -ext sql -dir pkg/database/migrations -seq $(table)

migrateup:
	migrate -database "postgres://postgres:postgres@localhost:5432/eater?sslmode=disable&search_path=public" -path ./pkg/database/migrations up

add-protos-submodules:
git submodule add git@github.com:RuzimurodovDilshodbek/service-protos.git ./src/infrastructure/protos

migratedown:
	migrate -database "postgres://postgres:postgres@localhost:5432/eater?sslmode=disable&search_path=public" -path ./pkg/database/migrations down

gen-eater-proto:
	protoc \
	--go_out=./src/application/protos \
	--go_opt=paths=import \
	--go-grpc_out=./src/application/protos \
	--go-grpc_opt=paths=import \
	-I=$(PROTOS_PATH)/eater \
	$(PROTOS_PATH)/eater/*.proto

docker:
	docker build --rm -t eater-svc -f ${PWD}/deploy/Dockerfile .

compose-up:
	docker-compose -f ./deploy/docker-compose.yml up

compose-down:
	docker-compose -f ./deploy/docker-compose.yml down