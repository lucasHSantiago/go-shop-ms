package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lucasHSantiago/go-shop-ms/category/category"
	"github.com/lucasHSantiago/go-shop-ms/category/category/postgres"
	"github.com/lucasHSantiago/go-shop-ms/category/config"
	"github.com/lucasHSantiago/go-shop-ms/category/internal/api"
	"github.com/lucasHSantiago/go-shop-ms/category/internal/grpc/server"
	"github.com/lucasHSantiago/go-shop-ms/foundation/dbsql"
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

	dbConn, err := dbsql.Open(ctx, dbsql.Config{
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Host:       cfg.DB.Host,
		Port:       cfg.DB.Port,
		Name:       cfg.DB.Name,
		Schema:     "public",
		DisableTLS: cfg.DB.DisableTLS,
	})
	if err != nil {
		log.Error().Err(err).Msg("connecting to db")
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Info().Str("host", cfg.DB.Host).Msg("stopping database support")
		dbConn.Close()
	}()

	if err := dbsql.StatusCheck(ctx, dbConn); err != nil {
		log.Error().Err(err).Msg("cannot connect to database")
		return fmt.Errorf("cannot connect to database: %w", err)
	}

	// -------------------------------------------------------------------------
	// Instantiate the repository, handler, service, and server.

	store := postgres.NewStore(dbConn)
	service := category.NewService(store)
	server := server.NewCategoryServer(service)

	// -------------------------------------------------------------------------
	// Start the server.

	err = api.Serve(ctx, server, cfg)
	if err != nil {
		log.Error().Err(err).Msg("cannot run server")
		return fmt.Errorf("cannot run server: %w", err)
	}

	return nil
}
