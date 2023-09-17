package handlers

import (
	"encoding/json"
	"net/http"
	"stlReader-DevStorm/stl"
)

// UploadFileHandler handles HTTP requests to upload and process an STL file.
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the uploaded file from the request
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Process the STL file to calculate total area and number of triangles
	areaTotal, numTriangles, err := stl.ProcessSTLFile(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a structure to store the results
	type Result struct {
		NumTriangles int     `json:"numTriangles"`
		AreaTotal    float64 `json:"areaTotal"`
	}

	result := Result{
		NumTriangles: numTriangles,
		AreaTotal:    areaTotal,
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
