package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	log.Printf(`
  ________               __                
 /  _____/  ____  __ ___/  |_  ___________ 
/   \  ___ /  _ \|  |  \   __\/ __ \_  __ \ 
\    \_\  (  <_> )  |  /|  | \  ___/|  | \/
 \______  /\____/|____/ |__|  \___  >__|
        \/                        \/`)
	log.Printf("Starting Gouter v0.5. A simple HTTP router for RESTful API calls.")
	http.HandleFunc("/", lib.HandleRequest)
	log.Printf("Listening for HTTP requests on Port '%v'", port)
}

func Init() {
	log.Printf("Initializing config for server")
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
	if err := viper.BindPFlag("port", ServerCmd.Flags().Lookup("port")); err != nil {
		log.Print(err)
	}
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
