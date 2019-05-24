package lib

// Structure that the routes will be represented by
type Route struct {
	Description string `yaml:"description"`
	Name        string `yaml:"name"`
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

var routes []Route

// Loads the Routes as per the configuration
func Load() {

}
