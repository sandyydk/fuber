package handlers

import (
	"fuber/location"
	"fuber/user"
	"testing"
)

func TestDistance(t *testing.T) {

	var testTables = []struct {
		x location.Location
		y location.Location
		z float64
	}{
		{location.Location{8, 100}, location.Location{11, 104}, 5},
		{location.Location{67, 100}, location.Location{70, 104}, 5},
		{location.Location{70, 96}, location.Location{77, 120}, 25},
	}

	for index, table := range testTables {
		totalDistance := distance(&table.x, &table.y)
		if totalDistance != table.z {
			t.Errorf("Distance returned is incorrect for index: %d , got: %f, expected:%f ", index, totalDistance, table.z)
		}
	}
}

func TestBill(t *testing.T) {
	var testTables = []struct {
		x user.User
		z int64
	}{
		{user.User{location.Location{67, 100}, location.Location{70, 104}, "pink", 1, 6}, 21},
		{user.User{location.Location{67, 100}, location.Location{70, 104}, "green", 1, 6}, 16},
	}

	for index, table := range testTables {
		bill, _ := calculateBill(table.x)
		if bill != table.z {
			t.Errorf("Bill calculation is incorrect for index: %d , got: %f, expected:%f ", index, bill, table.z)
		}
	}
}
