package uecontextmanagement

// amf3gppaccessregistration data update service

import (
	"TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Amfregs is the struct for this handler
type Amfregs struct {
	l *log.Logger
}

// AmsDataRetrieval creates the handler
func AmsRegDataUpdate(l *log.Logger) *Amfregs {
	return &Amfregs{l}
}

// GetAMSData returns the AMS data
func (amfreg Amfregs) UpdateAmsRegData(rw http.ResponseWriter, r *http.Request) {

	// extract supi from request path
	vars := mux.Vars(r)
	ueid := vars["ueid"]
	var updatedata models.Amf3gppAccessRegistration
	err := json.NewDecoder(r.Body).Decode(&updatedata)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// set specified headers
	rw.Header().Set("Location", "Contains the URI of the newly created resource, according to the structure: {apiRoot}/nudm-uecm/v1/{ueId}/registrations/amf-3gpp-access")
	rw.Header().Set("Content-Type", "application/json")

	updatedata.UeId = ueid

	// update data
	err = models.PutAmf3gpp(ueid, updatedata)
	if err != nil {
		http.Error(rw, "SUPI not found", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
