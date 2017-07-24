package cab

import (
	"testing"
)
// Testing the cab package
func TestGetCabs(t *testing.T) {
	var cabs []Cab
	InitializeCabs()
	cabs = GetAllCabs()
	if len(cabs) != 12 {
		t.Errorf("Total cabs obtained: %d, expected: 12 ", len(cabs))
	}
}
