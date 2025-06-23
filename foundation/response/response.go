package response

import (
	"encoding/json"
	"net/http"
)

// Response writes the provided data as a JSON response to the http.ResponseWriter.
func Response(w http.ResponseWriter, data any, statusCode int) error {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

// RequestError creates a JSON response with an error message and a specific status code.
func RequestError(w http.ResponseWriter, msg string, statusCode int) error {
	return Response(w, map[string]any{"error": msg}, statusCode)
}

// BadRequest creates a JSON response with an error message for a bad request.
func BadRequest(w http.ResponseWriter, err error) error {
	return Response(w, map[string]any{"error": err}, http.StatusBadRequest)
}

// InternalServerError creates a JSON response with an error message for a internal server error.
func InternalServerError(w http.ResponseWriter, err error) error {
	return Response(w, map[string]any{"error": err}, http.StatusInternalServerError)
}
