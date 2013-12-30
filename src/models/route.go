package models

import (
	"strings"
	"encoding/json"
	"log"
	"io"
	"io/ioutil"
)

type Route struct {
	Description string
	Name        string
	From        string
	To          string
}

func Load() []Route {
	routes := []Route{}
	// Load the Json from the file
	jsonStream, _ := ioutil.ReadFile("data/routes.json")
	dec := json.NewDecoder(strings.NewReader(string(jsonStream)))
	for {
		var r Route
		if err := dec.Decode(&r); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		routes = append(routes, r)
	}
	return routes
}

func Save() {

}
