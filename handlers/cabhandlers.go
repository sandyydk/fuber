package handlers

import (
	"encoding/json"
	"fmt"
	"fuber/cab"
	"fuber/location"
	"fuber/models"
	"fuber/services"
	"fuber/user"
	"math"
	"net/http"
	"strings"
)

// GetCabs - List all cabs irrespective of whether available or not. This is handled while assigning a cab to a user
func GetCabs(w http.ResponseWriter, r *http.Request) {

	// TO allow local testing , CORS when the UI connects from localhost:4200 to localhost:3200 of the server. Not needed if only testing using POSTMAN
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	if r.Method == "OPTIONS" {
		return
	}
	var cabsList models.CabsListResponseModel
	cabsList.Cabs = cab.GetAllCabs()
	cabs, err := json.Marshal(cabsList)
	if err != nil {
		fmt.Println(err)
		services.ErrorWithJSON(w, "Sorry something went wrong", http.StatusInternalServerError)
	}
	services.ResponseWithJSON(w, cabs, http.StatusOK)
}

func distance(userLocation *location.Location, cabLocation *location.Location) float64 {
	var distance = math.Sqrt(float64(((userLocation.GetLatitude() - cabLocation.GetLatitude()) * (userLocation.GetLatitude() - cabLocation.GetLatitude())) + ((userLocation.GetLongitude() - cabLocation.GetLongitude()) * (userLocation.GetLongitude() - cabLocation.GetLongitude()))))
	return distance
}

// Function to calculate bill amount. Distance between pickup and drop is taken to be in kms. Both distance and bill amount are given back in Whole numbers
func calculateBill(user user.User) (int64, int64) {
	totalDistance := distance(user.GetDrop(), user.GetPickup())
	var billAmount int64
	billAmount = int64((totalDistance * 2)) + user.TravelTime
	if strings.EqualFold(user.GetColorPreference(), "pink") {
		billAmount = billAmount + 5
	}
	return billAmount, int64(totalDistance)
}

// BookCab Handler function to book a cab. Extra detail like user drop location is included in the structure which is not used as of now but can be used if necessary in the future
func BeginRide(w http.ResponseWriter, r *http.Request) {

	// TO allow local testing , CORS when the UI connects from localhost:4200 to localhost:3200 of the server. Not needed if only testing using POSTMAN
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	if r.Method == "OPTIONS" {
		return
	}

	var cabs = cab.GetAllCabs()
	var bookingResponse models.CabBookedResponseModel
	var nearestCabDistance float64
	var userDetails user.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userDetails)
	if err != nil {
		fmt.Println(err)
		services.ErrorWithJSON(w, "Failed to Book a cab", http.StatusInternalServerError)
		return
	}
	// Need to filter if the user has specified a color preference. Golang doesn't provide Generics like filter etc. Need to define manually. Iterating is the way to go
	for _, c := range cabs {
		if userDetails.GetColorPreference() != "" {

			if c.Available() && strings.EqualFold(c.GetColor(), userDetails.GetColorPreference()) {
				dist := distance(userDetails.GetPickup(), c.GetLocation())
				if dist < nearestCabDistance || (bookingResponse.Cab == cab.Cab{}) {
					nearestCabDistance = dist
					bookingResponse.Cab = c
				}
			}
		} else {
			if c.Available() {
				dist := distance(userDetails.GetPickup(), c.GetLocation())
				if dist < nearestCabDistance || (bookingResponse.Cab == cab.Cab{}) {
					nearestCabDistance = dist
					bookingResponse.Cab = c
				}
			}
		}
	}

	if (bookingResponse.Cab != cab.Cab{}) {
		cab.SetAvailability(bookingResponse.Cab.ID, false)
		cab.UpdateCabLocation(bookingResponse.Cab.ID, &userDetails.Pickup)

		response, err := json.Marshal(bookingResponse)
		if err != nil {
			fmt.Println(err)
			services.ErrorWithJSON(w, "Failed to book a cab", http.StatusInternalServerError)
		}
		services.ResponseWithJSON(w, response, http.StatusOK)
	} else {
		services.ErrorWithJSON(w, "No Cab Found", http.StatusBadRequest)
	}
}

func EndRide(w http.ResponseWriter, r *http.Request) {

	// TO allow local testing , CORS when the UI connects from localhost:4200 to localhost:3200 of the server. Not needed if only testing using POSTMAN
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	if r.Method == "OPTIONS" {
		return
	}

	var cabs = cab.GetAllCabs()
	var bookingResponse models.CabRideEndResponseModel
	var userDetails user.User
	var validBooking = false
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userDetails)
	if err != nil {
		fmt.Println(err)
		services.ErrorWithJSON(w, "Failed to End Ride", http.StatusInternalServerError)
	}
	// Check to ensure all required fields are present
	if (userDetails.Drop != location.Location{} && userDetails.Pickup != location.Location{} && userDetails.TravelTime != 0 && userDetails.CabId != 0) {
		for _, c := range cabs {
			if userDetails.GetPickup().GetLatitude() == c.GetLocation().GetLatitude() && userDetails.GetPickup().GetLongitude() == c.GetLocation().GetLongitude() && userDetails.CabId == c.ID && !c.Available() {
				cab.UpdateCabLocation(c.ID, userDetails.GetDrop())
				cab.SetAvailability(c.ID, true)
				validBooking = true
				break
			}
		}
	}
	if validBooking {
		// Now once cab allocation to user is verified calculate the bill
		bookingResponse.Bill, bookingResponse.Distance = calculateBill(userDetails)
		bookingResponse.IsEnded = true

		response, err := json.Marshal(bookingResponse)
		if err != nil {
			fmt.Println(err)
			services.ErrorWithJSON(w, "Failed to book a cab", http.StatusInternalServerError)
		}
		services.ResponseWithJSON(w, response, http.StatusOK)
	} else {
		// Booking doesn't exist for this user
		services.ErrorWithJSON(w, "Invalid Booking", http.StatusBadRequest)
	}

}
