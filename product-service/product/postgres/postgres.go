package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lucasHSantiago/go-shop-ms/foundation/db"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
	"github.com/rs/zerolog/log"
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

	var dest productDb
	if err := db.NamedQueryStruct(ctx, s.db, query, &np, &dest); err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return toProduct(dest), nil
}

func (s *Store) GetAll(ctx context.Context, filter product.Filter, pageNumber int, rowsPerPage int) ([]*product.Product, error) {
	const query string = `
	SELECT id, name, description, price, category_id, created_at
	FROM products
	WHERE (name ILIKE COALESCE('%' || :name || '%', name))
	  AND (price = COALESCE(:price, price))
	  AND (category_id = COALESCE(:category_id, category_id))
	ORDER BY created_at DESC
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY
	`

	var dbPrds []productDb
	if err := db.NamedQuerySlice(ctx, s.db, query, toFilterDb(filter, pageNumber, rowsPerPage), &dbPrds); err != nil {
		log.Error().Err(err).Msg("failed to get products from the database")
		return nil, fmt.Errorf("failed to get products in the data base: %w", err)
	}

	return toDbProducts(dbPrds), nil
}
