package models

import (
	db "github.com/tommcclymont/TM5GProject/DB"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SmContextCreate struct {
	SmContextRef                 string              `json:"SmContextRef,omitempty" bson:"SmContextRef,omitempty"`
	SmContextCreateData          SmContextCreateData `json:"SmContextCreateData,omitempty" bson:"SmContextCreateData,omitempty"`
	BinaryDataN1SmMessage        []byte              `json:"BinaryDataN1SmMessage,omitempty" bson:"BinaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation    []byte              `json:"BinaryDataN2SmInformation,omitempty" bson:"BinaryDataN2SmInformation,omitempty"`
	BinaryDataN2SmInformationEx1 []byte              `json:"BinaryDataN2SmInformationEx1,omitempty" bson:"BinaryDataN2SmInformationEx1,omitempty"`
}

var ErrSmCtxRefNotFound = fmt.Errorf("SM Context Ref not found")

// insert data to DB
func PostSmCtx(smcontextref string, d SmContextCreate) error {

	collection := db.NewDBClient().Database("Repository").Collection("smcontextcreate")
	filter := bson.D{primitive.E{Key: "SmContextRef", Value: smcontextref}}
	opts := options.Replace().SetUpsert(true)

	_, err := collection.ReplaceOne(context.TODO(), filter, d, opts)
	if err != nil {
		return ErrSmCtxRefNotFound
	}

	return nil
}

// delete data in DB
func ReleaseSmCtx(smcontextref string) error {

	collection := db.NewDBClient().Database("Repository").Collection("smcontextcreate")
	filter := bson.D{primitive.E{Key: "SmContextRef", Value: smcontextref}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return ErrSmCtxRefNotFound
	}

	return nil
}

// template/example data
var ExSmCtxCreateData = SmContextCreate{
	SmContextRef:                 "dummysmcontextref",
	SmContextCreateData:          SmContextCreateDataList,
	BinaryDataN1SmMessage:        nil,
	BinaryDataN2SmInformation:    nil,
	BinaryDataN2SmInformationEx1: nil,
}
