package models

import (
	db "TM5GProject/DB"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// structure for amf3gppaccessregistation data as defined in specification 29.503 release 15
type Amf3gppAccessRegistration struct {
	UeId                        string               `json:"UeId,omitempty" bson:"UeId,omitempty"`
	AmfInstanceId               string               `json:"AmfInstanceId,omitempty" bson:"AmfInstanceId,omitempty"`
	DeregCallbackUri            string               `json:"DeregCallbackUri,omitempty" bson:"DeregCallbackUri,omitempty"`
	Guami                       Guami                `json:"Guami,omitempty" bson:"Guami,omitempty"`
	RatType                     RatType              `json:"RatType,omitempty" bson:"RatType,omitempty"`
	SupportedFeatures           string               `json:"SupportedFeatures,omitempty" bson:"SupportedFeatures,omitempty"`
	PurgeFlag                   bool                 `json:"PurgeFlag,omitempty" bson:"PurgeFlag,omitempty"`
	Pei                         string               `json:"Pei,omitempty" bson:"Pei,omitempty"`
	ImsVoPs                     string               `json:"ImsVoPs,omitempty" bson:"ImsVoPs,omitempty"`
	AmfServiceNameDereg         string               `json:"AmfServiceNameDereg,omitempty" bson:"AmfServiceNameDereg,omitempty"`
	PcscfRestorationCallbackUri string               `json:"PcscfRestorationCallbackUri,omitempty" bson:"PcscfRestorationCallbackUri,omitempty"`
	AmfServiceNamePcscfRest     string               `json:"AmfServiceNamePcscfRest,omitempty" bson:"AmfServiceNamePcscfRest,omitempty"`
	InitialRegistrationInd      bool                 `json:"InitialRegistrationInd,omitempty" bson:"InitialRegistrationInd,omitempty"`
	BackupAmfInfo               BackupAmfInfo        `json:"BackupAmfInfo,omitempty" bson:"BackupAmfInfo,omitempty"`
	DrFlag                      bool                 `json:"DrFlag,omitempty" bson:"DrFlag,omitempty"`
	UrrpIndicator               bool                 `json:"UrrpIndicator,omitempty" bson:"UrrpIndicator,omitempty"`
	AmfEeSubscriptionId         string               `json:"AmfEeSubscriptionId,omitempty" bson:"AmfEeSubscriptionId,omitempty"`
	EpsInterworkingInfo         map[string]EpsIwkPgw `json:"EpsInterworkingInfo,omitempty" bson:"EpsInterworkingInfo,omitempty"`
}

var ErrUeIdNotFound = fmt.Errorf("UEID not found")
var ErrSupiNotFound = fmt.Errorf("SUPI not found")

// update data
func PutAmf3gpp(ueid string, d Amf3gppAccessRegistration) error {

	upd := bson.M{"$set": d}
	collection := db.NewDBClient().Database("Repository").Collection("amf3gppaccessregistration")
	filter := bson.D{primitive.E{Key: "UeId", Value: ueid}}

	_, err := collection.UpdateOne(context.TODO(), filter, upd)
	if err != nil {
		return ErrUeIdNotFound
	}

	return nil
}

// template/example data
var ExAmf3gppData = Amf3gppAccessRegistration{
	AmfInstanceId:               "cf785448-6090-47bf-b28b-7371e2a737c6",
	DeregCallbackUri:            "https://127.0.0.78:9090/",
	Guami:                       GuamiData,
	RatType:                     RatTypeData,
	SupportedFeatures:           "D6e683bAfB2bbdCa",
	PurgeFlag:                   false,
	Pei:                         "imei-857390213674531",
	ImsVoPs:                     "HOMOGENEOUS_SUPPORT",
	AmfServiceNameDereg:         "testamf",
	PcscfRestorationCallbackUri: "https://127.0.0.78:9090/",
	AmfServiceNamePcscfRest:     "testamf",
	InitialRegistrationInd:      true,
	BackupAmfInfo:               BackupAmfInfoData,
	DrFlag:                      true,
	EpsInterworkingInfo:         EpsIwkPgwMap,
}
