package main

import (
	"flag"
	"fmt"
	"github.com/stevenandrewcarter/gouter/cmd/gouter"
	"github.com/stevenandrewcarter/gouter/cmd/gouter/controllers"
	"github.com/stevenandrewcarter/gouter/cmd/gouter/lib"
	"log"
	"net/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(-1)
	}
}

// RootCmd The Default command for running Gouter
var RootCmd = &cobra.Command{
	Use:   "gouter",
	Short: "Gouter is a simple proxy for routing traffic. It is designed to route dynamically.",
	Long:  "Gouter provides a solution for dynamically updating routing as and when needed.",
}

/*
 * Initialize the Cobra command handlers
 */
 func init() {
	RootCmd.AddCommand(ServerCmd)
	RootCmd.AddCommand(VersionCmd)
	ServerCmd.Flags().StringVarP(&elasticServerUrl,
		"elastic_url",
		"e",
		"http://localhost:9200",
		"The elastic server Url that Hydra will monitor")
	ServerCmd.Flags().IntVarP(&port,
		"port",
		"p",
		8080,
		"Port that Gouter will listen on")
	viper.BindPFlag("server.port", ServerCmd.Flags().Lookup("port"))
}
