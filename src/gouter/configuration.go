package gouter

import (
	"code.google.com/p/gcfg"
)

// Configuration structure that is loaded from the application configuration file
type Config struct {
	Database struct {
		Host string
		Database string
	}
}

// Loads the application configuration, will cause a panic if the configuration cannot be loaded
func Configuration() Config {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "config/app.gcfg")
	if err != nil {
		panic(err)
	}
	return cfg
}
