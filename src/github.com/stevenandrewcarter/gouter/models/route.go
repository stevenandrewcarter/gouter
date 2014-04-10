package models

import (
	"log"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// Structure that the routes will be represented by
type Route struct {
	Description string
	Name        string
	From        string
	To          string
}

// Load all of the routes from the database
func LoadRoutes() []Route {
	result := []Route{}
	fn := executeFunc(func(collection *mgo.Collection) {
		err := collection.Find(bson.M{}).All(&result)
		if err != nil { panic(err) }
	})
	execute(fn, ROUTE_COLLECTION)
	return result
}

// Finds the document in the collection for the given document name
// Returns the route if found, otherwise returns a blank route structure
func FindRoute(name string) Route {
	result := Route{}
	fn := executeFunc(func(collection *mgo.Collection) {
		err := collection.Find(bson.M{"name": name}).One(&result)
		if err != nil {
			result = Route{}
			log.Printf("Route '%v' not found. Error: %v", name, err)
		}
	})
	execute(fn, ROUTE_COLLECTION)
	return result
}

// Create a route from the provided route structure
// Returns true if the route was created successfully, or false if it already exists
func CreateRoute(route Route) bool {
	created := false
	checkRoute := FindRoute(route.Name)
	if checkRoute.Name == "" {
		fn := executeFunc(func(collection *mgo.Collection) {
			err := collection.Insert(&route)
			if err != nil { panic(err) }
			log.Printf("Route created '%v'", route.Name)
			created = true
		})
		execute(fn, ROUTE_COLLECTION)
	}
	return created
}

// Delete the route from the database with the given route name
// Returns a true if successful or false if the route could not be found
func DeleteRoute(name string) bool {
	deleted := false
	fn := executeFunc(func(collection *mgo.Collection) {
		err := collection.Remove(bson.M{"name": name})
		if err != nil { panic(err) }
		log.Printf("Route deleted '%v'", name)
		deleted = true
	})
	execute(fn, ROUTE_COLLECTION)
	return deleted
}
