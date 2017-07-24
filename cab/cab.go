package cab

import (
	"fuber/location"
)

// Cab structure
type Cab struct {
	Location    *location.Location `json:"Location"`
	Color       string             `json:"Color"`
	IsAvailable bool               `json:"IsAvailable"`
	ID          int32              `json:"ID"`
}

var cabsList []Cab

func InitializeCabs() {
	// Setting up my cabs here. Assuming a dozen cabs. This can be taken as an input too. For the sake
	// of convenience am setting it up here
	cabsList = []Cab{
		{
			ID: 1, Color: "pink", IsAvailable: true, Location: &location.Location{12, 100},
		},
		{
			ID: 2, Color: "green", IsAvailable: true, Location: &location.Location{8, 100},
		},
		{
			ID: 3, Color: "blue", IsAvailable: true, Location: &location.Location{15, 100},
		},
		{
			ID: 4, Color: "red", IsAvailable: true, Location: &location.Location{5, 100},
		},
		{
			ID: 5, Color: "red", IsAvailable: true, Location: &location.Location{10, 100},
		},
		{
			ID: 6, Color: "grey", IsAvailable: true, Location: &location.Location{10, 100},
		},
		{
			ID: 7, Color: "black", IsAvailable: true, Location: &location.Location{9, 120},
		},
		{
			ID: 8, Color: "pink", IsAvailable: true, Location: &location.Location{6, 150},
		},
		{
			ID: 9, Color: "pink", IsAvailable: true, Location: &location.Location{18, 110},
		},
		{
			ID: 10, Color: "yellow", IsAvailable: true, Location: &location.Location{10, 80},
		},
		{
			ID: 11, Color: "blue", IsAvailable: true, Location: &location.Location{10, 105},
		},
		{
			ID: 12, Color: "green", IsAvailable: true, Location: &location.Location{90, 140},
		},
	}
}

func GetAllCabs() []Cab {
	return cabsList
}

func (c *Cab) getID() int32 {
	return c.ID
}

func (c *Cab) GetLocation() *location.Location {
	return c.Location
}

func SetAvailability(cabId int32, status bool) {
	for index, val := range cabsList {
		if val.ID == cabId {
			cabsList[index].IsAvailable = status
		}
	}
}

func UpdateCabLocation(cabId int32, loc *location.Location) {
	// Range copies the slice, hence need to use the index to update the value
	for index, val := range cabsList {
		if val.ID == cabId {
			cabsList[index].Location = loc
		}
	}
}

func (c *Cab) SetAvail(status bool) {
	c.IsAvailable = status
}

func (c *Cab) Available() bool {
	return c.IsAvailable
}

func (c *Cab) GetColor() string {
	return c.Color
}
