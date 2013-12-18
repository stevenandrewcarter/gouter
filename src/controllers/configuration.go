package controllers

import (
	"fmt"
	"log"
	"net/http"
	"html"
)

func HandleConfiguration(w http.ResponseWriter, r *http.Request) {
	log.Printf("Configuration Request from: '%v", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
