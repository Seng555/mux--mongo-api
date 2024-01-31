// router/user.go
package router

import (
	usersController "mux-mongo-api/controller"

	"github.com/gorilla/mux"
)

// SetupUserRoutes sets up routes for "/api/user"
func SetupUserRoutes(r *mux.Router) {
	userRouter := r.PathPrefix("/api/user").Subrouter()

	//handles the /api/user/ endpoint
	userRouter.HandleFunc("/", usersController.GetUser).Methods("GET")
}
