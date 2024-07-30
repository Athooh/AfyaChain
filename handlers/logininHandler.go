package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Athooh/HealthChain/Backend/database"
	"github.com/Athooh/HealthChain/models"
	_ "github.com/go-sql-driver/mysql"
)

// Import the MySQL driver

// Credentials represents the user login credentials

func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Connect to the MySQL database
	db, err := sql.Open("mysql", "new_username:new_password@tcp(127.0.0.1:3306)/afya_chain_db") // Update with your database credentials
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query the database for the user
	var storedPassword string
	var storedAuthKey string
	var storedUserType string

	query := "SELECT password, authkey, usertype FROM users WHERE username = ?"
	err = db.QueryRow(query, creds.Username).Scan(&storedPassword, &storedAuthKey, &storedUserType)
	if err != nil {

		if err == sql.ErrNoRows {
			response := map[string]string{"message": "Invalid credentials"}
			fmt.Println("Login not successful")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}
		fmt.Println(err)
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}

	// Compare the input password and auth key with stored values
	if creds.Password == storedPassword && creds.AuthKey == storedAuthKey && creds.UserType == storedUserType {
		response := map[string]string{"message": "Login successful"}
		fmt.Println("Login successful")
		w.WriteHeader(http.StatusOK) // Set status code to 200 OK
		json.NewEncoder(w).Encode(response)
	} else {
		response := map[string]string{"message": "Invalid credentials"}
		fmt.Println("Login not successful")
		w.WriteHeader(http.StatusUnauthorized) // Set status code to 401 Unauthorized
		json.NewEncoder(w).Encode(response)
	}
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}
func Register(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	fmt.Println("Passed")
	// Connect to the MySQL database
	db, err := database.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	// Insert the new user into the database
	query := "INSERT INTO users (username, password, auth_key, user_type) VALUES (?, ?, ?, ?)"
	db.Exec(query, creds.Username, creds.Password, creds.AuthKey, creds.UserType)
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "Failed to register user", http.StatusInternalServerError)
	// 	return
	// }

	response := map[string]string{"message": "User registered successfully"}
	w.WriteHeader(http.StatusCreated) // Set status code to 201 Created
	json.NewEncoder(w).Encode(response)
}
func Dummy(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dummy.html"))
	tmpl.Execute(w, nil)
}
