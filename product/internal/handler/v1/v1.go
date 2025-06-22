package v1

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/lucasHSantiago/go-shop-ms/foundation/request"
	"github.com/lucasHSantiago/go-shop-ms/foundation/response"
)

type Storer interface {
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
		return
	}

	err := response.Respond(w, np, http.StatusCreated)
	if err != nil {
		log.Error().Err(err).Msg("failed to respond")
		return
	}
}
