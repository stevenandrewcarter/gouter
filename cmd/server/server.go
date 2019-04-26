package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stevenandrewcarter/gouter/controllers"
	"github.com/stevenandrewcarter/gouter/lib"
	"log"
	"net/http"
)

var port int
var url string
var host string

/*
 * Starts the Server and opens the port for listening
 */
func start() {
	log.Printf("\n  ________               __                \n /  _____/  ____  __ ___/  |_  ___________ \n/   \\  ___ /  _ \\|  |  \\   __\\/ __ \\_  __ \\ \n\\    \\_\\  (  <_> )  |  /|  | \\  ___/|  | \\/\n \\______  /\\____/|____/ |__|  \\___  >__|\n        \\/                        \\/")
	log.Printf("Starting Gouter v0.5. A simple HTTP router for RESTful API calls.")
	log.Printf("Please call %s:%v%s to configure Gouter.", host, port, url)
	http.HandleFunc("/", lib.HandleRequest)
	controllers.Load(url)
	log.Printf("Listening for HTTP requests on Port '%v'", port)
}

func Init() {
	ServerCmd.Flags().StringVarP(&url,
		"url",
		"e",
		"/admin",
		"Admin URL for Gouter")
	ServerCmd.Flags().IntVarP(&port,
		"port",
		"p",
		8080,
		"Port that Gouter will listen on")
	ServerCmd.Flags().StringVarP(&host,
		"host",
		"a",
		"http://localhost",
		"Host that Gouter is running on")
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the Gouter Server",
	Long:  "Starts up the Gouter Server and allows for proxy messages to be routed through it",
	Run: func(cmd *cobra.Command, args []string) {
		start()
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	},
}
