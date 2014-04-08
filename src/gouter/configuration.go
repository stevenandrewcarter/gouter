package gouter

import (
	"code.google.com/p/gcfg"
)

type Config struct {
	Database struct {
		Host string
		Database string
	}
}

func Configuration() Config {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "config/app.gcfg")
	if err != nil {
		panic(err)
	}
	return cfg
}
