package main

import (
	"bytes"
	"encoding/json"
	"fuber/location"
	"fuber/models"
	"fuber/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	myRouter.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetCabs(t *testing.T) {
	initialize()
	req, _ := http.NewRequest("GET", "/getcabs", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var cabsList models.CabsListResponseModel

	json.Unmarshal(response.Body.Bytes(), &cabsList)

	// Checking to see if error exists
	if cabsList.Errors != "" {
		t.Errorf("Expected a list of cabs for a response. Got error")
	}
}

func TestBookCab(t *testing.T) {
	initialize()
	var userDetails = user.User{Pickup: location.Location{67, 100}}
	payloadBytes, _ := json.Marshal(userDetails)
	req, _ := http.NewRequest("POST", "/beginride", bytes.NewBuffer(payloadBytes))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var cab models.CabBookedResponseModel

	json.Unmarshal(response.Body.Bytes(), &cab)

	if cab.Errors != "" {
		t.Errorf("Expected a cab. Got %s", cab.Errors)
	}

}

func TestEndRide(t *testing.T) {
	initialize()
	var invalidUserDetails = user.User{Pickup: location.Location{67, 100}}
	payloadBytes, _ := json.Marshal(invalidUserDetails)
	req, _ := http.NewRequest("POST", "/endride", bytes.NewBuffer(payloadBytes))
	response := executeRequest(req)
	// This request should fail to status 400
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	// Now check for valid end ride by booking a ride and then ending it
	var userBookingDetails = user.User{Pickup: location.Location{67, 100}}
	payloadBytes, _ = json.Marshal(userBookingDetails)
	req, _ = http.NewRequest("POST", "/beginride", bytes.NewBuffer(payloadBytes))
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var validUserDetails = user.User{Pickup: location.Location{67, 100}, Drop: location.Location{70, 104}, CabId: 1, TravelTime: 6}
	payloadBytes, _ = json.Marshal(validUserDetails)
	req, _ = http.NewRequest("POST", "/endride", bytes.NewBuffer(payloadBytes))
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var bill models.CabRideEndResponseModel

	json.Unmarshal(response.Body.Bytes(), &bill)

	if bill.Errors != "" {
		t.Errorf("Expected a bill. Got %s", bill.Errors)
	}

}
