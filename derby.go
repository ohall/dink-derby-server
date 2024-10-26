package main

import (
	"encoding/json"
	"net/http"
)

type Derby struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	Participants []User  `json:"participants,omitempty"`
	TopCatches   []Catch `json:"top_catches,omitempty"`
}

var derbies []Derby

func getDerbies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(derbies)
}

func createDerby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var derby Derby
	_ = json.NewDecoder(r.Body).Decode(&derby)
	derbies = append(derbies, derby)
	json.NewEncoder(w).Encode(derby)
}
