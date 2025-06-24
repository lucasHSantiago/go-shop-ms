package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	v1 "github.com/lucasHSantiago/go-shop-ms/product/internal/http/handler/v1"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
)

func NewHandler(s product.UseCase) http.Handler {
	r := chi.NewRouter()

	r.Mount("/", v1.NewHandler(s))

	return r
}
