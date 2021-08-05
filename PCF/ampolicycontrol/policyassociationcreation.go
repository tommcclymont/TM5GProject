package ampolicycontrol

// policy association creation service

import (
	"TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"
)

// pols is the struct for this handler
type pols struct {
	l *log.Logger
}

// PolDataCreation creates the handler
func PolDataCreation(l *log.Logger) *pols {
	return &pols{l}
}

func (pol pols) CreatePolData(rw http.ResponseWriter, r *http.Request) {

	var creationdata models.PolicyAssociationRequest
	err := json.NewDecoder(r.Body).Decode(&creationdata)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// set specified headers
	rw.Header().Set("Location", "https://127.0.0.79:9090/npcf-am-policy-control/v2/policies/dummypolassoid")
	rw.Header().Set("Content-Type", "application/json")

	polassoid := "dummypolassoid"
	creationdata.PolAssoId = polassoid

	// create policy association request data
	err = models.PostPolAsso(polassoid, creationdata)
	if err != nil {
		http.Error(rw, "Policy Association Id not found", http.StatusInternalServerError)
		return
	}

	// get policy association data and convert to json
	polassolist := models.GetPolAssoData(polassoid)
	d, err := json.Marshal(polassolist)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// write json for response
	rw.Write(d)
}
