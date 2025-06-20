package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

type app interface {
}

type server struct {
	app app
}

func NewServer(app app) *server {
	return &server{
		app: app,
	}
}

func (s *server) serve(ctx context.Context) error {
	// -------------------------------------------------------------------------
	// Load configuration.
	cfg := struct {
		Web struct {
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:10s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
			ShutdownTimeout time.Duration `conf:"default:20s"`
			APIHost         string        `conf:"default:0.0.0.0:3000"`
		}
	}{}

	const prefix = "PRODUCT"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// -------------------------------------------------------------------------
	// Run server

	srv := &http.Server{
		Addr:    cfg.Web.APIHost,
		Handler: s.routes(),
		// ErrorLog:     log.New(app.logger, "", 0),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// -------------------------------------------------------------------------
	// Capture the interrupt signals so we can gracefully shutdown the server.

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	waitGroup, ctx := errgroup.WithContext(ctx)

	// -------------------------------------------------------------------------
	// Start server and listen for shutdown

	waitGroup.Go(func() error {
		// log.Info().Msgf("start gateway server at %s", cfg.Web.APIHost)
		err = srv.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}

			// log.Fatal().Err(err).Msg("gateway server failed to serve")
			return err
		}

		return nil
	})

	// -------------------------------------------------------------------------
	// Gracefully shutdown the server when the context is done.

	waitGroup.Go(func() error {
		<-ctx.Done()
		// log.Info().Msg("graceful shutdown gateway server")
		srv.Shutdown(context.Background())
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			// log.Error().Err(err).Msg("failed to shutdown gateway server")
			return err
		}

		// log.Info().Msg("HTTP gateway server was stopped")
		return nil
	})

	// -------------------------------------------------------------------------
	// Wait for the server to finish serving or shutdown.

	err = waitGroup.Wait()
	if err != nil {
		// log.Fatal().Err(err).Msg("error from wait group")
	}

	return nil
}

func (s *server) routes() http.Handler {
	r := chi.NewRouter()

	return r
}
