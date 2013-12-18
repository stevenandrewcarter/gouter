package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"controllers"
	"router"
)

// Loads the parameters that got provided on the command line. If not provided will use the defaults instead
func loadParameters() (string) {
	wordPtr := flag.String("port", "8080", "Port Number for the Gouter to run at")
	flag.Parse()
	return *wordPtr
}

// Displays the start blurb for the router
func start(port string) {
	log.Printf("\n  ________               __                \n /  _____/  ____  __ ___/  |_  ___________ \n/   \\  ___ /  _ \\|  |  \\   __\\/ __ \\_  __ \\ \n\\    \\_\\  (  <_> )  |  /|  | \\  ___/|  | \\/\n \\______  /\\____/|____/ |__|  \\___  >__|\n        \\/                        \\/")
	log.Printf("Starting Gouter v0.1. A simple HTTP router for RESTful API calls.")
	log.Printf("Please call http://localhost:%v/config to configure Gouter.", port)
	http.HandleFunc("/", router.HandleRequest)
	controllers.Load()
	log.Printf("Listening for HTTP requests on Port '%v'", port)
}

// Main entry point for the Gouter project. Will listen on the configured port and models
// the http request to the matched request. The response will be returned to the original
// caller. 
func main() {
	port := loadParameters()
	start(port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
