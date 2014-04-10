// Package models includes the models used for the gouter configuration
package models

import (
	"labix.org/v2/mgo"
	"github.com/stevenandrewcarter/gouter"
)

const ROUTE_COLLECTION = "routes"
const ERROR_COLLECTION = "routes"

// Execute function signature, used as a helper for functions to database connections
type executeFunc func(collection *mgo.Collection)

// Executes the function provided against the database. Will connect to the database as per the
// application configuration and execute the function providing the desired collection
func execute(fn executeFunc, col string) {
	session, err := mgo.Dial(gouter.Configuration().Database.Host)
	if err != nil { panic(err) }
	defer session.Close()
	collection := session.DB(gouter.Configuration().Database.Database).C(col)
	// Call the provided function with the loaded collection
	fn(collection)
}
