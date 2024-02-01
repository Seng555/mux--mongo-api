// middleware/token.go
package middleware

import (
	"context"
	"errors"
	"mux-mongo-api/src/configs"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type UserInfo struct {
	Id string `json:"id"`
	// Add other user information as needed
}

// CheckToken is a middleware to check the JWT token
func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := extractToken(r)
		if err != nil {
			HandleBadRequest(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Validate and extract user information from the token
		userInfo, err := decodeJWT(token)
		if err != nil {
			HandleBadRequest(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Add user information to the request context
		ctx := context.WithValue(r.Context(), "auth", userInfo)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// extractToken extracts the JWT token from the request headers
func extractToken(r *http.Request) (string, error) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no token provided")
	}

	// Token usually has a format "Bearer <token>"
	// Split the header value to get the token part
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return "", errors.New("no token provided")
	}

	// Extract and return the actual token value
	token := authParts[1]
	//log.Print(token);
	return token, nil
}

// validateToken validates the JWT token and extracts user information
func decodeJWT(tokenString string) (*UserInfo, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		// Provide the secret key for validation
		return []byte(configs.EnvPrivateKey()), nil
	})

	// Check for errors
	if err != nil {
		return nil, err
	}

	// Verify and parse claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract user information from claims
		userInfo := &UserInfo{
			Id: claims["id"].(string),
			// Add other user information as needed
		}
		return userInfo, nil
	} else {
		return nil, errors.New("Invalid token")
	}
}
