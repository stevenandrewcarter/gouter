package models

// import (
// 	"github.com/stevenandrewcarter/gouter/cmd/gouter"
// 	// "labix.org/v2/mgo"
// 	// "labix.org/v2/mgo/bson"
// 	"log"
// )

// Structure that the routes will be represented by
type Route struct {
	Description string `yaml:"description"`
	Name        string `yaml:"name"`
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

// // Load all of the routes from the database
func LoadRoutes() []Route {
	result := []Route{}
	//fn := cmd.executeFunc(func(collection *mgo.Collection) {
	//	err := collection.Find(bson.M{}).All(&result)
	//	if err != nil { panic(err) }
	//})
	//cmd.execute(fn, cmd.ROUTE_COLLECTION)
	return result
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

// Create a route from the provided route structure
// Returns true if the route was created successfully, or false if it already exists
func CreateRoute(route Route) bool {
	created := false
	//checkRoute := FindRoute(route.Name)
	//if checkRoute.Name == "" {
	//	fn := cmd.executeFunc(func(collection *mgo.Collection) {
	//		err := collection.Insert(&route)
	//		if err != nil { panic(err) }
	//		log.Printf("Route created '%v'", route.Name)
	//		created = true
	//	})
	//	cmd.execute(fn, cmd.ROUTE_COLLECTION)
	//}
	return created
}

// Delete the route from the database with the given route name
// Returns a true if successful or false if the route could not be found
func DeleteRoute(name string) bool {
	deleted := false
	//fn := cmd.executeFunc(func(collection *mgo.Collection) {
	//	err := collection.Remove(bson.M{"name": name})
	//	if err != nil { panic(err) }
	//	log.Printf("Route deleted '%v'", name)
	//	deleted = true
	//})
	//cmd.execute(fn, cmd.ROUTE_COLLECTION)
	return deleted
}
