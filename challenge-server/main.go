package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"stlReader-DevStorm/web/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.UploadFormHandler).Methods("GET")
	r.HandleFunc("/upload", handlers.UploadFileHandler).Methods("POST")

	// Configurar as políticas de CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},  // Substitua pelo domínio real do seu aplicativo React
		AllowedMethods: []string{"GET", "POST", "OPTIONS"}, // Permita os métodos HTTP necessários
		AllowedHeaders: []string{"*"},                      // Permita todos os cabeçalhos
	})

	// Use o middleware CORS com seu roteador
	handler := c.Handler(r)

	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
