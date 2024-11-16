package views

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func ChangePassHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Define the path to the template file
	tmplPath := filepath.Join("templates", "forgot_pw.html")

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

	fmt.Println("Method: Get")
	// Serve the HTML file for GET requests
	err = tmpl.Execute(w, context)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
