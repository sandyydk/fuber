package main

import (
	"fmt"
	"fuber/cab"
	"fuber/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initialize() {
	// Initializing a fleet of cabs. Everything related to cabs is in cab package
	cab.InitializeCabs()
	myRouter.HandleFunc("/getcabs", handlers.GetCabs).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/beginride", handlers.BeginRide).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/endride", handlers.EndRide).Methods("POST", "OPTIONS")
}

var myRouter = mux.NewRouter().StrictSlash(true)

func handleRequests() {
	fmt.Println("Serving at 3002")
	log.Fatal(http.ListenAndServe(":3002", myRouter))
}

func main() {
	initialize()
	handleRequests()
}
