package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
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

	log.Println("Server started on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
