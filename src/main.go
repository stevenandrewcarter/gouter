package main

import (
	"log"
	"html"
	"net/http"
	"fmt"
	"flag"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling Request from: '%v", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func loadSettings() (string) {
	wordPtr := flag.String("port", "8080", "Port Number for the Gouter to run at")
	flag.Parse()
	return *wordPtr
}

// Main entry point for the Gouter project. Will listen on the configured port and route
// the http request to the matched request. The response will be returned to the original
// caller. 
func main() {
	port := loadSettings()
	log.Printf("\n  ________               __                \n /  _____/  ____  __ ___/  |_  ___________ \n/   \\  ___ /  _ \\|  |  \\   __\\/ __ \\_  __ \\ \n\\    \\_\\  (  <_> )  |  /|  | \\  ___/|  | \\/\n \\______  /\\____/|____/ |__|  \\___  >__|\n        \\/                        \\/")
	log.Printf("Starting Gouter v0.1. A simple HTTP router for RESTful API calls.")
	log.Printf("Please call http://localhost:%v/config to configure Gouter.", port)
	http.HandleFunc("/", handleRequest)
	log.Printf("Listening for HTTP requests on Port '%v'", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
