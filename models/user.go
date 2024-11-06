package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

func (e *User) FullName() string {
	return e.Name
}

// CreateLogin inserts new login entry to users db
func CreateLogin(db *sql.DB, username, name, password string) error {
	// Insert the new user into the database
	query := `INSERT INTO users (username, name, password, active) VALUES (?, ?, ?, ?)`
	// Execute the update query with parameters
	hashedPassword, _ := HashPassword(password)

	if hashedPassword != "" {
		_, err := db.Exec(query, username, name, hashedPassword, true)
		if err != nil {
			return fmt.Errorf("could not insert user: %v", err)
		}
	} else {
		fmt.Println("Cant hash password.")
	}
	fmt.Printf("Login successfully created %s\n", username)
	return nil
}

// GetAllUsers retrieves all the user available
func GetAllUsers(db *sql.DB) []User {
	users := []User{}
	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Active); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

// GetUserByLogin retrieves a user by their username and password
func GetUserByLogin(db *sql.DB, username string, password string) (*User, error) {
	user := &User{}
	query := "SELECT * FROM users WHERE username = ? AND password = ? AND active = ?"
	err := db.QueryRow(query, username, password, true).Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Active)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	user := &User{}
	query := "SELECT * FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Active)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(db *sql.DB, id int) (*User, error) {
	user := &User{}
	query := "SELECT * FROM users WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Active)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func UpdatePassword(db *sql.DB, username, password string) error {
	query := `UPDATE users SET password = ? WHERE username = ?`

	// Execute the update query with parameters
	hashedPassword, _ := HashPassword(password)
	if hashedPassword != "" {

		result, err := db.Exec(query, hashedPassword, username)
		if err != nil {
			return fmt.Errorf("could not update password: %v", err)
		}

		// Check how many rows were affected
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("could not retrieve rows affected: %v", err)
		}

		if rowsAffected == 0 {
			return fmt.Errorf("no user found with the username %s", username)
		}

	} else {
		fmt.Println("Cant hash password.")
	}
	fmt.Printf("Password updated successfully for user %s\n", username)
	return nil
}

func UpdateUserDetails(db *sql.DB, id int, key, value string) error {
	query := `UPDATE users SET ` + key + ` = ? WHERE id = ?`
	if key == "active" {
		value, _ := strconv.ParseBool(value)
		_, err := db.Exec(query, value, id)
		if err != nil {
			return fmt.Errorf("could not update active: %v", err)
		}
	} else {
		_, err := db.Exec(query, value, id)
		if err != nil {
			return fmt.Errorf("could not update active: %v", err)
		}
	}
	return nil
}
