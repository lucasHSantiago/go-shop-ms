package store

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lucasHSantiago/go-shop-ms/foundation/db"
)

type Store struct {
	db sqlx.ExtContext
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Create(ctx context.Context, prd Product) (Product, error) {
	const query = `
	INSERT INTO products (name, description, price, category_id)
	VALUES (:name, :description, :price, :category_id)
	RETURNING id, name, description, price, category_id, created_at
	`

	var dest Product
	if err := db.NamedQueryStruct(ctx, s.db, query, &prd, &dest); err != nil {
		return Product{}, fmt.Errorf("failed to create product: %w", err)
	}

	return dest, nil
}
