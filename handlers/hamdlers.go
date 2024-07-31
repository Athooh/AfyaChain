package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	blockchain "github.com/Athooh/HealthChain/Backend/blockChain"
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

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	var signupData models.SignupForm

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

	// Retrieve form values into the struct
	signupData.FirstName = r.FormValue("firstname")
	signupData.LastName = r.FormValue("lastname")
	signupData.Phone = r.FormValue("phone")
	signupData.Password = r.FormValue("password")
	signupData.ConfirmPassword = r.FormValue("confirm-password")
	signupData.Sex = r.FormValue("sex")
	signupData.Country = r.FormValue("country")
	signupData.City = r.FormValue("city")
	fmt.Println(signupData.Password)

	// Check if password and confirm password match
	if signupData.Password != signupData.ConfirmPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}
	difficulty := 4 // Example difficulty value; adjust as needed

	// Connect to the database
	db, err := database.ConnectDatabase()
	if err != nil {
		renderError(w, http.StatusInternalServerError, "HTTP status 500 - Internal Server Error")
		return
	}

	// Save the patient record to the database
	if err := db.Create(&signupData).Error; err != nil {
		renderError(w, http.StatusInternalServerError, "HTTP status 500 - Internal Server Error while inserting data")
		return
	}

	// Create the initial blockchain with a genesis block for the new patient
	genesisBlock := blockchain.Block{
		PatientID:    0,   // No patient ID for the genesis block
		UserID:       123, // Link to the patient
		Action:       "genesis",
		Timestamp:    time.Now(),
		PreviousHash: "0",
		Hash:         "", // Will be calculated later
		Pow:          0,  // Proof of work (if used)
	}

	// Calculate the hash for the genesis block
	genesisBlock.Hash = genesisBlock.CalculateHash()

	// Create the user blockchain
	userBlockchain := blockchain.Blockchain{
		UserID:     12,
		Difficulty: difficulty,
		Chain:      []blockchain.Block{genesisBlock}, // Initialize with the genesis block
	}

	// Save the user blockchain to the database
	if err := db.Create(&userBlockchain).Error; err != nil {
		renderError(w, http.StatusInternalServerError, "HTTP status 500 - Internal Server Error while inserting user blockchain")
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "created succesfuly",
	}
	json.NewEncoder(w).Encode(response)
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
