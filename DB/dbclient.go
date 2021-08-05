package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDBClient() *mongo.Client {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://Tom:h7fjn89o@tm5gproject.n1vef.mongodb.net")

	// Connect to MongoDB
	DBClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return DBClient
}
