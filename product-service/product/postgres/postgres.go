package postgres

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucasHSantiago/go-shop-ms/foundation/dbsql"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
	"github.com/rs/zerolog/log"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Create(ctx context.Context, nn []product.NewProduct) ([]*product.Product, error) {
	builder := sq.Insert("products").
		Columns("name", "description", "price", "category_id").
		Suffix("RETURNING id, name, description, price, category_id, created_at").
		PlaceholderFormat(sq.Dollar)

	for _, p := range nn {
		builder = builder.Values(p.Name, p.Description, p.Price, p.CategoryId)
	}

	query, args := builder.MustSql()
	dest, err := dbsql.QuerySlice[productDb](ctx, s.db, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("failed to create products in the database")
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return toProducts(dest), nil
}

func (s *Store) Get(ctx context.Context, filter product.Filter, pageNumber int, rowsPerPage int) ([]*product.Product, error) {
	builder := sq.Select("id", "name", "description", "price", "category_id", "created_at").
		From("products").
		OrderBy("created_at DESC").
		Offset(uint64((pageNumber - 1) * rowsPerPage)).
		Limit(uint64(rowsPerPage)).
		PlaceholderFormat(sq.Dollar)

	if filter.Name != nil {
		builder = builder.Where(sq.ILike{"name": fmt.Sprintf("%%%s%%", *filter.Name)})
	}

	if filter.Price != nil {
		builder = builder.Where(sq.Eq{"price": *filter.Price})
	}

	if filter.CategoryId != nil {
		builder = builder.Where(sq.Eq{"category_id": *filter.CategoryId})
	}

	query, args := builder.MustSql()
	fmt.Printf("Query: %s, Args: %v\n", query, args)

	res, err := dbsql.QuerySlice[productDb](ctx, s.db, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("failed to create products in the database")
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return toProducts(res), nil
}
