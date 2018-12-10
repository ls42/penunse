package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseGlob("templates/*.tmpl"))
var errorTemplate = `Error rendering template %s => %s`

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
