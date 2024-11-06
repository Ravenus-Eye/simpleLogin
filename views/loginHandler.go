package views

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/daambrocio/simple_login/models"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var (
	// Configure the session store with a secret key
	sessionStore = sessions.NewCookieStore([]byte("your-secret-key"))
)

// comparePasswords compares a plain password with a hashed password
func comparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil // returns true if passwords match
}

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
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
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
		password := r.FormValue("password")

		// Retrieve user by username
		user, _ := models.GetUserByUsername(db, username)
		if user == nil {
			// User does not exist
			context["ValidationError"] = "Username not found."
			err = tmpl.Execute(w, context)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
			return
		}
		// Check if the provided password matches the stored password
		if !comparePasswords(user.Password, password) {
			// Password is incorrect
			context["ValidationError"] = "Incorrect password."
			err = tmpl.Execute(w, context)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
			return
		}

		// Successful login - create a session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error creating session", http.StatusInternalServerError)
			return
		}

		// Set user ID in session values
		session.Values["userID"] = user.ID
		session.Save(r, w)

		// Redirect to dashboard
		http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
	}
}
