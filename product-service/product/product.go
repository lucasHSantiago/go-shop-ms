package product

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Product represents a product in the system.
type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	CategoryId  uuid.UUID
	Created_at  time.Time
}

// NewProduct represents the data required to create a new product.
type NewProduct struct {
	Name        string
	Description string
	Price       float64
	Category_id uuid.UUID
}

// Filter represents the criteria for filtering products.
type Filter struct {
	Name       *string
	Price      *float64
	CategoryId *uuid.UUID
}

var ErrNotFound = errors.New("product not found")

// UseCase defines the interface for product use cases.
type UseCase interface {
	Create(ctx context.Context, nn []NewProduct) ([]*Product, error)
	Get(ctx context.Context, filter Filter, pageNumber int, rowsPerPage int) ([]*Product, error)
}
