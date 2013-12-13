package main

import "fmt"
import "net/http"
import "html"
import "log"
import (
    "code.google.com/p/gcfg"
)

type Configuration struct {
  Port string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
  log.Printf("Handling Request from: '%v", html.EscapeString(r.URL.Path))
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// Main entry point for the gouter project. Will listen on the configured port and route
// the http request to the matched request. The response will be returned to the original
// caller. 
func main() {  
  decoder := json.NewDecoder(os.Open("conf.json"))
  configuration := & Configuration{}
  decoder.Decode(&configuration)
  fmt.Println(configuration.Port)
  log.Printf("\n  ________               __                \n /  _____/  ____  __ ___/  |_  ___________ \n/   \\  ___ /  _ \\|  |  \\   __\\/ __ \\_  __ \\ \n\\    \\_\\  (  <_> )  |  /|  | \\  ___/|  | \\/\n \\______  /\\____/|____/ |__|  \\___  >__|\n        \\/                        \\/")
  log.Printf("Starting gouter v0.1. A simple HTTP router for RESTful API calls.")
  log.Printf("Please call http://localhost:%v/config to configure the router.")  
  http.HandleFunc("/", handleRequest)
  log.Printf("Listening for HTTP requests on Port '%v'", 1337)
  http.ListenAndServe(":1337", nil)  
}