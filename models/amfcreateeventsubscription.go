package models

import (
	db "TM5GProject/DB"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AmfCreateEventSubscription struct {
	Subscription      AmfEventSubscription `json:"Subscription,omitempty" bson:"Subscription,omitempty"`
	SubscriptionId    string               `json:"SubscriptionId,omitempty" bson:"SubscriptionId,omitempty"`
	SupportedFeatures string               `json:"SupportedFeatures,omitempty" bson:"SupportedFeatures,omitempty"`
	OldGuami          Guami                `json:"OldGuami,omitempty" bson:"OldGuami,omitempty"`
}

var ErrSubIdNotFound = fmt.Errorf("Subscription ID not found")

// inserts data to DB
func PostEventExposeSub(subid string, d AmfCreateEventSubscription) error {

	collection := db.NewDBClient().Database("Repository").Collection("amfcreateeventsubscription")
	filter := bson.D{primitive.E{Key: "SubscriptionId", Value: subid}}
	opts := options.Replace().SetUpsert(true)

	_, err := collection.ReplaceOne(context.TODO(), filter, d, opts)
	if err != nil {
		return ErrSubIdNotFound
	}

	return nil
}
