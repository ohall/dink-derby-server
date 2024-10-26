package main

import (
	"encoding/json"
	"net/http"
)

type Location struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

var locations []Location

func getLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}

func createLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var location Location
	_ = json.NewDecoder(r.Body).Decode(&location)
	locations = append(locations, location)
	json.NewEncoder(w).Encode(location)
}
