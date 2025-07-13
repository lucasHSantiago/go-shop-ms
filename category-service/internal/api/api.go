package api

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/lucasHSantiago/go-shop-ms/category/config"
	"github.com/lucasHSantiago/go-shop-ms/category/internal/grpc/pb"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func Serve(ctx context.Context, server pb.CategoryServiceServer, cfg config.Config) error {
	// -------------------------------------------------------------------------
	// Capture the interrupt signals so we can gracefully shutdown the server.

	ctx, stop := signal.NotifyContext(ctx, interruptSignals...)
	defer stop()

	waitGroup, ctx := errgroup.WithContext(ctx)

	// -------------------------------------------------------------------------
	// Start gRPC server

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", cfg.Web.APIHost)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	// -------------------------------------------------------------------------
	// Start server and listen for shutdown

	waitGroup.Go(func() error {
		log.Info().Msgf("start gRPC server at %s\n", listener.Addr().String())

		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}

			log.Error().Err(err).Msg("gRPC failed to serve")
			return err
		}

		return nil
	})

	// -------------------------------------------------------------------------
	// Gracefully shutdown the server when the context is done.

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown gRPC server")
		grpcServer.GracefulStop()
		log.Info().Msg("gRPC server was stopped")

		return nil
	})

	// -------------------------------------------------------------------------
	// Wait for the server to finish serving or shutdown.

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}

	return nil
}
