package handlers

import (
	"html/template"
	"net/http"
)

func UploadFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
	
	`

	t, err := template.New("form").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}
