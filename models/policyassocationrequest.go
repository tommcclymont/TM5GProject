package models

import (
	db "github.com/tommcclymont/TM5GProject/DB"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// structure for policyassociationrequest data as defined in specification 29.507 release 15
type PolicyAssociationRequest struct {
	PolAssoId        string                 `json:"PolAssoId,omitempty" bson:"PolAssoId,omitempty"`
	NotificationUri  string                 `json:"NotificationUri,omitempty" bson:"NotificationUri,omitempty"`
	AltNotifIpv4Addr string                 `json:"AltNotifIpv4Addr,omitempty" bson:"AltNotifIpv4Addr,omitempty"`
	AltNotifIpv6Addr string                 `json:"AltNotifIpv6Addr,omitempty" bson:"AltNotifIpv6Addr,omitempty"`
	Supi             string                 `json:"Supi,omitempty" bson:"Supi,omitempty"`
	Gpsi             string                 `json:"Gpsi,omitempty" bson:"Gpsi,omitempty"`
	AccessType       string                 `json:"AccessType,omitempty" bson:"AccessType,omitempty"`
	Pei              string                 `json:"Pei,omitempty" bson:"Pei,omitempty"`
	UserLoc          UserLocation           `json:"UserLoc,omitempty" bson:"UserLoc,omitempty"`
	TimeZone         string                 `json:"TimeZone,omitempty" bson:"TimeZone,omitempty"`
	ServingPlmn      PlmnId                 `json:"ServingPlmn,omitempty" bson:"ServingPlmn,omitempty"`
	RatType          RatType                `json:"RatType,omitempty" bson:"RatType,omitempty"`
	GroupIds         []string               `json:"GroupIds,omitempty" bson:"GroupIds,omitempty"`
	ServAreaRes      ServiceAreaRestriction `json:"ServAreaRes,omitempty" bson:"ServAreaRes,omitempty"`
	Rfsp             uint                   `json:"Rfsp,omitempty" bson:"Rfsp,omitempty"`
	Guami            Guami                  `json:"Guami,omitempty" bson:"Guami,omitempty"`
	ServiceName      string                 `json:"ServiceName,omitempty" bson:"ServiceName,omitempty"`
	SuppFeat         string                 `json:"SuppFeat,omitempty" bson:"SuppFeat,omitempty"`
	//TraceReq
}

var ErrPolAssoIdNotFound = fmt.Errorf("Policy Association ID not found")

// insert data to DB
func PostPolAsso(polassoid string, d PolicyAssociationRequest) error {

	collection := db.NewDBClient().Database("Repository").Collection("policyassociationrequest")
	filter := bson.D{primitive.E{Key: "PolAssoId", Value: polassoid}}
	opts := options.Replace().SetUpsert(true)

	_, err := collection.ReplaceOne(context.TODO(), filter, d, opts)
	if err != nil {
		return ErrPolAssoIdNotFound
	}

	return nil
}
