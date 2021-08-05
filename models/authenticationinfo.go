package models

import (
	db "TM5GProject/DB"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthenticationInfo struct {
	AuthCtxId             string                `json:"AuthCtxId,omitempty" bson:"AuthCtxId,omitempty"`
	SupiOrSuci            string                `json:"SupiOrSuci,omitempty" bson:"SupiOrSuci,omitempty"`
	ServingNetworkName    string                `json:"ServingNetworkName,omitempty" bson:"ServingNetworkName,omitempty"`
	ResynchronizationInfo ResynchronizationInfo `json:"ResynchronizationInfo,omitempty" bson:"ResynchronizationInfo,omitempty"`
	Pei                   string                `json:"Pei,omitempty" bson:"Pei,omitempty"`
	TraceData             TraceData             `json:"TraceData,omitempty" bson:"TraceData,omitempty"`
}

var ErrAuthCtxIdNotFound = fmt.Errorf("Authentication Context ID not found")

// insert data to DB
func PostUeAuth(authctxid string, d AuthenticationInfo) error {

	collection := db.NewDBClient().Database("Repository").Collection("authenticationinfo")
	filter := bson.D{primitive.E{Key: "AuthCtxId", Value: authctxid}}
	opts := options.Replace().SetUpsert(true)

	_, err := collection.ReplaceOne(context.TODO(), filter, d, opts)
	if err != nil {
		return ErrAuthCtxIdNotFound
	}

	return nil
}
