package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	repo := NewRepo()
	hdlr := NewHandler(repo)
	srv := NewServer(hdlr)

	err := srv.serve(ctx)
	if err != nil {
		return fmt.Errorf("cannot run server: %w", err)
	}

	return nil
}
