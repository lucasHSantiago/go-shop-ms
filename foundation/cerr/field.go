package cerr

import (
	"encoding/json"
)

// FieldError is used to indicate an error with a specific request field.
type FieldError struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}

// FieldErrors represents a collection of field errors.
type FieldErrors []FieldError

// NewFieldsError creates an fields error.
func NewFieldsError(field string, err error) error {
	return FieldErrors{
		{
			Field: field,
			Err:   err.Error(),
		},
	}
}

// Error implements the error interface.
func (fe FieldErrors) Error() string {
	d, err := json.Marshal(fe)
	if err != nil {
		return err.Error()
	}
	return string(d)
}

// FieldErrorIndex is used to indicate an error with a specific request field at a specific index.
type FieldIndexError struct {
	Field string `json:"field"`
	Index int    `json:"index"`
	Err   string `json:"error"`
}

// FieldErrors represents a collection of field errors.
type FieldIndexErrors []FieldIndexError

// NewFieldsError creates an fields error.
func NewFieldIndexError(field string, idx int, err error) error {
	return FieldIndexErrors{
		{
			Field: field,
			Index: idx,
			Err:   err.Error(),
		},
	}
}

// Error implements the error interface.
func (fe FieldIndexErrors) Error() string {
	d, err := json.Marshal(fe)
	if err != nil {
		return err.Error()
	}
	return string(d)
}
