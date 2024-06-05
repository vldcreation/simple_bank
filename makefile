cache?=1
dev?=0
DB_URL?="postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_DATABASE)?sslmode=disable"

ifeq ($(dev), 1)
	DB_URL="postgresql://postgres:secret@localhost:5432/golang_masterclass?sslmode=disable"
endif

DB_start:
	$(shell /bin/sh ./.scripts/postgres.sh);

DB_stop:
	docker stop postgres10

createdb:
	docker exec -it postgres10 createdb --username=postgres --owner=postgres golang_masterclass

dropdb:
	docker exec -it postgres10 dropdb golang_masterclass

migrateup:
	migrate -path db/sql/postgresql/migration -database $(DB_URL) -verbose up

migrateup1:
	migrate -path db/sql/postgresql/migration -database $(DB_URL) -verbose up 1

migratedown:
	migrate -path db/sql/postgresql/migration -database $(DB_URL) -verbose down

migratedown1:
	migrate -path db/sql/postgresql/migration -database $(DB_URL) -verbose down 1

sqlc:
	sqlc generate

test_main_db:
	@echo "Running tests...."
	@echo "Cache is $(cache)"
	@if [ "$(cache)" = 1 ]; then \
		go test -timeout 30s -v -cover github.com/vldcreation/simple_bank/db/sql/postgresql/sqlc -run ^TestMain; \
	else \
		go test -timeout 30s -v -cover -count=1 github.com/vldcreation/simple_bank/db/sql/postgresql/sqlc -run ^TestMain; \
	fi

test_create_account:
	@echo "Running tests...."
	@echo "Cache is $(cache)"
	@if [ "$(cache)" = 1 ]; then \
		go test -timeout 30s -v -cover github.com/vldcreation/simple_bank/db/sql/postgresql/sqlc -run ^TestCreateAccount; \
	else \
		go test -timeout 30s -v -cover -count=1 github.com/vldcreation/simple_bank/db/sql/postgresql/sqlc -run ^TestCreateAccount; \
	fi
test:
	@echo "Running tests...."
	@echo "Cache is $(cache)"
	@if [ "$(cache)" = 1 ]; then \
		rm -f cover.out cover.html && \
		go test -v -cover -coverprofile cover.out ./... && \
		go tool cover -html cover.out -o cover.html && \
		echo "Dev mode is $(debug)" && \
		if [ "$(debug)" = 1 ]; then \
			open cover.html ; \
		fi; \
	else \
		rm -f cover.out cover.html && \
		go test -v -cover -coverprofile cover.out -count=1 ./... && \
		go tool cover -html cover.out -o cover.html && \
		echo "Dev mode is $(debug)" && \
		if [ "$(debug)" = 1 ]; then \
			open cover.html ; \
		fi; \
	fi

install:
	@echo "Installing dependencies...."
	@rm -rf vendor
	@go mod download && go mod tidy && go mod vendor

start:
	@echo "Starting server...."
	@go run main.go

.PHONY: DB_start DB_stop createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test test_create_account test_main_db install start