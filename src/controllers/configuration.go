package controllers

import (
	"log"
	"net/http"
	"html"
	"html/template"
)

// Handles the Requests to the /config resource
func handleConfiguration(w http.ResponseWriter, r *http.Request) {
	log.Printf("Configuration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
	// Check what type of request was made (GET / POST)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("tmpl/index.html")
		t.Execute(w, t)
	}
}

// Loads the Configuration Controller for Gouter
func Load() {
	log.Printf("Configuration: Loading the configuration handlers.")
	http.HandleFunc("/config", handleConfiguration)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./tmpl/css"))))
}
