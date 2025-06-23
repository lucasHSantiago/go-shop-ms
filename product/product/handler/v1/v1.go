package v1

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/lucasHSantiago/go-shop-ms/foundation/request"
	"github.com/lucasHSantiago/go-shop-ms/foundation/response"
	"github.com/lucasHSantiago/go-shop-ms/product/product/store"
)

type Storer interface {
	Create(ctx context.Context, prd store.Product) (store.Product, error)
}

type Handler struct {
	storer Storer
}

func NewHandler(storer Storer) *Handler {
	return &Handler{
		storer: storer,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var np NewProduct
	if err := request.Decode(r, &np); err != nil {
		log.Error().Err(err).Msg("failed to decode request")
		response.BadRequest(w, err)
		return
	}

	// TODO: validate if category_id exists in the database

	prd, err := h.storer.Create(r.Context(), toDBProduct(np))
	if err != nil {
		log.Error().Err(err).Msg("failed to create product")
		response.InternalServerError(w, err)
		return
	}

	err = response.Response(w, prd, http.StatusCreated)
	if err != nil {
		log.Error().Err(err).Msg("failed to respond")
		return
	}
}
