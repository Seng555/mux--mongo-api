// controller/getuser.go
package usersController

import (
	"encoding/json"
	"mux-mongo-api/model"
	"net/http"
)

// GetUser handles the /api/user endpoint
// @Summary Get user details
// @Description Retrieve user details with a sample response
// @Produce json
// @Success 200 {object} model.Response "OK"
// @Router /user/ [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Sending an array of objects
	arrayOfObjects := []map[string]interface{}{
		{"id": 1, "name": "John"},
		{"id": 2, "name": "Jane"},
	}

	response := model.Response{
		Status:  http.StatusOK,
		Message: "Hello, Swagger!",
		Data:    arrayOfObjects,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
