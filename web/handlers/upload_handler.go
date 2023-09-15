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
       <!-- Bootstrap 5 CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
		<title>File Upload</title>
	</head>
	<body>
		<div class="container mt-5">
			<h1 class="text-center">Upload STL File</h1>
    	<div class="result-container mt-4 p-4">
				
			<form action="/upload" method="post" enctype="multipart/form-data">
				<input type="file" name="arquivo" id="arquivo" accept=".stl">
				<input type="submit" value="Analysis">
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
