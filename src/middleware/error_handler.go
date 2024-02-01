// middleware/error_handler.go
package middleware

import (
	"encoding/json"
	"net/http"

	"mux-mongo-api/src/model"
)

func HandleandleServerError(w http.ResponseWriter, status int, errMessage string) {
	// Handle 500 Internal Server Error
	response := model.NewErrorResponse(status, errMessage)
	writeResponse(w, response, status)
}

func HandleBadRequest(w http.ResponseWriter, status int, errMessage string) {
	// Handle 400 Bad Request with custom error message
	response := model.NewErrorResponse(status, errMessage)
	writeResponse(w, response, status)
}

func writeResponse(w http.ResponseWriter, response *model.Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Encode the response to JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
