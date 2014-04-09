package controllers

import (
	"log"
	"net/http"
	"html"
	"html/template"
	"github.com/stevenandrewcarter/gouter/models"
	"net/url"
)

// Indicates what the contents of a page should look like
type Page struct {
	Routes []models.Route
	MessageState string
	Message string
}

// Handles the Requests to the /config resource
func index(w http.ResponseWriter, r *http.Request) {
	message := ""
	messageState := ""
	log.Printf("Administration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
	// Check for a Delete request
	query := r.URL.Query()
	if len(query["name"]) != 0 && len(query["action"]) != 0 {
		log.Printf("Administration: Deleting '%v'", query["name"][0])
		// Delete the given route
		if models.Delete(query["name"][0]) {
			message = "Route succesfully deleted!"
			messageState = "success"
		} else {
			message = "Could not delete the route!"
			messageState = "warning"
		}
	}
	// Check if the request was a POST
	if r.Method == "POST" {
		r.ParseForm()
		if create(r.Form) {
			message = "Route succesfully created!"
			messageState = "success"
		} else {
			message = "Could not create the route. A route with the same name already exists!"
			messageState = "warning"
		}
	}
	result := models.Routes()
	page := &Page{Routes: result, Message: message, MessageState: messageState}
	t := template.Must(template.ParseGlob("tmpl/*.html"))
	// t, _ := template.ParseGlob("tmpl/*")
	t.ExecuteTemplate(w, "indexPage", page)
}

// Handles the POST request for the configuration handler
// url.Values will be a hash of string arrays. This seems strange, but that is what it is
func create(params url.Values) bool {
	log.Printf("Creating new route '%v', Params: (Description: '%v', From: '%v', To: '%v'", params["name"][0], params["description"][0], params["from"][0], params["to"][0])
	return models.Create(models.Route{params["description"][0], params["name"][0], params["from"][0], params["to"][0]})
}

// Handles the Requests to the /config/edit resource (Currently Not Implemented)
func update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Administration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
}

// Loads the Configuration Controller for Gouter
func Load() {
	log.Printf("Administration: Loading the configuration handlers.")
	// Handle the Controller actions
	http.HandleFunc("/admin", index)
	http.HandleFunc("/admin/edit", update)
	// Standard handling of the css, js and fonts paths
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./tmpl/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./tmpl/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./tmpl/fonts"))))
}
