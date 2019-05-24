package configs

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func getWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func deleteFile(path string) error {
	// delete file
	return os.Remove(path)
}

func TestConfig_InitEmptyPath(t *testing.T) {
	config := Config{}
	config.Init()
	if config.Path == "" {
		t.Error("Should create a default Path location if not provided")
	}
}

func TestConfig_InitProvidedPath(t *testing.T) {
	path := fmt.Sprintf("%s/t.yml", getWorkingDir())
	config := Config{
		Path: path,
	}
	config.Init()
	if config.Path != path {
		t.Error("Should use the provided Path location provided")
	}
	deleteFile(path)
}

func TestConfig_GetConfigSuccess(t *testing.T) {
	config := Config{}
	config.Init()
	_, err := config.GetConfig()
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}
