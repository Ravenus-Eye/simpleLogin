package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

// Define a function to return JSON
func ReturnJSON(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// APIHandler !
func APIHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
	path := r.URL.Path

	switch path {
	case "changePassword":
		ChangePasshandler(w, r, db)
		return
	case "login":
		Login(w, r, db)
		return
	case "create_account":
		CreateAccount(w, r, db)
		return
	}
}
