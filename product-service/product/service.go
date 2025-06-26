package product

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Storer interface {
	Create(ctx context.Context, np []NewProduct) ([]*Product, error)
	GetAll(ctx context.Context, filter Filter, pageNumber int, rowsPerPage int) ([]*Product, error)
}

type Service struct {
	storer Storer
}

func NewService(s Storer) *Service {
	return &Service{
		storer: s,
	}
}

func (s *Service) Create(ctx context.Context, nn []NewProduct) ([]*Product, error) {
	// TODO: validate if category_id exists in the database

	pp, err := s.storer.Create(ctx, nn)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return pp, nil
}

func (s *Service) GetAll(ctx context.Context, filter Filter, pageNumber int, rowsPerPage int) ([]*Product, error) {
	pp, err := s.storer.GetAll(ctx, filter, pageNumber, rowsPerPage)
	if err != nil {
		log.Error().Err(err).Msg("failed to get products from the database")
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return pp, nil
}
