package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates func
func LoadTemplates(pattern string) {
	LoggingInfoFile("Loading Templates...")
	templates = template.Must(template.ParseGlob(pattern))
}

// ExecuteTemplate func
func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

// ExecuteTemplate func
func ExecuteTemplateWithContentType(w http.ResponseWriter, tmpl string, data interface{}) {
	w.Header().Set("Content-Type", "application/pdf")
	templates.ExecuteTemplate(w, tmpl, data)
}
