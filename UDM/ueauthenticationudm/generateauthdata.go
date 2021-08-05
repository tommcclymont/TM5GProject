package ueauthenticationudm

// ue authentication udm service

import (
	"TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// genauths is the struct for this handler
type genauths struct {
	l *log.Logger
}

// AuthGeneration creates the handler
func AuthGeneration(l *log.Logger) *genauths {
	return &genauths{l}
}

func (genauth genauths) GenerateAuth(rw http.ResponseWriter, r *http.Request) {

	var generatedata models.AuthenticationInfoRequest
	err := json.NewDecoder(r.Body).Decode(&generatedata)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// set specified headers
	rw.Header().Set("Content-Type", "application/json")

	// extract supi from request path
	vars := mux.Vars(r)
	supi := vars["supi"]
	generatedata.Supi = supi

	// create authentication request data
	err = models.PostGenAuth(supi, generatedata)
	if err != nil {
		http.Error(rw, "Supi not found", http.StatusInternalServerError)
		return
	}

	// get authenication info for response
	AuthenticationInfoResultData := models.GetAuthInfoRes(supi)
	d, err := json.Marshal(AuthenticationInfoResultData)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// write json for response
	rw.Write(d)
}
