package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
)

func NewHandler(s product.UseCase) http.Handler {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Post("/product", create(s))
	})

	return r
}
