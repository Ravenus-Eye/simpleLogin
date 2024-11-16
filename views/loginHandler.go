package views

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/sessions"
)

var (
	// Configure the session store with a secret key
	sessionStore = sessions.NewCookieStore([]byte("your-secret-key"))
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Retrieve the session
	session, err := sessionStore.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Error retrieving session", http.StatusInternalServerError)
		return
	}

	// Check if user is already logged in
	if _, ok := session.Values["userID"]; ok {
		// If userID exists in session, redirect to dashboard
		http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
		return
	}

	// Define the path to the template file
	tmplPath := filepath.Join("templates", "login.html")

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
