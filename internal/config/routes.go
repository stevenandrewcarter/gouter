package config

import (
	"gopkg.in/yaml.v2"
)

type Routes struct {
	Routes []Route `yaml:"routes"`
}

// Structure that the routes will be represented by
type Route struct {
	Description string `yaml:"description"`
	Name        string `yaml:"name"`
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

// // Load all of the routes from the database
func (r *Routes) Load(config string) error {
	return yaml.Unmarshal([]byte(config), r)
}

// // Finds the document in the collection for the given document name
// // Returns the route if found, otherwise returns a blank route structure
// func FindRoute(name string) Route {
// 	result := Route{}
// 	fn := cmd.executeFunc(func(collection *mgo.Collection) {
// 		err := collection.Find(bson.M{"name": name}).One(&result)
// 		if err != nil {
// 			result = Route{}
// 			log.Printf("Route '%v' not found. Error: %v", name, err)
// 		}
// 	})
// 	cmd.execute(fn, cmd.ROUTE_COLLECTION)
// 	return result
// }

// func FindRouteByFrom(from string) (Route, error) {
// 	result := Route{}
// 	var err error = nil
// 	fn := cmd.executeFunc(func(collection *mgo.Collection) {
// 		err = collection.Find(bson.M{"from": from}).One(&result)
// 		if err != nil {
// 			log.Printf("Route '%v' not found. Error: %v", from, err)
// 		}
// 	})
// 	cmd.execute(fn, cmd.ROUTE_COLLECTION)
// 	return result, err
// }
