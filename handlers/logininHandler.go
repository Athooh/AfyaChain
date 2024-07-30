package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/Athooh/HealthChain/Backend/database"
	"github.com/Athooh/HealthChain/models"
	"github.com/jinzhu/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials

	// Check Content-Type
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Print the raw JSON data
	log.Printf("Raw JSON request body: %s\n", body)

	// Decode JSON request body
	err = json.Unmarshal(body, &creds) // Use Unmarshal since we already read the body
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	fmt.Printf("Decoded credentials: %+v\n", creds)

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	// Query the database for the user
	var user models.Credentials // Assuming you have a model defined for credentials
	if err := db.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response := map[string]string{"message": "Invalid credentials"}
			log.Println("Login not successful: User not found")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}
		log.Printf("Database query failed: %v", err)
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}

	// Compare the input password and auth key with stored values
	if creds.Password == user.Password && creds.AuthKey == user.AuthKey {
		response := map[string]string{"message": "Login successful"}
		log.Println("Login successful")
		w.WriteHeader(http.StatusOK) // Set status code to 200 OK
		json.NewEncoder(w).Encode(response)
	} else {
		response := map[string]string{"message": "Invalid credentials"}
		log.Println("Login not successful: Incorrect password or auth key")
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

	// Check Content-Type
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		http.Error(w, "Content-Type must be application/x-www-form-urlencoded", http.StatusUnsupportedMediaType)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Retrieve form values
	creds.Username = r.FormValue("sname")
	creds.Password = r.FormValue("passwd")
	creds.AuthKey = r.FormValue("auth")
	creds.UserType = r.FormValue("type")

	// Debugging information
	log.Println("Decoded credentials:", creds)

	// Connect to the MySQL database
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	db.AutoMigrate(&models.Credentials{})

	// Insert the new user into the database
	result := db.Create(&creds)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	response := map[string]string{"message": "User registered successfully"}
	w.WriteHeader(http.StatusCreated) // Set status code to 201 Created
	json.NewEncoder(w).Encode(response)
}

func Dummy(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dummy.html"))
	tmpl.Execute(w, nil)
}
