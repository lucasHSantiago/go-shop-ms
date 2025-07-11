package postgres

import (
	"context"

	"github.com/lucasHSantiago/go-shop-ms/category/category"
)

type Store struct {
}

func (s Store) Create(ctx context.Context, np []category.NewCategory) ([]*category.Category, error) {
	panic("not implemented") // TODO: Implement
}

func (s Store) Get(ctx context.Context, filter category.Filter, pageNumber int, rowsPerPage int) ([]*category.Category, error) {
	panic("not implemented") // TODO: Implement
}
