package config

import (
	"testing"
)

func TestRoutesLoadEmptyContent(t *testing.T) {
	routes := Routes{}
	if err := routes.Load(""); err != nil {
		t.Error("unexpected error when attempting to load empty content")
	}
	if len(routes.Routes) > 0 {
		t.Error("should have no active routes")
	}
}

func TestRoutesLoadInvalidConfig(t *testing.T) {
	routes := Routes{}
	if err := routes.Load("this_is_not_valid_yaml"); err == nil {
		t.Error("expected an error when attempting to load an empty path")
	}
}

func TestRoutesLoadSuccess(t *testing.T) {
	routeYaml := `
routes:
  - source: /admin
    destination: www.google.com
`
	routes := Routes{}
	if err := routes.Load(routeYaml); err != nil {
		t.Errorf("expected a successful load call %s", err)
	}
	if len(routes.Routes) == 0 {
		t.Error("expected to load at least one route")
	}
}
