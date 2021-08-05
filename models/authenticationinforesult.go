package models

import (
	db "github.com/tommcclymont/TM5GProject/DB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthenticationInfoResult struct {
	AuthType             string               `json:"AuthType,omitempty" bson:"AuthType,omitempty"`
	AuthenticationVector AuthenticationVector `json:"AuthenticationVector,omitempty" bson:"AuthenticationVector,omitempty"`
	Supi                 string               `json:"Supi,omitempty" bson:"Supi,omitempty"`
	SupportedFeatures    string               `json:"SupportedFeatures,omitempty" bson:"SupportedFeatures,omitempty"`
}

// get data from DB
func GetAuthInfoRes(supi string) *AuthenticationInfoResult {

	collection := db.NewDBClient().Database("Repository").Collection("authenticationinforesult")
	filter := bson.D{primitive.E{Key: "Supi", Value: supi}}

	var result AuthenticationInfoResult

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

// template/example data
var ExAuthInfo = AuthenticationInfoResult{
	AuthType:             "5G_AKA",
	AuthenticationVector: AuthenticationVector{Av5GHeAkaData},
	SupportedFeatures:    "D6e683bAfB2bbdCa",
}
