package controllers

import (
	"log"
	"net/http"
	"html"
	"html/template"
	"github.com/stevenandrewcarter/gouter/models"
)

// Indicates what the contents of a page should look like
type Page struct {
	Routes       []models.Route
	MessageState string
	Message      string
}

// Handles the Requests to the /config resource
func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("Administration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
	message, messageState := deleteRequest(r)
	message, messageState = createRequest(r)
	result := models.LoadRoutes()
	page := &Page{Routes: result, Message: message, MessageState: messageState}
	t := template.Must(template.ParseGlob("tmpl/*.html"))
	// t, _ := template.ParseGlob("tmpl/*")
	t.ExecuteTemplate(w, "indexPage", page)
}

// Handles the Requests to the /config/edit resource (Currently Not Implemented)
func update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Administration: Handling '%v' Request to: '%v", r.Method, html.EscapeString(r.URL.Path))
}

// Performs the delete request for the admin controller
func deleteRequest(r *http.Request) (string, string) {
	message := ""
	messageState := ""
	query := r.URL.Query()
	if len(query["name"]) != 0 && len(query["action"]) != 0 {
		log.Printf("Administration: Deleting '%v'", query["name"][0])
		if models.DeleteRoute(query["name"][0]) {
			message = "Route succesfully deleted!"
			messageState = "success"
		} else {
			message = "Could not delete the route!"
			messageState = "warning"
		}
	}
	return message, messageState
}

func createRequest(r *http.Request) (string, string) {
	message := ""
	messageState := ""
	// Check if the request was a POST
	if r.Method == "POST" {
		r.ParseForm()
		params := r.Form
		log.Printf("Administration: Creating new route '%v', Params: (Description: '%v', From: '%v', To: '%v'", params["name"][0], params["description"][0], params["from"][0], params["to"][0])
		if models.CreateRoute(models.Route{params["description"][0], params["name"][0], params["from"][0], params["to"][0]}) {
			message = "Route succesfully created!"
			messageState = "success"
		} else {
			message = "Could not create the route. A route with the same name already exists!"
			messageState = "warning"
		}
	}
	return message, messageState
}
