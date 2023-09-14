package handlers

import (
	"html/template"
	"net/http"
)

func UploadFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Upload de Arquivo</title>
	</head>
	<body>
		<h1>Upload de Arquivo para o Servidor Go</h1>
		<form action="/upload" method="post" enctype="multipart/form-data">
			<input type="file" name="arquivo" id="arquivo" accept=".stl">
			<input type="submit" value="Enviar Arquivo">
		</form>
	</body>
	</html>
	`

	t, err := template.New("form").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}
