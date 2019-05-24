package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	li "github.com/stevenandrewcarter/gouter/internal/logging"
	"github.com/stevenandrewcarter/gouter/internal/middleware"
	"net/http"
)

var logger = li.NewLogger()
var port int
var host string

/*
 * Starts the Server and opens the port for listening
 */
func start() {
	logger.Infof("Starting Gouter v0.5. A simple HTTP router for RESTful API calls.")
	m := middleware.Middleware{}
	http.HandleFunc("/", m.HandleRequest)
	logger.Infof("Listening for HTTP requests on Port '%v'", port)
}

func Init() {
	logger.Infof("Initializing config for server")
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
	logger.Infof("Binding settings from Config")
	if err := viper.BindPFlag("server.port", ServerCmd.Flags().Lookup("port")); err != nil {
		logger.Error(err)
	}
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the Gouter Server",
	Long:  "Starts up the Gouter Server and allows for proxy messages to be routed through it",
	Run: func(cmd *cobra.Command, args []string) {
		start()
		logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	},
}
