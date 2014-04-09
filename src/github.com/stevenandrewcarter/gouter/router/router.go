package router

import (
	"log"
	"net/http"
	"html"
	"fmt"
)

// Handles the requests to the router
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling Request from: '%v", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
