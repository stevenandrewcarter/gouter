package controllers

import (
	"log"
	"net/http"
	"html"
	"html/template"
	"models"
)

type Page struct {
	Routes []models.Route
}

// Handles the Requests to the /config resource
func handleConfiguration(w http.ResponseWriter, r *http.Request) {
	log.Printf("Configuration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
	// Check what type of request was made (GET / POST)
	if r.Method == "GET" {
		routes := make([]models.Route, 1)
		routes[0] = models.Route{ Source: "/test", Destination: "/test" }
		page := &Page{Routes: routes}
		t, _ := template.ParseFiles("tmpl/index.html")
		t.Execute(w, page)
	}
}

// Loads the Configuration Controller for Gouter
func Load() {
	log.Printf("Configuration: Loading the configuration handlers.")
	http.HandleFunc("/config", handleConfiguration)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./tmpl/css"))))
}
