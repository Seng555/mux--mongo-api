// controller/getuser.go
package usersController

import (
	"context"
	"encoding/json"
	"log"
	"mux-mongo-api/src/configs"
	"mux-mongo-api/src/helper"
	"mux-mongo-api/src/middleware"
	"mux-mongo-api/src/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = configs.GetCollection("users")

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	// Parse the request body into newUser
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		middleware.HandleBadRequest(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Check if Email and Password are provided
	if newUser.Email == "" || newUser.Password == "" {
		middleware.HandleBadRequest(w, http.StatusBadRequest, "Email and Password are required")
		return
	}

	// Hash the user's password
	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		log.Print(err.Error())
		middleware.HandleandleServerError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	// Create a new user with default values
	newUser = *model.NewUser(newUser.Email, hashedPassword)

	// Validate email format
	if !helper.IsValidEmail(newUser.Email) {
		middleware.HandleBadRequest(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	// Check if the user already exists
	existingUser, err := getUserByEmail(ctx, newUser.Email)
	if err != nil {
		middleware.HandleandleServerError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if existingUser != nil {
		middleware.HandleBadRequest(w, http.StatusBadRequest, "User with this email already exists")
		return
	}

	// For the sake of this example, directly insert newUser into the database
	newSave, insertErr := userCollection.InsertOne(ctx, newUser)
	if insertErr != nil {
		middleware.HandleandleServerError(w, http.StatusInternalServerError, insertErr.Error())
		return
	}

	// Sending a single object (user)
	response := model.NewSuccessResponse(newSave)

	// Respond with the JSON-encoded response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		middleware.HandleandleServerError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	auth, _ := ctx.Value("auth").(*middleware.UserInfo)
	log.Print(auth.Id)
	// Parse the request body into credentials
	var credentials model.Login
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		middleware.HandleBadRequest(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Check if Email and Password are provided
	if credentials.Email == "" || credentials.Password == "" {
		middleware.HandleBadRequest(w, http.StatusBadRequest, "Email and Password are required")
		return
	}

	// Retrieve the user from the database based on the provided email
	existingUser, err := getUserByEmail(ctx, credentials.Email)
	if err != nil {
		middleware.HandleandleServerError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if existingUser == nil {
		middleware.HandleBadRequest(w, http.StatusUnauthorized, "Invalid email")
		return
	}

	// Check if the provided password matches the hashed password in the database
	if !CheckPasswordHash(credentials.Password, existingUser.Password) {
		middleware.HandleBadRequest(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	// Generate JWT token
	token, err := generateJWT(existingUser.Id.Hex()) // assuming _id is an ObjectId
	if err != nil {
		middleware.HandleandleServerError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the JWT token
	response := model.NewSuccessResponse(map[string]string{"token": token})
	if err := json.NewEncoder(w).Encode(response); err != nil {
		middleware.HandleandleServerError(w, http.StatusInternalServerError, err.Error())
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// generateJWT generates a JWT token for the given email
func generateJWT(user_id string) (string, error) {
	// Set up the claims
	claims := jwt.MapClaims{
		"id":  user_id,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(configs.EnvPrivateKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
