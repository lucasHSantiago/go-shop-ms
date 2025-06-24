package v1

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/lucasHSantiago/go-shop-ms/foundation/request"
	"github.com/lucasHSantiago/go-shop-ms/foundation/response"
	"github.com/lucasHSantiago/go-shop-ms/product/product/service"
)

type Service interface {
	Create(ctx context.Context, prd service.NewProduct) (service.Product, error)
}

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto NewProduct
	if err := request.Decode(r, &dto); err != nil {
		log.Error().Err(err).Msg("failed to decode request")
		response.BadRequest(w, err)
		return
	}

	np := service.NewProduct{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Category_id: dto.Category_id,
	}

	prd, err := h.service.Create(r.Context(), np)
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
