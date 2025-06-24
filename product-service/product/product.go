package product

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	Category_id uuid.UUID
	Created_at  time.Time
}

type NewProduct struct {
	Name        string
	Description string
	Price       float64
	Category_id uuid.UUID
}

type UseCase interface {
	Create(ctx context.Context, prd NewProduct) (Product, error)
}
