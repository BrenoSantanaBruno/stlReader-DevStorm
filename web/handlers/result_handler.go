package handlers

import (
	"html/template"
	"net/http"
	"stlReader-DevStorm/stl"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("arquivo")
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

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Resultado da Análise</title>
	</head>
	<body>
		<h1>Resultado da Análise</h1>
		<p>Número de Triângulos: {{.NumTriangles}}</p>
		<p>Área Total: {{.AreaTotal}}</p>
	</body>
	</html>
	`

	t, err := template.New("result").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"NumTriangles": numTriangles,
		"AreaTotal":    areaTotal,
	}

	t.Execute(w, data)
}
