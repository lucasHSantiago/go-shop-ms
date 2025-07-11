package category

import (
	"context"
	"time"
)

type Category struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

type NewCategory struct {
	Name string `validate:"required"`
}

type Filter struct {
	Name *string
}

type UseCase interface {
	Create(ctx context.Context, nn []NewCategory) ([]*Category, error)
	Get(ctx context.Context, filter Filter, pageNumber int, rowsPerPage int) ([]*Category, error)
}
