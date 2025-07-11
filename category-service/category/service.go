package category

import "context"

// -------------------------------------------------------------------------
// Storer

type Storer interface {
	Create(ctx context.Context, np []NewCategory) ([]*Category, error)
	Get(ctx context.Context, filter Filter, pageNumber int, rowsPerPage int) ([]*Category, error)
}

// -------------------------------------------------------------------------
// Service

type Service struct {
	storer Storer
}

func (s Service) Create(ctx context.Context, nn []NewCategory) ([]*Category, error) {
	panic("not implemented") // TODO: Implement
}

func (s Service) Get(ctx context.Context, filter Filter, pageNumber int, rowsPerPage int) ([]*Category, error) {
	panic("not implemented") // TODO: Implement
}
