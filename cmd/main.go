package main

import (
	"github.com/spf13/cobra"
	"github.com/stevenandrewcarter/gouter/cmd/server"
	"github.com/stevenandrewcarter/gouter/cmd/version"
	"github.com/stevenandrewcarter/gouter/configs"
	li "github.com/stevenandrewcarter/gouter/internal/logging"
)

var logger = li.NewLogger()

func main() {
	if err := RootCmd.Execute(); err != nil {
		logger.Fatalf(err.Error())
	}
}

// RootCmd The Default command for running Gouter. All of the other commands will be based off the Root Command.
var RootCmd = &cobra.Command{
	Use:   "gouter",
	Short: "Gouter is a simple API Manager for routing traffic.",
	Long: `
			Gouter provides the basic functionality for a simple API Manager. It should be able to handle a fair amount 
			of traffic and should be hopefully fairly easy to use. You should be able to configure most of the features
			via the flags (Or the configuration file).
		   `,
}

var config configs.Config

/*
 * Initialize the Cobra command handlers
 */
func init() {
	cobra.OnInitialize(config.Init)
	RootCmd.PersistentFlags().StringVar(&config.Path, "config", "", "config file (default if $HOME/.gouter.yaml")
	RootCmd.AddCommand(server.ServerCmd)
	RootCmd.AddCommand(cmd.VersionCmd)
	server.Init()
}
