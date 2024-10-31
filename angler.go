package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Angler struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Derbies []Derby `json:"derbies,omitempty"`
	Catches []Catch `json:"catches,omitempty"`
}

var anglers []Angler

func getAnglers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(anglers)
}

func createAngler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var anglers Angler
	_ = json.NewDecoder(r.Body).Decode(&anglers)

	newAngler, err := db.InsertDocument("DinkDerby", "anglers", anglers)
	if err != nil {
		LogHTTPError(err, r)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	LogHTTPInfo("Angler created", r)
	json.NewEncoder(w).Encode(newAngler)
}

func getAngler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid anglers ID", http.StatusBadRequest)
		LogHTTPError(err, r)
		return
	}

	anglersResult, err := db.FindDocument("DinkDerby", "anglers", bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		LogHTTPError(err, r)
		return
	}

	if anglersResult.Err() == mongo.ErrNoDocuments {
		NotFoundHandler(w, r)
		LogHTTPError(err, r)
		return
	}

	var anglers Angler
	if err := anglersResult.Decode(&anglers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		LogHTTPError(err, r)
		return
	}
	LogHTTPInfo("Angler retrieved", r)
	json.NewEncoder(w).Encode(anglers)

}

func updateAngler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid anglers ID", http.StatusBadRequest)
		LogHTTPError(err, r)
		return
	}
	var anglers Angler
	if err := json.NewDecoder(r.Body).Decode(&anglers); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		LogHTTPError(err, r)
		return
	}

	update := bson.M{
		"$set": anglers,
	}

	result, err := db.UpdateDocument("DinkDerby", "anglers", bson.M{"_id": objectID}, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		LogHTTPError(err, r)
		return
	}

	if result.MatchedCount == 0 {
		NotFoundHandler(w, r)
		LogHTTPError(err, r)
		return
	}
	LogHTTPInfo("Angler updated", r)
	json.NewEncoder(w).Encode(result)

}

func deleteAngler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid anglers ID", http.StatusBadRequest)
		LogHTTPError(err, r)
		return
	}

	result, err := db.DeleteDocument("DinkDerby", "anglers", bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		LogHTTPError(err, r)
		return
	}

	if result.DeletedCount == 0 {
		NotFoundHandler(w, r)
		LogHTTPError(err, r)
		return
	}

	json.NewEncoder(w).Encode(bson.M{"message": "Angler deleted"})
}
