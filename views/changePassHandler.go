package views

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/daambrocio/simple_login/models"
)

func ChangePassHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	path := filepath.Join("templates", "forgot_pw.html") // path to index.html
	if r.Method == http.MethodGet {
		// Serve the HTML file for GET requests
		http.ServeFile(w, r, path)
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
		user, _ := models.GetUserByUsername(db, username)
		if user != nil {
			password1 := r.FormValue("password1")
			password2 := r.FormValue("password2")
			if password1 == password2 {
				err := models.UpdatePassword(db, username, password1)
				if err != nil {
					fmt.Println("error: ", err)
				} else {
					http.Redirect(w, r, "/login", http.StatusSeeOther)
				}
			}
		}
	}

}
