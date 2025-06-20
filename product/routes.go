package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *server) routes() http.Handler {
	r := chi.NewRouter()

	return r
}
