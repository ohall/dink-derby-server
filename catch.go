package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Catch struct {
	AnglerID primitive.ObjectID `json:"angler_id"`
	DerbyID  string             `json:"derby_id"`
	Species  string             `json:"species"`
	Weight   float64            `json:"weight"`
	Location Location           `json:"location"`
}

func getCatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid catch ID", http.StatusBadRequest)
		return
	}

	catchResult, err := db.FindDocument("DinkDerby", "catches", bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if catchResult.Err() == mongo.ErrNoDocuments {
		NotFoundHandler(w, r)
		return
	}

	var catch Catch
	if err := catchResult.Decode(&catch); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(catch)
}

func createCatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var catch Catch
	if err := json.NewDecoder(r.Body).Decode(&catch); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newCatch, err := db.InsertDocument("DinkDerby", "catches", catch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newCatch)

}
