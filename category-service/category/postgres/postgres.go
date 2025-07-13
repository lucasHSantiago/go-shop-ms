package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucasHSantiago/go-shop-ms/category/category"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

func (s Store) Create(ctx context.Context, np []category.NewCategory) ([]*category.Category, error) {
	panic("not implemented") // TODO: Implement
}

func (s Store) Get(ctx context.Context, filter category.Filter, pageNumber int, rowsPerPage int) ([]*category.Category, error) {
	panic("not implemented") // TODO: Implement
}
