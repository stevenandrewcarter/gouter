package configs

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	li "github.com/stevenandrewcarter/gouter/internal/logging"
	"io/ioutil"
	"os"
)

var logger = li.NewLogger()

type Config struct {
	Path string
}

// Embedded Default Configuration File. This is used to generate a Config File if one does
// not already exist.
var defaultConfig = `
# Gouter Configuration
# ==============================
# Contains the configuration for the application.
# Configuration information includes the default port, database information, etc
---
# Server Configuration defines the default port that the application expects
server:
  port: 8080

# Routes define the middleware that the default application accepts
routes:
  - source: /admin
    destination: www.google.com
`

func getHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	return home
}

func (c *Config) createDefaultIfRequired() error {
	logger.Debugf("Checking if a configuration file exists at %s", c.Path)
	if _, err := os.Stat(c.Path); err == nil {
		// File exists, nothing more to do here
		return nil
	} else if os.IsNotExist(err) {
		// File does not exist, create a new one instead
		logger.Infof("Creating new default configuration at %s", c.Path)
		return ioutil.WriteFile(c.Path, []byte(defaultConfig), 0644)
	} else {
		// Something really weird is happening, just fail at this point
		return err
	}
}

// Build the path to the configuration, if one was not provided it will attempt to use
// $HOMEDIR/.gouter.yml
func (c *Config) buildPath() {
	if c.Path == "" {
		home := getHomeDir()
		c.Path = fmt.Sprintf("%s/.gouter.yml", home)
	}
}

func (c *Config) Init() {
	c.buildPath()
	if err := c.createDefaultIfRequired(); err != nil {
		// Could not read or create the configuration, this is a failure and will break
		logger.Errorf("Can't read config:", err)
		os.Exit(1)
	}
	logger.Infof("Loading config from %s", c.Path)
	viper.SetConfigFile(c.Path)
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Can't read config:", err)
		os.Exit(1)
	}
}

func (c *Config) GetConfig() (string, error) {
	content, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
