package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lucasHSantiago/go-shop-ms/foundation/dbsql"
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

type pageDb struct {
	Offset      int `db:"offset"`
	RowsPerPage int `db:"rows_per_page"`
}

type filterDb struct {
	pageDb
	Name       pgtype.Text   `db:"name"`
	Price      pgtype.Float8 `db:"price"`
	CategoryId pgtype.UUID   `db:"category_id"`
}

func toFilterDb(f product.Filter, pageNumber int, rowsPerPage int) filterDb {
	return filterDb{
		pageDb: pageDb{
			Offset:      (pageNumber - 1) * rowsPerPage,
			RowsPerPage: rowsPerPage,
		},
		Name:       dbsql.StringToText(f.Name),
		Price:      dbsql.Float64ToFloat8(f.Price),
		CategoryId: dbsql.UUIDToUUID(f.CategoryId),
	}
}
