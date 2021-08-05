package models

import (
	db "TM5GProject/DB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// structure for UE context in SMF data as defined in specification 29.503 release 15
type UeContextInSmfData struct {
	Supi          string                `json:"Supi,omitempty" bson:"Supi,omitempty"`
	PduSessions   map[string]PduSession `json:"PduSessions,omitempty" bson:"PduSessions,omitempty"`
	PgwInfo       []PgwInfo             `json:"PgwInfo,omitempty" bson:"PgwInfo,omitempty"`
	EmergencyInfo EmergencyInfo         `json:"EmergencyInfo,omitempty" bson:"EmergencyInfo,omitempty"`
}

// get data from DB
func GetUeCtxSmfData(supi string) *UeContextInSmfData {

	collection := db.NewDBClient().Database("Repository").Collection("uecontextinsmfdata")
	filter := bson.D{primitive.E{Key: "Supi", Value: supi}}

	var result UeContextInSmfData

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

// template/example data
var ExUeSmfCtx = UeContextInSmfData{
	Supi:          "892881507",
	PduSessions:   PduSessionsMap,
	PgwInfo:       []PgwInfo{PgwInfoData},
	EmergencyInfo: EmergencyInfoData,
}
