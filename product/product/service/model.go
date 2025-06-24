package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-shop-ms/product/product/store"
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

func (np NewProduct) toDBProduct() store.Product {
	return store.Product{
		Name:        np.Name,
		Description: np.Description,
		Price:       np.Price,
		Category_id: np.Category_id,
	}
}
