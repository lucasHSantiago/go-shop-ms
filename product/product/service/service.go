package service

import (
	"context"
	"fmt"

	"github.com/lucasHSantiago/go-shop-ms/product/product/store"
)

type Storer interface {
	Create(ctx context.Context, prd store.Product) (store.Product, error)
}

type Service struct {
	storer Storer
}

func NewService(s Storer) *Service {
	return &Service{
		storer: s,
	}
}

func (s *Service) Create(ctx context.Context, np NewProduct) (Product, error) {
	// TODO: validate if category_id exists in the database

	prd, err := s.storer.Create(ctx, np.toDBProduct())
	if err != nil {
		return Product{}, fmt.Errorf("failed to create product: %w", err)
	}

	return Product{
		ID:          prd.ID,
		Name:        prd.Name,
		Description: prd.Description,
		Price:       prd.Price,
		Category_id: prd.Category_id,
		Created_at:  prd.Created_at,
	}, nil
}
