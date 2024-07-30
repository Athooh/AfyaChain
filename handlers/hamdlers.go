package handler

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Athooh/HealthChain/Backend/database"
	"github.com/Athooh/HealthChain/models"
)

type PageData struct {
	Text    string
	Art     string
	Error   string
	Code    int
	Message string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, 404, "HTTP status 404 - Page not found")
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		renderError(w, 500, "HTTP status 500 - Internal Server Error")
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/about.html"))
	tmpl.Execute(w, nil)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/signup.html"))
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}

func SignupFacilityHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/signup_facility.html"))
	tmpl.Execute(w, nil)
}

// func renderForm(w http.ResponseWriter, data PageData) {
// 	tmpl := template.Must(template.ParseFiles("templates/form.html"))
// 	err := tmpl.Execute(w, data)
// 	if err != nil {
// 		log.Printf("Error rendering form: %v", err)
// 		renderError(w, 500, "HTTP status 500 - Internal Server Error")
// 	}
// }

func renderError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	err := tmpl.Execute(w, PageData{Code: code, Message: message})
	if err != nil {
		log.Printf("Error rendering error page: %v", err)
		http.Error(w, "HTTP status 500 - Internal Server Error", http.StatusInternalServerError)
	}
}

var db *sql.DB

// SetDB sets the database connection for the handlers
func SetDB(databaseConnection *sql.DB) {
	db = databaseConnection
}

// HealthCheck checks the database connectivity
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	err := db.Ping()
	if err != nil {
		http.Error(w, "Database connection is not available", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database connection is healthy"))
}

// CreatePatientHandler handles the creation of a new patient
func CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := database.CreatePatient(db, &patient)
	if err != nil {
		http.Error(w, "Failed to create patient: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

// GetPatientHandler retrieves a patient by ID
func GetPatientHandler(w http.ResponseWriter, r *http.Request) {
	// Extract patient ID from the URL (e.g., /patient/1)
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing patient ID", http.StatusBadRequest)
		return
	}
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	patient, err := database.GetPatient(db, num)
	if err != nil {
		http.Error(w, "Patient not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(patient)
}
