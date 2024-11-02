package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func routes() *mux.Router {
	router := mux.NewRouter()
	//TODO: Add auth, goroutines, and error handling

	// Angler routes
	router.HandleFunc("/api/anglers/derby/{id}", getAnglersByDerby).Methods("GET")
	router.HandleFunc("/api/anglers", createAngler).Methods("POST")
	router.HandleFunc("/api/anglers/{id}", getAngler).Methods("GET")
	router.HandleFunc("/api/anglers/{id}", updateAngler).Methods("PUT")
	router.HandleFunc("/api/anglers/{id}", deleteAngler).Methods("DELETE")

	// Derby routes
	router.HandleFunc("/api/derbies/angler/{id}", getDerbiesByAngler).Methods("GET")
	router.HandleFunc("/api/derbies", createDerby).Methods("POST")
	router.HandleFunc("/api/derby/{id}", getDerby).Methods("GET")

	// Catch routes
	router.HandleFunc("/api/catches", getCatches).Methods("GET")
	router.HandleFunc("/api/catches", createCatch).Methods("POST")

	// Location routes
	router.HandleFunc("/api/locations", getLocations).Methods("GET")
	router.HandleFunc("/api/locations", createLocation).Methods("POST")

	return router
}

var db *MongoDBService

func main() {
	err := godotenv.Load(".env")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbService, err := NewMongoDBService(os.Getenv("MONGO_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("Error initializing MongoDB service: %v", err)
	}

	db = dbService

	router := routes()

	LogInfo("Server started on port 3000")
	http.ListenAndServe("localhost:3000", router)
}
