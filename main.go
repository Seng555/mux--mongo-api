// main.go
package main

import (
	"mux-mongo-api/docs"
	"mux-mongo-api/router"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Your API Title"
	docs.SwaggerInfo.Description = "Your API description"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := mux.NewRouter()

	// Routes path are defined
	router.SetupUserRoutes(r)

	// Swagger endpoint
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	http.Handle("/", r)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
