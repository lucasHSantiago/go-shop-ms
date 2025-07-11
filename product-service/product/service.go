package product

import (
	"context"
	"fmt"

	"github.com/lucasHSantiago/go-shop-ms/foundation/validate"
	"github.com/rs/zerolog/log"
)

// -------------------------------------------------------------------------
// Storer

type Storer interface {
	Create(ctx context.Context, np []NewProduct) ([]*Product, error)
	Get(ctx context.Context, filter Filter, pageNumber int, rowsPerPage int) ([]*Product, error)
}

type Service struct {
	storer Storer
}

// -------------------------------------------------------------------------
// Service

func NewService(s Storer) *Service {
	return &Service{
		storer: s,
	}
}

func (s Service) Create(ctx context.Context, nn []NewProduct) ([]*Product, error) {
	// Validate each new product
	for i, np := range nn {
		if err := validate.CheckWithIndex(np, i); err != nil {
			return nil, err
		}
	}

	// TODO: validate if category_id exists in the database

	pp, err := s.storer.Create(ctx, nn)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return pp, nil
}

func (s Service) Get(ctx context.Context, filter Filter, pageNumber int, rowsPerPage int) ([]*Product, error) {
	if pageNumber < 1 {
		pageNumber = 1 // Default to the first page
	}

	if rowsPerPage < 1 {
		rowsPerPage = 10 // Default value
	}

	pp, err := s.storer.Get(ctx, filter, pageNumber, rowsPerPage)
	if err != nil {
		log.Error().Err(err).Msg("failed to get products from the database")
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return pp, nil
}
