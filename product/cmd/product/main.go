package main

import (
	"context"
	"fmt"
	"os"

	v1 "github.com/lucasHSantiago/go-shop-ms/product/internal/handler/v1"
	"github.com/lucasHSantiago/go-shop-ms/product/internal/server"
	"github.com/lucasHSantiago/go-shop-ms/product/internal/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	if build == "develop" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// -------------------------------------------------------------------------
	// Instantiate the repository, handler, and server.

	store := store.NewStore()
	v1 := v1.NewHandler(store)
	srv := server.NewServer(v1)

	// -------------------------------------------------------------------------
	// Start the server.

	err := srv.Serve(ctx)
	if err != nil {
		log.Error().Err(err).Msg("cannot run server")
		return err
	}

	return nil
}
