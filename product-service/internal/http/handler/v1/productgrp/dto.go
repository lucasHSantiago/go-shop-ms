package productgrp

import (
	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-shop-ms/foundation/validate"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
)

type NewProduct struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Category_id uuid.UUID `json:"category_id" validate:"required,uuid4"`
}

type NewProducts []NewProduct

func (np NewProduct) Validate() error {
	if err := validate.Check(np); err != nil {
		return err
	}

	return nil
}

func (np NewProduct) toNewProduct() product.NewProduct {
	return product.NewProduct{
		Name:        np.Name,
		Description: np.Description,
		Price:       np.Price,
		Category_id: np.Category_id,
	}
}

func (nps NewProducts) toNewProducts() []product.NewProduct {
	nn := make([]product.NewProduct, len(nps))
	for i, np := range nps {
		nn[i] = np.toNewProduct()
	}
	return nn
}

type Filter struct {
	Name        *string    `form:"name" validate:"omitempty"`
	Price       *float64   `form:"price" validate:"omitempty,gt=0"`
	Category_id *uuid.UUID `form:"category_id" validate:"omitempty,uuid4"`
}

func (p Filter) Validate() error {
	if err := validate.Check(p); err != nil {
		return err
	}

	return nil
}

func (p Filter) toProductFilter() product.Filter {
	return product.Filter{
		Name:       p.Name,
		Price:      p.Price,
		CategoryId: p.Category_id,
	}
}
