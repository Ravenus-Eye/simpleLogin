package views

import (
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session
	session, err := sessionStore.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Error retrieving session", http.StatusInternalServerError)
		return
	}

	// Clear the session by setting MaxAge to -1 and saving it
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Error saving session", http.StatusInternalServerError)
		return
	}

	// Manually delete the cookie by setting its MaxAge to -1
	http.SetCookie(w, &http.Cookie{
		Name:     "session-name",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false, // Set to true if using HTTPS in production
	})

	// Redirect to the login page after logout
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
