package models

import (
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
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("t2_dev").C("services_models_services")
	result := []Route{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		panic(err)
	}
	return result
}

// Finds the document in the collection for the given document name
func Find(name string) Route {
	result := Route{}
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("t2_dev").C("services_models_services")
	err = c.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		result = Route{}
		// panic(err)
	}
	return result
}

func Create(route Route) bool {
	checkRoute := Find(route.Name)
	if checkRoute.Name == "" {
		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		c := session.DB("t2_dev").C("services_models_services")
		err = c.Insert(&route)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

//func Connect() *mgo.Collection {
//	session, err := mgo.Dial("localhost")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//	return session.DB("t2_dev").C("services_models_services")
//}
