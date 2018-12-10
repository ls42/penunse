package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseGlob("templates/*.tmpl"))
var errorTemplate = `
<html>
	<body>
		<h1>Error rendering template %s</h1>
		<p>%s</p>
	</body>
</html>
`

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	err := t.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf(errorTemplate, name, err),
			http.StatusInternalServerError,
		)
	}
}
