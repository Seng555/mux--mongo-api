// main.go
package main

import (
	"log"
	"mux-mongo-api/configs"
	"mux-mongo-api/router"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Run database
	if configs.DB == nil {
		log.Fatal("MongoDB connection not established. Exiting...")
	}

	// Routes path are defined
	r := mux.NewRouter()
	router.SetupUserRoutes(r)
	// Use the ErrorHandlerMiddleware globally
	//r.Use(middleware.ErrorHandlerMiddleware)

	// Start the server
	serverAddr := ":8080"
	log.Printf("Server is listening on %s...\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, r))
}
