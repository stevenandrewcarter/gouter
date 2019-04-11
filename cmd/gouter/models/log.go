package models

import (
	"github.com/stevenandrewcarter/gouter/cmd/gouter"
	// "labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
)

// Structure that the routes will be represented by
type Log struct {
	CreatedAt                string
	From                     string
	To                       string
	ProcessingTimeSec        float64
	ResponseCode             int
}

// Load all of the logs from the database
// func LoadLogs() []Log {
// 	result := []Log{}
// 	fn := cmd.executeFunc(func(collection *mgo.Collection) {
// 		err := collection.Find(bson.M{}).All(&result)
// 		if err != nil { panic(err) }
// 	})
// 	cmd.execute(fn, cmd.LOG_COLLECTION)
// 	return result
// }

// // Create a route from the provided route structure
// // Returns true if the route was created successfully, or false if it already exists
// func CreateLog(log Log) bool {
// 	created := false
// 	fn := cmd.executeFunc(func(collection *mgo.Collection) {
// 		err := collection.Insert(&log)
// 		if err != nil { panic(err) }
// 		created = true
// 	})
// 	cmd.execute(fn, cmd.LOG_COLLECTION)
// 	return created
// }
