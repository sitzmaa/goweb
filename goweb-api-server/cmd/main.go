package main

import (
	"fmt"
	"goweb-server/internal/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register routes
	api.RegisterRoutes(r)

	// Add API Key middleware to the router
	r.Use(api.APIKeyMiddleware)

	// Start the server
	http.Handle("/", r)
	port := ":8080"
	fmt.Println("Server started at port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
