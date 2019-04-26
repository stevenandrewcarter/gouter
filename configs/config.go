package configs

import (
	"io/ioutil"
	"log"
)

var defaultConfig = `
---
server:
  port: 8080
  url: "http://localhost:8080/admin"
`

func Create(path string) {
	err := ioutil.WriteFile(path, []byte(defaultConfig), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
