package user

import (
	"fuber/location"
)

// User structure to book a cab. THis would normally be in DB where cab allocated is mapped to user's transaction/usage. Hence the cabId here which would be used to verify before ending a ride
type User struct {
	Pickup     location.Location `json:"Pickup"`
	Drop       location.Location `json:"Drop"`
	Color      string            `json:"Color"`
	CabId      int32             `json:"CabId"`
	TravelTime int64             `json:"TravelTime"`
}

func (u *User) GetPickup() *location.Location {
	return &u.Pickup
}

func (u *User) GetDrop() *location.Location {
	return &u.Drop
}

func (u *User) GetColorPreference() string {
	return u.Color
}
