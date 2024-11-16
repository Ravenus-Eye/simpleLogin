package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/daambrocio/simple_login/models"
)

func CreateAccount(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the values from the form
	username := r.FormValue("username")
	name := r.FormValue("name")
	password := r.FormValue("password")
	if username == "" || name == "" || password == "" {
		ReturnJSON(w, map[string]interface{}{
			"status":  "error",
			"err_msg": "Please fill up the form completely.",
		})
	} else {
		user, _ := models.GetUserByUsername(db, username)
		if user != nil {
			fmt.Println("user: ", user)
			ReturnJSON(w, map[string]interface{}{
				"status":  "error",
				"err_msg": "Username already taken.",
			})
		} else {
			err := models.CreateLogin(db, username, name, password)
			if err == nil {
				ReturnJSON(w, map[string]interface{}{
					"status":  "success",
					"err_msg": "Account uccessfully created.",
				})
			} else {
				ReturnJSON(w, map[string]interface{}{
					"status":  "error",
					"err_msg": "Failed to create account.",
				})
			}
		}
	}

}
