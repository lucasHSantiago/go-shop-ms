package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) routes() http.Handler {
	r := chi.NewRouter()

	// -------------------------------------------------------------------------
	// v1

	r.Route("/v1/product", func(r chi.Router) {
		r.Post("/", s.prdHdlr.Create)
	})

	// -------------------------------------------------------------------------

	return r
}
