// router/user.go
package router

import (
	usersController "mux-mongo-api/src/controller"
	"mux-mongo-api/src/middleware"

	"github.com/gorilla/mux"
)

// SetupUserRoutes sets up routes for "/api/user"
func SetupUserRoutes(r *mux.Router) {
	userRouter := r.PathPrefix("/api/auth").Subrouter()
	// Middleware to check JWT token
	userRouter.Use(middleware.CheckToken)
	//handles the /api/auth/register endpoint
	userRouter.HandleFunc("/register", usersController.CreateUser).Methods("POST")
	//handles the /api/auth/login endpoint
	userRouter.HandleFunc("/login", usersController.LoginUser).Methods("POST")
}
