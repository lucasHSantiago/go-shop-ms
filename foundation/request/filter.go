package request

import (
	"net/http"

	"github.com/go-playground/form"
	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-shop-ms/foundation/cerr"
)

// ParseFilter decodes the query parameters from the request into the provided value.
func ParseFilter(r *http.Request, val any) error {
	decoder := form.NewDecoder()

	// -------------------------------------------------------------------------
	// Custom type functions registration

	decoder.RegisterCustomTypeFunc(decodeUUID, uuid.UUID{})

	// -------------------------------------------------------------------------

	err := decoder.Decode(val, r.URL.Query())
	if err != nil {
		derrors, ok := err.(form.DecodeErrors)
		if !ok {
			return err
		}

		var fields cerr.FieldErrors
		for field, derror := range derrors {
			field := cerr.FieldError{
				Field: field,
				Err:   derror.Error(), // No translator available for form errors
			}
			fields = append(fields, field)
		}

		return fields
	}

	if v, ok := val.(validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Used to decod UUID values from query parameters.
func decodeUUID(val []string) (interface{}, error) {
	if len(val) == 0 {
		return nil, nil // No value provided, return nil
	}

	return uuid.Parse(val[0])
}
