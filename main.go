package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func routes() *mux.Router {
	router := mux.NewRouter()
	//TODO: Add auth, goroutines, and error handling

	// User routes
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	// // Derby routes
	// router.HandleFunc("/api/derbies", getDerbies).Methods("GET")
	// router.HandleFunc("/api/derbies", createDerby).Methods("POST")

	// // Catch routes
	// router.HandleFunc("/api/catches", getCatches).Methods("GET")
	// router.HandleFunc("/api/catches", createCatch).Methods("POST")

	// // Location routes
	// router.HandleFunc("/api/locations", getLocations).Methods("GET")
	// router.HandleFunc("/api/locations", createLocation).Methods("POST")

	return router
}

var db *MongoDBService

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbService, err := NewMongoDBService(os.Getenv("MONGO_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("Error initializing MongoDB service: %v", err)
	}

	db = dbService

	router := routes()

	log.Println("Server started on port 3000")
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
