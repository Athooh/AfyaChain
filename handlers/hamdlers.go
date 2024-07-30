package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

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

func DoctorDashHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/doctors_dash.html"))
	tmpl.Execute(w, nil)
}

func PatientDashHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/patients_dash.html"))
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

// CreatePatientHandler handles the creation of a new patient
func CreatePatient(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	phone := r.FormValue("phone")
	email := r.FormValue("email")
	address := r.FormValue("address")
	gender := r.FormValue("gender")

	patient := models.Patient{
		FirstName: firstname,
		LastName:  lastname,
		DOB:       date,
		Phone:     phone,
		Email:     email,
		Address:   address,
		Gender:    gender,
	}
	db, err := database.ConnectDatabase()
	if err != nil {
		renderError(w, 500, "HTTP status 500 - Internal Server Error")
		return
	}
	result := db.Create(&patient)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td></tr>", patient.FirstName, patient.LastName)
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

	patient := database.GetPatient(num)
	if err != nil {
		http.Error(w, "Patient not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(patient)
}

func GetAllPatients(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return
	}
	var patients []models.Patient
	if err := db.Find(&patients).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	sendJSONResponse(w, patients, http.StatusOK)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
