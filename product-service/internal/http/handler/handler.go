package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasHSantiago/go-shop-ms/product/internal/http/handler/v1/productgrp"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
)

func NewHandler(s product.UseCase) http.Handler {
	r := chi.NewRouter()

	// -------------------------------------------------------------------------
	// v1 routes

	r.Route("/v1", func(r chi.Router) {
		r.Post("/product", productgrp.Create(s))
		r.Get("/product", productgrp.GetAll(s))
	})

	// -------------------------------------------------------------------------

	return r
}
