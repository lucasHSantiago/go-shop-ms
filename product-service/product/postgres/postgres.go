package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lucasHSantiago/go-shop-ms/foundation/db"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
)

type Store struct {
	db sqlx.ExtContext
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Create(ctx context.Context, np product.NewProduct) (*product.Product, error) {
	const query = `
	INSERT INTO products (name, description, price, category_id)
	VALUES (:name, :description, :price, :category_id)
	RETURNING id, name, description, price, category_id, created_at
	`

	var dest Product
	if err := db.NamedQueryStruct(ctx, s.db, query, &np, &dest); err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return toProduct(dest), nil
}

func toProduct(prd Product) *product.Product {
	return &product.Product{
		ID:          prd.ID,
		Name:        prd.Name,
		Description: prd.Description,
		Price:       prd.Price,
		Category_id: prd.Category_id,
		Created_at:  prd.Created_at,
	}
}
