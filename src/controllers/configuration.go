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
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Configuration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
	// Check what type of request was made (GET / POST)
	routes := models.Load()
	page := &Page{Routes: routes}
	t := template.Must(template.ParseGlob("tmpl/*.html"))
	// t, _ := template.ParseGlob("tmpl/*")
	t.ExecuteTemplate(w, "indexPage", page)
}

func create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Configuration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
	// Check what type of request was made (GET / POST)
	routes := models.Load()
	page := &Page{Routes: routes}
	t, _ := template.ParseFiles("tmpl/index.html")
	t.Execute(w, page)
}

func update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Configuration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
	// Check what type of request was made (GET / POST)
	routes := models.Load()
	page := &Page{Routes: routes}
	t, _ := template.ParseFiles("tmpl/index.html")
	t.Execute(w, page)
}

// Loads the Configuration Controller for Gouter
func Load() {
	log.Printf("Configuration: Loading the configuration handlers.")
	http.HandleFunc("/config", indexHandler)
	http.HandleFunc("/config/new", create)
	http.HandleFunc("/config/edit", update)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./tmpl/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./tmpl/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./tmpl/fonts"))))
}
