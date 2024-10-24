package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Structs for Users, Derbies, Catches, Locations

type User struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Derbies []Derby   `json:"derbies,omitempty"`
	Catches []Catch   `json:"catches,omitempty"`
}

type Derby struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Participants []User  `json:"participants,omitempty"`
	TopCatches   []Catch `json:"top_catches,omitempty"`
}

type Catch struct {
	ID       string   `json:"id"`
	UserID   string   `json:"user_id"`
	DerbyID  string   `json:"derby_id"`
	Species  string   `json:"species"`
	Weight   float64  `json:"weight"`
	Location Location `json:"location"`
}

type Location struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// In-memory storage (simulated database)
var users []User
var derbies []Derby
var catches []Catch
var locations []Location

// Handlers

// Get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// Get a user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

// Update a user by ID
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range users {
		if item.ID == params["id"] {
			users = append(users[:i], users[i+1:]...) // Delete the old user
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

// Delete a user by ID
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range users {
		if item.ID == params["id"] {
			users = append(users[:i], users[i+1:]...)
			json.NewEncoder(w).Encode(users)
			return
		}
	}
	http.NotFound(w, r)
}

// Get all derbies
func getDerbies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(derbies)
}

// Create a new derby
func createDerby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var derby Derby
	_ = json.NewDecoder(r.Body).Decode(&derby)
	derbies = append(derbies, derby)
	json.NewEncoder(w).Encode(derby)
}

// Get all catches
func getCatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catches)
}

// Record a new catch
func createCatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCatch Catch
	_ = json.NewDecoder(r.Body).Decode(&newCatch)
	catches = append(catches, newCatch)
	json.NewEncoder(w).Encode(newCatch)
}

// Get all locations
func getLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}

// Create a new location
func createLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var location Location
	_ = json.NewDecoder(r.Body).Decode(&location)
	locations = append(locations, location)
	json.NewEncoder(w).Encode(location)
}

// Main function to start the server
func main() {
	// Initialize router
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	// Derby routes
	router.HandleFunc("/api/derbies", getDerbies).Methods("GET")
	router.HandleFunc("/api/derbies", createDerby).Methods("POST")

	// Catch routes
	router.HandleFunc("/api/catches", getCatches).Methods("GET")
	router.HandleFunc("/api/catches", createCatch).Methods("POST")

	// Location routes
	router.HandleFunc("/api/locations", getLocations).Methods("GET")
	router.HandleFunc("/api/locations", createLocation).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", router))
}
