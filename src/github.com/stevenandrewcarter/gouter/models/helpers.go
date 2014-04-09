package models

import (
	"labix.org/v2/mgo"
	"github.com/stevenandrewcarter/gouter"
)

type executeFunc func(collection *mgo.Collection)

func execute(fn executeFunc) {
	session, err := mgo.Dial(gouter.Configuration().Database.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection := session.DB(gouter.Configuration().Database.Database).C(routeCollection)
	// Call the provided function
	fn(collection)
}
