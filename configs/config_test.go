package configs

import (
	"fmt"
	"testing"
)

func TestConfig_InitEmptyPath(t *testing.T) {
	config := Config{}
	config.Init()
	if config.Path == "" {
		t.Error("Should create a default Path location if not provided")
	}
}

func TestConfig_InitProvidedPath(t *testing.T) {
	path := fmt.Sprintf("%s/t.yml", getHomeDir())
	config := Config{
		Path: path,
	}
	config.Init()
	if config.Path != path {
		t.Error("Should use the provided Path location provided")
	}
}

func TestConfig_GetConfigSuccess(t *testing.T) {
	config := Config{}
	config.Init()
	_, err := config.GetConfig()
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}
