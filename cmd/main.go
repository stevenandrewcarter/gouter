package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevenandrewcarter/gouter/cmd/server"
	"github.com/stevenandrewcarter/gouter/cmd/version"
	"log"
	"os"
)

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf(err.Error())
		os.Exit(-1)
	}
}

// RootCmd The Default command for running Gouter
var RootCmd = &cobra.Command{
	Use:   "gouter",
	Short: "Gouter is a simple proxy for routing traffic. It is designed to route dynamically.",
	Long:  "Gouter provides a solution for dynamically updating routing as and when needed.",
}

var port = 8080
var url = "http://localhost"

/*
 * Initialize the Cobra command handlers
 */
 func init() {
	RootCmd.AddCommand(server.ServerCmd)
	RootCmd.AddCommand(cmd.VersionCmd)
	server.ServerCmd.Flags().StringVarP(&url,
		"elastic_url",
		"e",
		"http://localhost:9200",
		"The elastic server Url that Hydra will monitor")
	server.ServerCmd.Flags().IntVarP(&port,
		"port",
		"p",
		8080,
		"Port that Gouter will listen on")
	viper.BindPFlag("server.port", server.ServerCmd.Flags().Lookup("port"))
}
