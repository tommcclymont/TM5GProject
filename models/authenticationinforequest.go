package models

import (
	db "github.com/tommcclymont/TM5GProject/DB"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthenticationInfoRequest struct {
	Supi                  string                `json:"Supi,omitempty" bson:"Supi,omitempty"`
	ServingNetworkName    string                `json:"ServingNetworkName,omitempty" bson:"ServingNetworkName,omitempty"`
	ResynchronizationInfo ResynchronizationInfo `json:"ResynchronizationInfo,omitempty" bson:"ResynchronizationInfo,omitempty"`
	SupportedFeatures     string                `json:"SupportedFeatures,omitempty" bson:"SupportedFeatures,omitempty"`
	AusfInstanceId        string                `json:"AusfInstanceId,omitempty" bson:"AusfInstanceId,omitempty"`
}

// insert data to DB
func PostGenAuth(supi string, d AuthenticationInfoRequest) error {

	collection := db.NewDBClient().Database("Repository").Collection("authenticationinforequest")
	filter := bson.D{primitive.E{Key: "Supi", Value: supi}}
	opts := options.Replace().SetUpsert(true)

	_, err := collection.ReplaceOne(context.TODO(), filter, d, opts)
	if err != nil {
		return ErrSupiNotFound
	}

	return nil
}
