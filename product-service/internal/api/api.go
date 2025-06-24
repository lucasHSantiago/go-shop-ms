package api

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lucasHSantiago/go-shop-ms/foundation/logger"
	"github.com/lucasHSantiago/go-shop-ms/product/config"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func Serve(ctx context.Context, hdl http.Handler, cfg config.Config) error {
	// -------------------------------------------------------------------------
	// Run server

	srv := &http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      hdl,
		ErrorLog:     logger.NewStdLogger(),
		IdleTimeout:  cfg.Web.IdleTimeout,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
	}

	// -------------------------------------------------------------------------
	// Capture the interrupt signals so we can gracefully shutdown the server.

	ctx, stop := signal.NotifyContext(ctx, interruptSignals...)
	defer stop()

	waitGroup, ctx := errgroup.WithContext(ctx)

	// -------------------------------------------------------------------------
	// Start server and listen for shutdown

	waitGroup.Go(func() error {
		log.Info().Msgf("start product service at %s", cfg.Web.APIHost)
		err := srv.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}

			log.Fatal().Err(err).Msg("product service failed to serve")
			return err
		}

		return nil
	})

	// -------------------------------------------------------------------------
	// Gracefully shutdown the server when the context is done.

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown product service")

		ctx, cancel := context.WithTimeout(ctx, cfg.Web.ShutdownTimeout)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("failed to shutdown product service")
			return err
		}

		log.Info().Msg("HTTP product service was stopped")
		return nil
	})

	// -------------------------------------------------------------------------
	// Wait for the server to finish serving or shutdown.

	err := waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}

	return nil
}
