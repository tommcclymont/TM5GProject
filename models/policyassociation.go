package models

import (
	db "github.com/tommcclymont/TM5GProject/DB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PolicyAssociation struct {
	PolAssoId   string                   `json:"PolAssoId,omitempty" bson:"PolAssoId,omitempty"`
	Request     PolicyAssociationRequest `json:"Request,omitempty" bson:"Request,omitempty"`
	Triggers    []string                 `json:"Triggers,omitempty" bson:"Triggers,omitempty"`
	ServAreaRes ServiceAreaRestriction   `json:"ServAreaRes,omitempty" bson:"ServAreaRes,omitempty"`
	Rfsp        uint                     `json:"Rfsp,omitempty" bson:"Rfsp,omitempty"`
	SuppFeat    string                   `json:"SuppFeat,omitempty" bson:"SuppFeat,omitempty"`
}

// get data from DB
func GetPolAssoData(polassoid string) *PolicyAssociation {

	collection := db.NewDBClient().Database("Repository").Collection("policyassociation")
	filter := bson.D{primitive.E{Key: "PolAssoId", Value: polassoid}}

	var result PolicyAssociation

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

// delete data in DB
func DeletePolAsso(polassoid string) error {

	collection := db.NewDBClient().Database("Repository").Collection("policyassociation")
	filter := bson.D{primitive.E{Key: "PolAssoId", Value: polassoid}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return ErrPolAssoIdNotFound
	}

	return nil
}

// template/example data
var ExPolAssoData = PolicyAssociation{
	PolAssoId: "dummypolassoid",
	Request: PolicyAssociationRequest{
		NotificationUri:  "https://127.0.0.78:9090/",
		AltNotifIpv4Addr: "1.1.1.1",
		AltNotifIpv6Addr: "1111:a1a1:1111:a1a1:1111:a1a1:1111:a1a1",
		Supi:             "892881507",
		Gpsi:             "msisdn-61738596071",
		AccessType:       "3GPP_ACCESS",
		Pei:              "imei-857390213674531",
		UserLoc:          UserlocationData,
		TimeZone:         "00:00+0",
		ServingPlmn:      PlmnIdData,
		RatType:          RatTypeData,
		GroupIds:         []string{"12345678-111-22-3", "12345678-444-55-6"},
		ServAreaRes:      ServiceAreaRestrictionData,
		Rfsp:             1,
		Guami:            GuamiData,
		ServiceName:      "testamf",
		SuppFeat:         "D6e683bAfB2bbdCa",
	},
	Triggers:    []string{"LOC_CH"},
	ServAreaRes: ServiceAreaRestrictionData,
	Rfsp:        1,
	//Pras: ,
	SuppFeat: "D6e683bAfB2bbdCa",
}
