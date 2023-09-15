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
    <!-- Inclua o CSS do Bootstrap -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
    <div class="container mt-5">
        <h1 class="text-center">Upload de Arquivo para o Servidor Go</h1>
        <form action="/upload" method="post" enctype="multipart/form-data" class="mt-4">
            <div class="mb-3">
                <input class="form-control" type="file" name="arquivo" id="arquivo" accept=".stl">
            </div>
            <button type="submit" class="btn btn-primary">Enviar Arquivo</button>
        </form>
    </div>
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
