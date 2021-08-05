package models

import (
	db "github.com/tommcclymont/TM5GProject/DB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// structure for SMF selection subscription data as defined in specification 29.503 release 15
type SmfSelectionSubscriptionData struct {
	Supi                  string             `json:"Supi,omitempty" bson:"Supi,omitempty"`
	SupportedFeatures     string             `json:"SupportedFeatures,omitempty" bson:"SupportedFeatures,omitempty"`
	SubscribedSnssaiInfos map[string]DnnInfo `json:"SubscribedSnssaiInfos,omitempty" bson:"SubscribedSnssaiInfos,omitempty"`
	SharedSnssaiInfosId   string             `json:"SharedSnssaiInfosId,omitempty" bson:"SharedSnssaiInfosId,omitempty"`
}

// get data from DB
func GetSMFSubData(supi string) *SmfSelectionSubscriptionData {

	collection := db.NewDBClient().Database("Repository").Collection("smfselectionsubscriptiondata")
	filter := bson.D{primitive.E{Key: "Supi", Value: supi}}

	var result SmfSelectionSubscriptionData

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

// template/example data
var ExSelSubData = SmfSelectionSubscriptionData{
	Supi:                  "892881507",
	SupportedFeatures:     "D6e683bAfB2bbdCa",
	SubscribedSnssaiInfos: SubscribedSnssaiInfosMap,
	SharedSnssaiInfosId:   "85562-test",
}
