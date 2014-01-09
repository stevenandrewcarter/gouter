package controllers

import (
	"log"
	"net/http"
	"html"
	"html/template"
	"models"
	"net/url"
)

// Indicates what the contents of a page should look like
type Page struct {
	Routes []models.Route
}

// Handles the Requests to the /config resource
func index(w http.ResponseWriter, r *http.Request) {
	result := models.Routes()
	log.Printf("Configuration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
	// Check if the request was a POST
	if r.Method == "POST" {
		r.ParseForm()
		create(r.Form)
	}
	page := &Page{Routes: result}
	t := template.Must(template.ParseGlob("tmpl/*.html"))
	// t, _ := template.ParseGlob("tmpl/*")
	t.ExecuteTemplate(w, "indexPage", page)
}

func create(params url.Values) {
	log.Printf("Params: '%v'", params["name"])
}

// Handles the Requests to the /config/edit resource (Currently Not Implemented)
func update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Configuration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
}

// Loads the Configuration Controller for Gouter
func Load() {
	log.Printf("Configuration: Loading the configuration handlers.")
	// Handle the Controller actions
	http.HandleFunc("/config", index)
	http.HandleFunc("/config/edit", update)
	// Standard handling of the css, js and fonts paths
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./tmpl/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./tmpl/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./tmpl/fonts"))))
}
