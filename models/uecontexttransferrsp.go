package models

import (
	db "github.com/tommcclymont/TM5GProject/DB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UeContextTransferRsp struct {
	UeContextId                 string                   `json:"UeContextId,omitempty" bson:"UeContextId,omitempty"`
	UeContextTransferRspData    UeContextTransferRspData `json:"UeContextTransferRspData,omitempty" bson:"UeContextTransferRspData,omitempty"`
	BinaryDataN2Information     []byte                   `json:"BinaryDataN2Information,omitempty" bson:"BinaryDataN2Information,omitempty"`
	BinaryDataN2InformationExt1 []byte                   `json:"BinaryDataN2InformationExt1,omitempty" bson:"BinaryDataN2InformationExt1,omitempty"`
}

// GetAMSData returns the dummy AMS data
func GetUeContextData(uecontextid string) *UeContextTransferRsp {

	collection := db.NewDBClient().Database("Repository").Collection("uecontexttransferrsp")
	filter := bson.D{primitive.E{Key: "UeContextId", Value: uecontextid}}

	var result UeContextTransferRsp

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}
