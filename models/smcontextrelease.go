package models

import (
	db "github.com/tommcclymont/TM5GProject/DB"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SmContextRelease struct {
	SmContextRef              string               `json:"SmContextRef,omitempty" bson:"SmContextRef,omitempty"`
	SmContextReleaseData      SmContextReleaseData `json:"SmContextReleaseData,omitempty" bson:"SmContextReleaseData,omitempty"`
	BinaryDataN2SmInformation []byte               `json:"BinaryDataN2SmInformation,omitempty" bson:"BinaryDataN2SmInformation,omitempty"`
}

// insert data to DB
func PostSmCtxRel(smcontextref string, d SmContextRelease) error {

	collection := db.NewDBClient().Database("Repository").Collection("smcontextrelease")
	filter := bson.D{primitive.E{Key: "SmContextRef", Value: smcontextref}}
	opts := options.Replace().SetUpsert(true)

	_, err := collection.ReplaceOne(context.TODO(), filter, d, opts)
	if err != nil {
		return ErrSmCtxRefNotFound
	}

	return nil
}
