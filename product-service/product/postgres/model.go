package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Category_id uuid.UUID `db:"category_id"`
	Created_at  time.Time `db:"created_at"`
}
