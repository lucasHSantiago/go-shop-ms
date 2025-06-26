# ==============================================================================
# Migrate

.PHONY: migrateup
migrateup:
	migrate -path ./product/postgres/migrate -database "postgresql://postgres:admin@localhost:5432/product?sslmode=disable" -verbose up $(or $(n))

.PHONY: migratedown
migratedown:
	migrate -path ./product/postgres/migrate -database "postgresql://postgres:admin@localhost:5432/product?sslmode=disable" -verbose down $(or $(n))

.PHONY: new_migration
new_migration:
	migrate create -ext sql -dir migrate -seq $(name)


# ==============================================================================
# Run

.PHONY: run
run:
	go run ./...