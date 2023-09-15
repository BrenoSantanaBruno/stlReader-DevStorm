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
    <title>Analysis Result</title>
    <!-- Bootstrap 5 CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
<div class="container mt-5">
    <h1 class="text-center">Analysis Result</h1>
    <div class="result-container mt-4 p-4">
        <p>Numbers of Triangles: {{.NumTriangles}}</p>
        <p>Total Area: {{.AreaTotal}}</p>
    </div>
</div>
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
