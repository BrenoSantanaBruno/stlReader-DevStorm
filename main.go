package main

import (
	"fmt"
	"net/http"
	"stlReader-DevStorm/stlreader"
	"text/template"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/", http.FileServer(http.Dir("templates")))

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Maximum file size limit: 10MB
	file, _, err := r.FormFile("stlFile")
	if err != nil {
		http.Error(w, "Error receiving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	totalArea, numTriangles, err := stlreader.ProcessSTLFile(file)
	if err != nil {
		http.Error(w, "Error processing STL file", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error processing HTML template", http.StatusInternalServerError)
		return
	}

	data := struct {
		TotalArea    float64
		NumTriangles int
	}{
		TotalArea:    totalArea,
		NumTriangles: numTriangles,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error displaying HTML template", http.StatusInternalServerError)
		return
	}
}
