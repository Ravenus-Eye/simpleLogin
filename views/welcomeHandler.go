package views

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/daambrocio/simple_login/models"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Define the path to the template file
	tmplPath := filepath.Join("templates", "welcome.html")

	// Parse the template file
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Retrieve the session
	session, err := sessionStore.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Error retrieving session", http.StatusInternalServerError)
		return
	}

	// Check if the user ID is set in the session
	userID, ok := session.Values["userID"].(int)
	if !ok {
		// If not logged in, redirect to login
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return
	}

	// Get user details from the database
	user, err := models.GetUserByID(db, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// If the user is admin, retrieve all user details
	adminDetails := []models.User{}
	if user.Username == "admin" {
		adminDetails = models.GetAllUsers(db)
	}

	// Prepare data to pass to the template
	context := map[string]interface{}{
		"User":         user,
		"AdminDetails": adminDetails,
		"Status":       "",
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
		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		username := r.FormValue("username")
		password := r.FormValue("password")
		active := r.FormValue("active")

		// Get User by ID
		chosenUser, _ := models.GetUserByID(db, id)
		if chosenUser.Name != name {
			err = models.UpdateUserDetails(db, id, "name", name)
			if err == nil {
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				return
			} else {
				fmt.Println("Error: name issue")
			}
		}
		if chosenUser.Username != username {
			err = models.UpdateUserDetails(db, id, "username", username)
			if err == nil {
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				return
			} else {
				fmt.Println("Error: username issue")
			}
		}
		if chosenUser.Password != password {
			err = models.UpdatePassword(db, username, password)
			if err == nil {
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				return
			} else {
				fmt.Println("Error: password issue")
			}
		}
		boolActive, _ := strconv.ParseBool(active)
		if chosenUser.Active != boolActive {
			err = models.UpdateUserDetails(db, id, "active", active)
			if err == nil {
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				return
			} else {
				fmt.Println("Error: active issue")
			}
		}
		return
	}
}
