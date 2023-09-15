package handlers

import (
	"encoding/json"
	"net/http"
	"stlReader-DevStorm/stl"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

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

	// Convert the structure to JSON
	jsonResult, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write response JSON
	w.Write(jsonResult)
}
