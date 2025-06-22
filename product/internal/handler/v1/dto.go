package v1

import "github.com/google/uuid"

type NewProduct struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category_id uuid.UUID `json:"category_id"`
}
