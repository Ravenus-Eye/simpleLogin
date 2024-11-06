package views

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/daambrocio/simple_login/models"
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

	if r.Method == http.MethodGet {
		// Serve the HTML file for GET requests
		err = tmpl.Execute(w, context)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get the values from the form
		username := r.FormValue("username")
		name := r.FormValue("name")
		password := r.FormValue("password")

		err = models.CreateLogin(db, username, name, password)
		if err == nil {
			// Redirect to login
			http.Redirect(w, r, "/login/", http.StatusSeeOther)
		}
	}
}
