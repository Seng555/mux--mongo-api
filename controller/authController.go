// controller/getuser.go
package usersController

import (
	"context"
	"encoding/json"
	"log"
	"mux-mongo-api/configs"
	"mux-mongo-api/helper"
	"mux-mongo-api/model"
	"net/http"

	"mux-mongo-api/middleware"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection("users")

// CreateUser handles the /auth/register endpoint
func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	// Sample user data
	var newUser model.User

	// Parse the request body into newUser

	err := json.NewDecoder(r.Body).Decode(&newUser)
	log.Print(newUser)
	if err != nil {
		middleware.HandleBadRequest(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Create a new user with default values
	newUser = *model.NewUser(newUser.Email, newUser.Password)

	// Validate email format
	if !helper.IsValidEmail(newUser.Email) {
		middleware.HandleBadRequest(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	existingUser, err := getUserByEmail(ctx, newUser.Email)
	if err != nil {
		// Handle the error as needed
		middleware.HandleBadRequest(w, http.StatusBadRequest, err.Error())
		return
	}
	if existingUser != nil {
		// User with the given email already exists, handle accordingly
		middleware.HandleBadRequest(w, http.StatusBadRequest, "User with this email already exists")
		return
	}
	// For the sake of this example, directly insert newUser into the database
	newSave, insertErr := userCollection.InsertOne(ctx, newUser)
	if insertErr != nil {
		middleware.HandleBadRequest(w, http.StatusInternalServerError, insertErr.Error())
		return
	}

	// Sending a single object (user)
	response := model.NewSuccessResponse(newSave)

	// Respond with the JSON-encoded response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		middleware.HandleBadRequest(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func getUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var existingUser model.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		return nil, nil // No user found, return nil without error
	} else if err != nil {
		// Handle other errors, log or return an appropriate error message
		return nil, err
	}
	return &existingUser, nil
}
