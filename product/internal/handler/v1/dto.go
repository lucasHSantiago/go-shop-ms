package v1

import (
	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-shop-ms/foundation/validate"
	"github.com/lucasHSantiago/go-shop-ms/product/internal/store"
)

type NewProduct struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Category_id uuid.UUID `json:"category_id" validate:"required,uuid4"`
}

func (app NewProduct) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}

	return nil
}

func toDBProduct(app NewProduct) store.Product {
	return store.Product{
		Name:        app.Name,
		Description: app.Description,
		Price:       app.Price,
		Category_id: app.Category_id,
	}
}
