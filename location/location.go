package location

import (
	"fmt"
)

// Define location struct to be used in cab package. Use capital letters to be able to export to other packages
type Location struct {
	Latitude  int64 `json:"Latitude"`
	Longitude int64 `json:"Longitude"`
}

// Go doesn't have constructors the way Java or C++ has that the compiler would run. Instead the values are set
// to their default values like a string to empty string.

func (l *Location) displayLocation() {
	fmt.Sprintf("Location is  Latitude: %d Longitude: %d", l.Latitude, l.Longitude)
}

func (l *Location) GetLatitude() int64 {
	return l.Latitude
}

func (l *Location) GetLongitude() int64 {
	return l.Longitude
}
