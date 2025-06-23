package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lucasHSantiago/go-shop-ms/foundation/db"
	"github.com/lucasHSantiago/go-shop-ms/product/config"
	v1 "github.com/lucasHSantiago/go-shop-ms/product/product/handler/v1"
	"github.com/lucasHSantiago/go-shop-ms/product/product/store"
	"github.com/lucasHSantiago/go-shop-ms/product/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var build = "develop"

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// -------------------------------------------------------------------------
	// Config logger

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if build == "develop" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// -------------------------------------------------------------------------
	// Load configuration.

	cfg, err := config.Parse()
	if err != nil {
		log.Error().Err(err).Msg("cannot parse configuration")
		return fmt.Errorf("cannot parse configuration: %w", err)
	}

	// -------------------------------------------------------------------------
	// Load Database

	log.Info().Str("host", cfg.DB.Host).Msg("initializing database support")

	dbConn, err := db.Open(db.Config{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Info().Str("host", cfg.DB.Host).Msg("stopping database support")
		dbConn.Close()
	}()

	if err := db.StatusCheck(ctx, dbConn); err != nil {
		log.Error().Err(err).Msg("cannot connect to database")
		return fmt.Errorf("cannot connect to database: %w", err)
	}

	// -------------------------------------------------------------------------
	// Instantiate the repository, handler, and server.

	store := store.NewStore(dbConn)
	v1 := v1.NewHandler(store)
	srv := server.NewServer(v1)

	// -------------------------------------------------------------------------
	// Start the server.

	err = srv.Serve(ctx, cfg)
	if err != nil {
		log.Error().Err(err).Msg("cannot run server")
		return fmt.Errorf("cannot run server: %w", err)
	}

	return nil
}
