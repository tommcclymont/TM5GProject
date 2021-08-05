package models

import (
	db "TM5GProject/DB"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UeContextTransferReq struct {
	UeContextId              string                   `json:"UeContextId,omitempty" bson:"UeContextId,omitempty"`
	UeContextTransferReqData UeContextTransferReqData `json:"UeContextTransferReqData,omitempty" bson:"UeContextTransferReqData,omitempty"`
	BinaryDataN1Message      []byte                   `json:"BinaryDataN1Message,omitempty" bson:"BinaryDataN1Message,omitempty"`
}

var ErrUeCtxidNotFound = fmt.Errorf("UE Context ID not found")

// insert data to DB
func PostUeContextReq(uecontextid string, d UeContextTransferReq) error {

	collection := db.NewDBClient().Database("Repository").Collection("uecontexttransferreq")
	filter := bson.D{primitive.E{Key: "UeContextId", Value: uecontextid}}
	opts := options.Replace().SetUpsert(true)

	_, err := collection.ReplaceOne(context.TODO(), filter, d, opts)
	if err != nil {
		return ErrUeCtxidNotFound
	}

	return nil
}
