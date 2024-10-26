package main

import (
	"encoding/json"
	"net/http"
)

type Catch struct {
	ID       string   `json:"id"`
	UserID   string   `json:"user_id"`
	DerbyID  string   `json:"derby_id"`
	Species  string   `json:"species"`
	Weight   float64  `json:"weight"`
	Location Location `json:"location"`
}

var catches []Catch

func getCatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catches)
}

func createCatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCatch Catch
	_ = json.NewDecoder(r.Body).Decode(&newCatch)
	catches = append(catches, newCatch)
	json.NewEncoder(w).Encode(newCatch)
}
