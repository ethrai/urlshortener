.PHONY: all clean test run

DB_DRIVER=sqlite3
MIGRATIONS_DIR=migrations

create-migration:
	migrate create -ext=sql -dir=$(MIGRATIONS_DIR) -seq init

migrate-up:
	migrate -path=$(MIGRATIONS_DIR) -database=$(DB_DRIVER)://${DB_URL} -verbose up

migrate-down:
	migrate -path=$(MIGRATIONS_DIR) -database=$(DB_DRIVER)://${DB_URL} -verbose down

