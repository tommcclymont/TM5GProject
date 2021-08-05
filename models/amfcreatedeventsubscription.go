package models

import (
	db "TM5GProject/DB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AmfCreatedEventSubscription struct {
	Subscription      AmfEventSubscription `json:"Subscription,omitempty" bson:"Subscription,omitempty"`
	SubscriptionId    string               `json:"SubscriptionId,omitempty" bson:"SubscriptionId,omitempty"`
	SupportedFeatures string               `json:"SupportedFeatures,omitempty" bson:"SupportedFeatures,omitempty"`
}

// returns data
func GetEventExposeSubData(subid string) *AmfCreatedEventSubscription {

	collection := db.NewDBClient().Database("Repository").Collection("amfcreatedeventsubscription")
	filter := bson.D{primitive.E{Key: "SubscriptionId", Value: subid}}

	var result AmfCreatedEventSubscription

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

// template/example data
var ExEventSub = AmfCreatedEventSubscription{
	Subscription:      AmfEventSubscriptionData,
	SubscriptionId:    "dummysubid",
	SupportedFeatures: "D6e683bAfB2bbdCa",
}
