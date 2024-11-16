package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/daambrocio/simple_login/models"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var (
	// Configure the session store with options
	sessionStore = sessions.NewCookieStore([]byte("your-secret-key"))
)

func init() {
	sessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 1 week
		HttpOnly: true,
		Secure:   false,                // Set to true if running on HTTPS
		SameSite: http.SameSiteLaxMode, // Allows the cookie for same-site requests
	}
}

func comparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if err := r.ParseForm(); err != nil {
		ReturnJSON(w, map[string]interface{}{
			"status":  "error",
			"err_msg": "Unable to parse form",
		})
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := models.GetUserByUsername(db, username)
	if err != nil || user == nil {
		ReturnJSON(w, map[string]interface{}{
			"status":  "error",
			"err_msg": "Username not found",
		})
		return
	}

	if !comparePasswords(user.Password, password) {
		ReturnJSON(w, map[string]interface{}{
			"status":  "error",
			"err_msg": "Incorrect password",
		})
		return
	}

	// Retrieve or create a session
	session, err := sessionStore.Get(r, "session-name")
	if err != nil {
		ReturnJSON(w, map[string]interface{}{
			"status":  "error",
			"err_msg": "Error creating session",
		})
		return
	}

	// Set session value
	session.Values["userID"] = user.ID

	// Save session and check for errors
	if err := session.Save(r, w); err != nil {
		log.Println("Error saving session:", err)
		ReturnJSON(w, map[string]interface{}{
			"status":  "error",
			"err_msg": "Error saving session",
		})
		return
	}

	// Confirm successful login
	ReturnJSON(w, map[string]interface{}{
		"status":  "success",
		"message": "Login successful",
	})
}
