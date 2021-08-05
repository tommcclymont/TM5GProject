package models

import (
	db "TM5GProject/DB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EirResponseData struct {
	Pei    string `json:"Pei,omitempty" bson:"Pei,omitempty"`
	Status string `json:"Status,omitempty" bson:"Status,omitempty"`
}

// get data from DB
func GetEirResponseData(pei string) *EirResponseData {

	collection := db.NewDBClient().Database("Repository").Collection("eirresponsedata")
	filter := bson.D{primitive.E{Key: "Pei", Value: pei}}

	var result EirResponseData

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

// template/example data
var ExEirData = EirResponseData{
	Pei:    "imei-857390213674531",
	Status: "WHITELISTED",
}
