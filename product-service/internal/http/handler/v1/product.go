package v1

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-shop-ms/foundation/request"
	"github.com/lucasHSantiago/go-shop-ms/foundation/response"
	"github.com/lucasHSantiago/go-shop-ms/foundation/validate"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
	"github.com/rs/zerolog/log"
)

type NewProduct struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Category_id uuid.UUID `json:"category_id" validate:"required,uuid4"`
}

func (app NewProduct) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}

	return nil
}

func create(s product.UseCase) http.HandlerFunc {
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
