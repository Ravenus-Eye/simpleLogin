package api

import (
	"database/sql"
	"net/http"

	"github.com/daambrocio/simple_login/models"
)

func ChangePasshandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
				ReturnJSON(w, map[string]interface{}{
					"status":  "error",
					"err_msg": "Cannot change Password.",
				})
				return
			} else {
				ReturnJSON(w, map[string]interface{}{
					"status": "ok",
					"msg":    "Password successfully changed.",
				})
				return
			}
		} else {
			ReturnJSON(w, map[string]interface{}{
				"status":  "error",
				"err_msg": "Password not matched.",
			})
			return
		}
	} else {
		ReturnJSON(w, map[string]interface{}{
			"status":  "error",
			"err_msg": "Username not found.",
		})
		return
	}
}
