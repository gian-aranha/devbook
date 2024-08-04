package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates inserts the HTML templates into the templates variable
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecTemplate renders a HTML page
func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}