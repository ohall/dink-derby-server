package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func routes() *mux.Router {
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

	return router
}

func initializeMongoDB() (*mongo.Client, error) {
	// Set client options from env variable for connection string
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}

func main() {
	router := routes()

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbService, err := NewMongoDBService(os.Getenv("MONGO_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("Error initializing MongoDB service: %v", err)
	}

	log.Println("Server started on port 3000")
	log.Println("DB Service: ", dbService)
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
