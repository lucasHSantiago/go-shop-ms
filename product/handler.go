package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/lucasHSantiago/go-shop-ms/foundation/request"
	"github.com/lucasHSantiago/go-shop-ms/foundation/response"
)

type storer interface {
}

type handler struct {
	storer storer
}

func NewHandler(storer storer) *handler {
	return &handler{
		storer: storer,
	}
}

type NewProduct struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category_id uuid.UUID `json:"category_id"`
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var np NewProduct
	if err := request.Decode(r, &np); err != nil {
		log.Error().Err(err).Msg("failed to decode request")
		return
	}

	err := response.Respond(w, np, http.StatusCreated)
	if err != nil {
		log.Error().Err(err).Msg("failed to respond")
		return
	}
}
