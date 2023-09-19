package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"mime/multipart"
	"net/http"
	"stlReader-DevStorm/stl"
	"stlReader-DevStorm/web/handlers"
)

func main() {
	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Define a route for handling file uploads via POST requests
	r.HandleFunc("/process-ascii-stl", handlers.UploadFileHandler).Methods("POST")

	// Define a route for handling binary STL file processing
	r.HandleFunc("/process-binary-stl", ProcessBinarySTLHandler).Methods("POST")

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
	fmt.Println("Server listening on :8080...")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

// ProcessBinarySTLHandler handles binary STL file processing
func ProcessBinarySTLHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the uploaded file from the request
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(file)

	// Process the binary STL file
	_, numTriangles, err := stl.ProcessBinarySTL(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a structure to store the results
	type Result struct {
		NumTriangles int `json:"numTriangles"`
	}

	result := Result{
		NumTriangles: numTriangles,
	}

	// Convert the result structure to JSON
	jsonResult, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON response
	w.Write(jsonResult)
}
