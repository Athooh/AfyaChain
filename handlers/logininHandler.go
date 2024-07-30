package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthKey  string `json:"authkey"`
	UserType string `json:"usertype"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	fmt.Println("Passed")
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
	// Dummy authentication logic; replace with actual database lookup
	if creds.Username == "admin" && creds.Password == "password" && creds.AuthKey == "@doc1234" && creds.UserType == "patient" {
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
