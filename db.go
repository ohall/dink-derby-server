package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBService struct {
	client *mongo.Client
}

func NewMongoDBService(uri string) (*MongoDBService, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return &MongoDBService{client: client}, nil
}

func (s *MongoDBService) InsertDocument(database, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	coll := s.client.Database(database).Collection(collection)
	insertResult, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

func (s *MongoDBService) FindDocument(database, collection string, filter interface{}) (*mongo.SingleResult, error) {
	coll := s.client.Database(database).Collection(collection)
	result := coll.FindOne(context.TODO(), filter)
	return result, nil
}

func (s *MongoDBService) UpdateDocument(database, collection string, filter, update interface{}) (*mongo.UpdateResult, error) {
	coll := s.client.Database(database).Collection(collection)
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *MongoDBService) DeleteDocument(database, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	coll := s.client.Database(database).Collection(collection)
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *MongoDBService) Close() error {
	return s.client.Disconnect(context.TODO())
}
