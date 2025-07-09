# ==============================================================================
# Migrate

.PHONY: migrateup
migrateup:
	migrate -path ./category-service/postgres/migrate -database "postgresql://postgres:admin@localhost:5432/product?sslmode=disable" -verbose up $(or $(n))

.PHONY: migratedown
migratedown:
	migrate -path ./category-service/postgres/migrate -database "postgresql://postgres:admin@localhost:5432/product?sslmode=disable" -verbose down $(or $(n))

.PHONY: new_migration
new_migration:
	migrate create -ext sql -dir ./category-service/postgres/migrate -seq $(name)