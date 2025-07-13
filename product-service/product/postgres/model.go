package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
)

type productDb struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Category_id uuid.UUID `db:"category_id"`
	Created_at  time.Time `db:"created_at"`
}

func (prd productDb) toProduct() *product.Product {
	return &product.Product{
		ID:          prd.ID,
		Name:        prd.Name,
		Description: prd.Description,
		Price:       prd.Price,
		CategoryId:  prd.Category_id,
		CreatedAt:   prd.Created_at,
	}
}

func toProducts(pp []*productDb) []*product.Product {
	products := make([]*product.Product, len(pp))
	for i, prd := range pp {
		products[i] = prd.toProduct()
	}
	return products
}
