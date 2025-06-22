package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type validator interface {
	Validate() error
}

// Decode reads the request body and decodes it into the provided value.
func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	if v, ok := val.(validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
