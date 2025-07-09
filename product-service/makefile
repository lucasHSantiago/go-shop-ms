# ==============================================================================
# Migrate

.PHONY: migrateup
migrateup:
	migrate -path ./product-service/postgres/migrate -database "postgresql://postgres:admin@localhost:5432/product?sslmode=disable" -verbose up $(or $(n))

.PHONY: migratedown
migratedown:
	migrate -path ./product-service/postgres/migrate -database "postgresql://postgres:admin@localhost:5432/product?sslmode=disable" -verbose down $(or $(n))

.PHONY: new_migration
new_migration:
	migrate create -ext sql -dir ./product-service/postgres/migrate -seq $(name)


# ==============================================================================
# Run

.PHONY: run
run:
	go run ./...

# ==============================================================================
# Protoc

.PHONY: protoc
protoc:
	rm -f internal/grpc/pb/*.go
	protoc --proto_path=internal/grpc/proto --go_out=internal/grpc/pb --go_opt=paths=source_relative \
	--go-grpc_out=internal/grpc/pb --go-grpc_opt=paths=source_relative \
	internal/grpc/proto/product.proto