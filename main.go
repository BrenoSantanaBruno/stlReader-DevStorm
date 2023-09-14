package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"stlReader-DevStorm/web/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.UploadFormHandler).Methods("GET")
	r.HandleFunc("/upload", handlers.UploadFileHandler).Methods("POST")
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
