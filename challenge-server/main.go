package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"stlReader-DevStorm/web/handlers"
)

func main() {
	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Define a route for handling file uploads via POST requests
	r.HandleFunc("/upload", handlers.UploadFileHandler).Methods("POST")

	// Configure CORS policies
	c := cors.New(cors.Options{
		// Specify the allowed origins (replace with your React app's real domain)
		AllowedOrigins: []string{"http://localhost:3000"},

		// Allow necessary HTTP methods
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},

		// Allow all headers
		AllowedHeaders: []string{"*"},
	})

	// Use the CORS middleware with your router
	handler := c.Handler(r)

	// Set up an HTTP server to listen on port 8080
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
