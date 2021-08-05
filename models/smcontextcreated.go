package models

import (
	db "TM5GProject/DB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SmContextCreated struct {
	SmContextRef              string               `json:"SmContextRef,omitempty" bson:"SmContextRef,omitempty"`
	SmContextCreatedData      SmContextCreatedData `json:"SmContextCreatedData,omitempty" bson:"SmContextCreatedData,omitempty"`
	BinaryDataN2SmInformation []byte               `json:"BinaryDataN2SmInformation,omitempty" bson:"BinaryDataN2SmInformation,omitempty"`
}

// get data from DB
func GetSMCtxCreatedData(smcontextref string) *SmContextCreated {

	collection := db.NewDBClient().Database("Repository").Collection("smcontextcreated")
	filter := bson.D{primitive.E{Key: "SmContextRef", Value: smcontextref}}

	var result SmContextCreated

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

// template/example data
var ExSmCtx = SmContextCreated{
	SmContextRef:              "dummysmcontextref",
	SmContextCreatedData:      SmContextCreatedDataList,
	BinaryDataN2SmInformation: nil,
}
