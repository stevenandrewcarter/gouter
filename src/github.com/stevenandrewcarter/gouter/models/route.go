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
func Routes() []Route {
	result := []Route{}
	fn := executeFunc(func(collection *mgo.Collection) {
		err := collection.Find(bson.M{}).All(&result)
		if err != nil {
			panic(err)
		}
	})
	execute(fn)
	return result
}

// Finds the document in the collection for the given document name
func Find(name string) Route {
	result := Route{}
	fn := executeFunc(func(collection *mgo.Collection) {
		err := collection.Find(bson.M{"name": name}).One(&result)
		if err != nil {
			result = Route{}
			log.Printf("Route '%v' not found. Error: %v", name, err)
		}
	})
	execute(fn)
	return result
}

func Create(route Route) bool {
	created := false
	checkRoute := Find(route.Name)
	if checkRoute.Name == "" {
		fn := executeFunc(func(collection *mgo.Collection) {
			err := collection.Insert(&route)
			if err != nil {
				panic(err)
			}
			log.Printf("Route created '%v'", route.Name)
			created = true
		})
		execute(fn)
	}
	return created
}

func Delete(name string) bool {
	deleted := false
	fn := executeFunc(func(collection *mgo.Collection) {
		err := collection.Remove(bson.M{"name": name})
		if err != nil {
			panic(err)
		}
		log.Printf("Route deleted '%v'", name)
		deleted = true
	})
	execute(fn)
	return deleted
}
