package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"log"
)

/*
 * Starts the Server and opens the port for listening
 */
func start(port int) {
	log.Printf("\n  ________               __                \n /  _____/  ____  __ ___/  |_  ___________ \n/   \\  ___ /  _ \\|  |  \\   __\\/ __ \\_  __ \\ \n\\    \\_\\  (  <_> )  |  /|  | \\  ___/|  | \\/\n \\______  /\\____/|____/ |__|  \\___  >__|\n        \\/                        \\/")
	log.Printf("Starting Gouter v0.5. A simple HTTP router for RESTful API calls.")
	log.Printf("Please call http://localhost:%v%v to configure Gouter.", port, gouter.Configuration().Application.AdminUrl)
	http.HandleFunc("/", lib.HandleRequest)
	controllers.Load()
	log.Printf("Listening for HTTP requests on Port '%v'", port)
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the Gouter Server",
	Long:  "Starts up the Gouter Server and allows for proxy messages to be routed through it",
	Run: func(cmd *cobra.Command, args []string) {
		start(util.Config.GetPort())
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", util.Config.GetPort()), nil))
	},
}
