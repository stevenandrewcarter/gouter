package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevenandrewcarter/gouter/cmd/server"
	"github.com/stevenandrewcarter/gouter/cmd/version"
	"github.com/stevenandrewcarter/gouter/configs"
	"log"
	"os"
)

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf(err.Error())
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
	RootCmd.AddCommand(server.ServerCmd)
	RootCmd.AddCommand(cmd.VersionCmd)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default if $HOME/.gouter.yaml")
	server.Init()
	if err := viper.BindPFlag("server.port", server.ServerCmd.Flags().Lookup("port")); err != nil {
		log.Print(err)
	}
}

func getHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home
}

func initConfig() {
	home := getHomeDir()
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".gouter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gouter")
	}

	if err := viper.ReadInConfig(); err != nil {
		// Attempt to write a default file
		configs.Create(home + "/.gouter.yaml")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}
	}
}
