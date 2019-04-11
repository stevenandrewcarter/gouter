package controllers

import (	
	"github.com/stevenandrewcarter/gouter/configs"
	"log"
	"net/http"
)

// Loads the route handers for configuration for the go router
func Load() {
	log.Printf("Administration: Loading the configuration handlers.")
	// Handle the Controller actions
	http.HandleFunc(gouter.Configuration().Application.AdminUrl, cmd.index)
	// Standard handling of the css, js and fonts paths
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./tmpl/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./tmpl/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./tmpl/fonts"))))
	http.Handle("/favicon.ico", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./tmpl/images"))))
}
