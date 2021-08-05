package models

import (
	db "github.com/tommcclymont/TM5GProject/DB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// structure for AMS data as defined in specification 29.503 release 15
type AccessAndMobilitySubscriptionData struct {
	Supi                        string                 `json:"Supi,omitempty" bson:"Supi,omitempty"`
	SupportedFeatures           string                 `json:"SupportedFeatures,omitempty" bson:"SupportedFeatures,omitempty"`
	Gpsis                       []string               `json:"Gpsis,omitempty" bson:"Gpsis,omitempty"`
	InternalGroupIds            []string               `json:"InternalGroupIds,omitempty" bson:"InternalGroupIds,omitempty"`
	SubscribedUeAmbr            Ambr                   `json:"SubscribedUeAmbr,omitempty" bson:"SubscribedUeAmbr,omitempty"`
	Nssai                       Nssai                  `json:"Nssai,omitempty" bson:"Nssai,omitempty"`
	RatRestrictions             []RatType              `json:"RatRestrictions,omitempty" bson:"RatRestrictions,omitempty"`
	ForbiddenAreas              []Area                 `json:"ForbiddenAreas,omitempty" bson:"ForbiddenAreas,omitempty"`
	ServiceAreaRestriction      ServiceAreaRestriction `json:"ServiceAreaRestriction,omitempty" bson:"ServiceAreaRestriction,omitempty"`
	CoreNetworkTypeRestrictions []CoreNetworkType      `json:"CoreNetworkTypeRestrictions,omitempty" bson:"CoreNetworkTypeRestrictions,omitempty"`
	RfspIndex                   uint                   `json:"RfspIndex,omitempty" bson:"RfspIndex,omitempty"`
	SubsRegTimer                uint                   `json:"SubsRegTimer,omitempty" bson:"SubsRegTimer,omitempty"`
	UeUsageType                 int                    `json:"UeUsageType,omitempty" bson:"UeUsageType,omitempty"`
	MpsPriority                 bool                   `json:"MpsPriority,omitempty" bson:"MpsPriority,omitempty"`
	McsPriority                 bool                   `json:"McsPriority,omitempty" bson:"McsPriority,omitempty"`
	ActiveTime                  uint                   `json:"ActiveTime,omitempty" bson:"ActiveTime,omitempty"`
	DlPacketCount               int                    `json:"DlPacketCount,omitempty" bson:"DlPacketCount,omitempty"`
	MicoAllowed                 bool                   `json:"MicoAllowed,omitempty" bson:"MicoAllowed,omitempty"`
	SharedAmDataIds             []string               `json:"SharedAmDataIds,omitempty" bson:"SharedAmDataIds,omitempty"`
	OdbPacketServices           OdbPacketServices      `json:"OdbPacketServices,omitempty" bson:"OdbPacketServices,omitempty"`
	SubscribedDnnList           []string               `json:"SubscribedDnnList,omitempty" bson:"SubscribedDnnList,omitempty"`
	NssaiInclusionAllowed       bool                   `json:"NssaiInclusionAllowed,omitempty" bson:"NssaiInclusionAllowed,omitempty"`
}

// GetAMSData returns the AMS data
func GetAMSData(supi string) *AccessAndMobilitySubscriptionData {

	collection := db.NewDBClient().Database("Repository").Collection("accessandmobilitysubscriptiondata")
	filter := bson.D{primitive.E{Key: "Supi", Value: supi}}

	var result AccessAndMobilitySubscriptionData

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

// template/example data
var ExAMSData = AccessAndMobilitySubscriptionData{
	Supi:                        "892881507",
	SupportedFeatures:           "D6e683bAfB2bbdCa",
	Gpsis:                       []string{"msisdn-61738596071"},
	InternalGroupIds:            []string{"056eEbD4-310-248-4963a9e7B4ee"},
	SubscribedUeAmbr:            AmbrData,
	Nssai:                       NssaiData,
	RatRestrictions:             []RatType{RatTypeData},
	ForbiddenAreas:              []Area{AreaData1},
	ServiceAreaRestriction:      ServiceAreaRestrictionData,
	CoreNetworkTypeRestrictions: []CoreNetworkType{CoreNetworkTypeData},
	RfspIndex:                   1,
	SubsRegTimer:                10000,
	UeUsageType:                 00010100,
	MpsPriority:                 false,
	McsPriority:                 false,
	ActiveTime:                  10000,
	DlPacketCount:               0,
	// sorInfo: ,
	// upuInfo: ,
	MicoAllowed:           false,
	SharedAmDataIds:       []string{"056eEbD4-310-248-4963a9e7B4ee"},
	OdbPacketServices:     OdbPacketServicesData,
	SubscribedDnnList:     []string{"networka", "networkb"},
	NssaiInclusionAllowed: false,
}
