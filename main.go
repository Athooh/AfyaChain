package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Athooh/HealthChain/Backend/database"
	handler "github.com/Athooh/HealthChain/handlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}
	connStr := "user=afyadmin dbname=ehrdb password=pass host=localhost sslmode=disable"
	db, err := database.OpenDatabase(connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	handler.SetDB(db)

	http.HandleFunc("/health", handler.HealthCheck)
	http.HandleFunc("/patient", handler.CreatePatientHandler) // Use POST for creating
	http.HandleFunc("/patient/", handler.GetPatientHandler)
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/signup", handler.SignupHandler)
	http.HandleFunc("/signup/facility", handler.SignupFacilityHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("Server started at http://localhost:8080")

	err1 := http.ListenAndServe(":8080", nil)
	if err1 != nil {
		fmt.Println("Error starting server:", err)
	}
}
