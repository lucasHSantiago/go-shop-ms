package productgrp

import (
	"net/http"

	"github.com/lucasHSantiago/go-shop-ms/foundation/request"
	"github.com/lucasHSantiago/go-shop-ms/foundation/response"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
	"github.com/rs/zerolog/log"
)

func Create(s product.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto NewProduct
		if err := request.Decode(r, &dto); err != nil {
			log.Error().Err(err).Msg("failed to decode request")
			response.BadRequest(w, err)
			return
		}

		np := product.NewProduct{
			Name:        dto.Name,
			Description: dto.Description,
			Price:       dto.Price,
			Category_id: dto.Category_id,
		}

		prd, err := s.Create(r.Context(), np)
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
}

func GetAll(s product.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filter Filter
		err := request.ParseFilter(r, &filter)
		if err != nil {
			log.Error().Err(err).Msg("failed to parse filter")
			response.BadRequest(w, err)
			return
		}

		page, err := request.ParsePage(r)
		if err != nil {
			log.Error().Err(err).Msg("failed to get page number")
			response.BadRequest(w, err)
			return
		}

		prds, err := s.GetAll(r.Context(), filter.toProductFilter(), page.Number, page.RowsPerPage)
		if err != nil {
			log.Error().Err(err).Msg("failed to get products")
			response.InternalServerError(w, err)
			return
		}

		status := http.StatusOK
		if len(prds) == 0 {
			status = http.StatusNoContent
		}

		err = response.Response(w, prds, status)
		if err != nil {
			log.Error().Err(err).Msg("failed to respond")
			return
		}
	}
}
