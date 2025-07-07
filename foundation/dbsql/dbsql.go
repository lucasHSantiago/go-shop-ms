// Package db provides support for access the database.
package dbsql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
)

const (
	uniqueViolation = "23505"
	undefinedTable  = "42P01"
)

// Set of error variables for CRUD operations.
var (
	ErrDBNotFound        = sql.ErrNoRows
	ErrDBDuplicatedEntry = errors.New("duplicated entry")
	ErrUndefinedTable    = errors.New("undefined table")
)

// Config is the required properties to use the database.
type Config struct {
	User       string
	Password   string
	Host       string
	Port       string
	Name       string
	Schema     string
	DisableTLS bool
}

func (c *Config) pgxConnString() string {
	sslMode := "require"
	if c.DisableTLS {
		sslMode = "disable"
	}
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		sslMode,
	)
}

// Open knows how to open a database connection based on the configuration.
func Open(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, cfg.pgxConnString())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// StatusCheck returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise.
func StatusCheck(ctx context.Context, db *pgxpool.Pool) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second)
		defer cancel()
	}

	var pingError error
	for attempts := 1; ; attempts++ {
		pingError = db.Ping(ctx)
		if pingError == nil {
			break
		}
		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Run a simple query to determine connectivity.
	// Running this query forces a round trip through the database.
	const q = `SELECT true`
	var tmp bool
	return db.QueryRow(ctx, q).Scan(&tmp)
}

func QuerySlice[T any](ctx context.Context, db *pgxpool.Pool, query string, args ...any) ([]*T, error) {
	q := queryString(query, args)
	log.Info().Str("query", q).Msg("database.QuerySlice")

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		if pqerr, ok := err.(*pgconn.PgError); ok && pqerr.Code == undefinedTable {
			switch pqerr.Code {
			case undefinedTable:
				return nil, ErrUndefinedTable
			case uniqueViolation:
				return nil, ErrDBDuplicatedEntry
			}
		}
		return nil, err
	}

	dest, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[T])
	if err != nil {
		return nil, err
	}

	return dest, nil
}

// queryString provides a pretty print version of the query and parameters using pgx style ($1, $2, ...).
func queryString(query string, args any) string {
	var params []any

	switch v := args.(type) {
	case []any:
		params = v
	default:
		params = []any{args}
	}

	for i, param := range params {
		var value string
		switch v := param.(type) {
		case string:
			value = fmt.Sprintf("'%s'", v)
		case []byte:
			value = fmt.Sprintf("'%s'", string(v))
		case uuid.UUID:
			value = fmt.Sprintf("'%s'", v.String())
		default:
			value = fmt.Sprintf("%#v", v)
		}
		placeholder := fmt.Sprintf("$%d", i+1)
		query = strings.Replace(query, placeholder, value, 1)
	}

	query = strings.ReplaceAll(query, "\t", "")
	query = strings.ReplaceAll(query, "\n", " ")

	return strings.Trim(query, " ")
}
