// middleware/error_handler.go
package middleware

import (
	"encoding/json"
	"net/http"

	"mux-mongo-api/model"
)

// ErrorHandlerMiddleware is a middleware function to handle errors globally
func ErrorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom ResponseWriter to capture the status code
		responseWriter := NewStatusRecorder(w)

		// Pass the custom ResponseWriter to the next handler
		next.ServeHTTP(responseWriter, r)

		// Check the captured status code
		status := responseWriter.Status()
		if status >= http.StatusInternalServerError {
			// Handle 500 Internal Server Error
			handleServerError(w, status)
		} else if status >= http.StatusBadRequest {
			// Handle 400 Bad Request
			HandleBadRequest(w, status, "Bad Request")
		}
	})
}

// NewStatusRecorder creates a new StatusRecorder
func NewStatusRecorder(w http.ResponseWriter) *StatusRecorder {
	return &StatusRecorder{
		ResponseWriter: w,
		status:         http.StatusOK,
	}
}

// StatusRecorder is a custom ResponseWriter to capture the status code
type StatusRecorder struct {
	http.ResponseWriter
	status int
}

// WriteHeader captures the status code
func (sr *StatusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}

// Status returns the captured status code
func (sr *StatusRecorder) Status() int {
	return sr.status
}

func handleServerError(w http.ResponseWriter, status int) {
	// Handle 500 Internal Server Error
	response := model.NewErrorResponse(status, "Internal Server Error")
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
