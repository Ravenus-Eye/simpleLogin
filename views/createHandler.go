package views

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Define the path to the template file
	tmplPath := filepath.Join("templates", "register.html")

	// Parse the template file
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Initialize the context for template data
	context := map[string]interface{}{
		"ValidationError": "",
	}

	// Serve the HTML file for GET requests
	err = tmpl.Execute(w, context)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
