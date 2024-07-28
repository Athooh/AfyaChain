package handler

import (
	"html/template"
	"log"
	"net/http"
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
