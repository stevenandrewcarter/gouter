package main

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevenandrewcarter/gouter/cmd/server"
	"github.com/stevenandrewcarter/gouter/cmd/version"
	"github.com/stevenandrewcarter/gouter/configs"
	li "github.com/stevenandrewcarter/gouter/internal/logging"
	"os"
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

var cfgFile string

/*
 * Initialize the Cobra command handlers
 */
func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default if $HOME/.gouter.yaml")
	RootCmd.AddCommand(server.ServerCmd)
	RootCmd.AddCommand(cmd.VersionCmd)
	server.Init()
}

func getHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	return home
}

func initConfig() {
	home := getHomeDir()
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		logger.Infof("Loading config from %s", cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".gouter" (without extension).
		logger.Infof("Loading config from %s", home)
		viper.AddConfigPath(home)
		viper.SetConfigName(".gouter")
	}

	if err := viper.ReadInConfig(); err != nil {
		// Attempt to write a default file
		configs.Create(home + "/.gouter.yaml")
		if err := viper.ReadInConfig(); err != nil {
			logger.Errorf("Can't read config:", err)
			os.Exit(1)
		}
	}
	logger.Infof("%s", viper.AllSettings())
}
