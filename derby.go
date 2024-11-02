package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Derby struct {
	Name       string               `json:"name"`
	StartDate  string               `json:"start_date"`
	EndDate    string               `json:"end_date"`
	Image      string               `json:"image"`
	Active     bool                 `json:"active"`
	Location   primitive.ObjectID   `json:"location"`
	Winner     primitive.ObjectID   `json:"winner"`
	Anglers    []primitive.ObjectID `json:"anglers,omitempty"`
	Catches    []primitive.ObjectID `json:"catches,omitempty"`
}

func createDerby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var derby Derby
	_ = json.NewDecoder(r.Body).Decode(&derby)

	newDerby, err := db.InsertDocument("DinkDerby", "derbies", derby)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newDerby)
}

func getDerby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid derby ID", http.StatusBadRequest)
		return
	}

	derbyResult, err := db.FindDocument("DinkDerby", "derbies", bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if derbyResult.Err() == mongo.ErrNoDocuments {
		NotFoundHandler(w, r)
		return
	}

	var derby Derby
	if err := derbyResult.Decode(&derby); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(derby)
}

func getDerbiesByAngler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	anglerID := params["id"]
	log.Println("anglerID", anglerID)

	objectID, err := primitive.ObjectIDFromHex(anglerID)
	if err != nil {
		http.Error(w, "Invalid angler ID", http.StatusBadRequest)
		return
	}

	filter := bson.D{{Key: "anglers", Value: objectID}}

	derbiesCursor, err := db.FindDocuments("DinkDerby", "derbies", filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var derbies []Derby
	if err := derbiesCursor.All(context.TODO(), &derbies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("derbies", derbies)
	if len(derbies) == 0 {
		NotFoundHandler(w, r)
		return
	}
	json.NewEncoder(w).Encode(derbies)
}
