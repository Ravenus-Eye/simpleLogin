package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/daambrocio/simple_login/views"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := Database()
	CheckErr(err)
	_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS user_db;")
	CheckErr(err)
	_, err = db.Exec("USE user_db;")
	CheckErr(err)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), username VARCHAR(255) UNIQUE, password VARCHAR(255), active BOOLEAN)")
	CheckErr(err)
	fmt.Println("Users table initialized")
	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS enty_login (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), password VARCHAR(255))")
	// CheckErr(err)
	// fmt.Println("Users table initialized")
	_, err = db.Exec("INSERT INTO users (name, username, password, active) SELECT 'admin', 'admin', 'admin', True FROM DUAL WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = 'admin')")
	fmt.Println("User Admin initialized")
	CheckErr(err)
	fmt.Println("start program")
	defer db.Close()

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/login/", func(w http.ResponseWriter, r *http.Request) {
		views.LoginHandler(w, r, db)
	})
	http.HandleFunc("/create_login/", func(w http.ResponseWriter, r *http.Request) {
		views.CreateHandler(w, r, db)
	})
	http.HandleFunc("/change_password/", func(w http.ResponseWriter, r *http.Request) {
		views.ChangePassHandler(w, r, db)
	})
	http.HandleFunc("/dashboard/", func(w http.ResponseWriter, r *http.Request) {
		views.WelcomeHandler(w, r, db)
	})
	http.HandleFunc("/logout/", views.LogoutHandler)

	log.Println("Server started on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func Database() (db *sql.DB, err error) {
	// Open a connection to the database
	db, err = sql.Open("mysql", "{db Username}:{db Password}@tcp(127.0.0.1:3306)/portal")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	} else {
		fmt.Println("Successfully connected to database.")
	}
	return db, err
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
