package models

import (
	"fuber/cab"
)

type CabBookedResponseModel struct {
	Cab    cab.Cab `json:"Cab"`
	Errors string  `json:"Errors"`
}

type CabsListResponseModel struct {
	Cabs   []cab.Cab `json:"Cabs"`
	Errors string    `json:"Errors"`
}

type CabRideEndResponseModel struct {
	Bill     int64  `json:"Bill"`
	Distance int64  `json:"Distance"`
	IsEnded  bool   `json:"IsEnded"`
	Errors   string `json:"Errors"`
}
