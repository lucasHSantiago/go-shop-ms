package response

import (
	"encoding/json"
	"net/http"
)

// Respond writes the provided data as a JSON response to the http.ResponseWriter.
func Respond(w http.ResponseWriter, data any, statusCode int) error {
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

func BadRequest(w http.ResponseWriter, err error) error {
	return Respond(w, map[string]any{"error": err}, http.StatusBadRequest)
}
