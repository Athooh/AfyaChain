package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	blockchain "github.com/Athooh/HealthChain/Backend/blockChain"
	"github.com/Athooh/HealthChain/Backend/database"
	handler "github.com/Athooh/HealthChain/handlers"
	"github.com/Athooh/HealthChain/models"
)

func main() {

	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Auto Migrate the schema
	err = db.AutoMigrate(&models.Patient{}, &models.SignupForm{}, &blockchain.Blockchain{}, &blockchain.Block{}, &models.Facility{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	http.HandleFunc("/dum", handler.Dummy)
	http.HandleFunc("/patient", handler.CreatePatient) // Use POST for creating
	http.HandleFunc("/all", handler.GetAllPatients)    // Use POST for creating
	http.HandleFunc("/patient/", handler.GetPatientHandler)
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/facilitylog", handler.FacilityLogin)
	http.HandleFunc("/facilityReg", handler.RegisterFacility)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/admin", handler.Login)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/signup", handler.SignupHandler)
	http.HandleFunc("/signup/facility", handler.SignupFacilityHandler)
	http.HandleFunc("/facility/dashboard", handler.DoctorDashHandler)
	http.HandleFunc("/patient/dashboard", handler.PatientDashHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("Server started at http://localhost:8081")

	err1 := http.ListenAndServe(":8081", nil)
	if err1 != nil {
		fmt.Println("Error starting server:", err)
	}
}
