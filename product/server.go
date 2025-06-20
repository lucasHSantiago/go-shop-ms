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
)

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

	srv := &http.Server{
		Addr:    cfg.Web.APIHost,
		Handler: s.routes(),
		// ErrorLog:     log.New(app.logger, "", 0),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		// s := <-quit

		// app.logger.PrintInfo("shutting down server", map[string]string{
		// 	"signal": s.String(),
		// })

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		// app.logger.PrintInfo("completing background tasks", map[string]string{
		// 	"addr": srv.Addr,
		// })

		shutdownError <- nil
	}()

	// app.logger.PrintInfo("starting %s server on %s", map[string]string{
	// 	"addr": srv.Addr,
	// 	"env":  app.config.env,
	// })

	err = srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	// app.logger.PrintInfo("stopped server", map[string]string{
	// 	"addr": srv.Addr,
	// })

	return nil
}
