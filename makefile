postgres_start:
	$(shell /bin/sh ./.scripts/postgres.sh);

postgres_stop:
	docker stop postgres10

createdb:
	docker exec -it postgres10 createdb --username=postgres --owner=postgres golang_masterclass

dropdb:
	docker exec -it postgres10 dropdb golang_masterclass

migrateup:
	migrate -path db/sql/postgresql/migration -database "postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/golang_masterclass?sslmode=disable" -verbose up

migratedown:
	migrate -path db/sql/postgresql/migration -database "postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/golang_masterclass?sslmode=disable" -verbose down

sqlc:
	sqlc generate

cache?=1

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
		go test -v -cover ./...; \
	else \
		go test -v -cover -count=1 ./...; \
	fi

install:
	@echo "Installing dependencies...."
	@rm -rf vendor
	@go mod download && go mod tidy && go mod vendor

.PHONY: postgres_start postgres_stop createdb dropdb migrateup migratedown sqlc test test_create_account