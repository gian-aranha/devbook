package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a JSON response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if erro := json.NewEncoder(w).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Error returns an error in JSON
func Error(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: erro.Error(),
	})
}