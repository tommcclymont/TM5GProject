package pdusession

// sm context creation service

import (
	"TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"
)

// smctxs is the struct for this handler
type smctxs struct {
	l *log.Logger
}

// SmContextCreate creates the handler
func SmContextCreate(l *log.Logger) *smctxs {
	return &smctxs{l}
}

func (smctx smctxs) CreateSmContext(rw http.ResponseWriter, r *http.Request) {

	var createdata models.SmContextCreate
	err := json.NewDecoder(r.Body).Decode(&createdata)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// set specified headers
	rw.Header().Set("Location", "https://127.0.0.82:9090/nsmf-pdusession/v2/sm-contexts/dummysmcontextref")
	rw.Header().Set("Content-Type", "application/json")

	smcontextref := "dummysmcontextref"
	createdata.SmContextRef = smcontextref

	// create sm context
	err = models.PostSmCtx(smcontextref, createdata)
	if err != nil {
		http.Error(rw, "SM Context Ref not found", http.StatusInternalServerError)
		return
	}

	// get created data for response
	smctxlist := models.GetSMCtxCreatedData(smcontextref)
	d, err := json.Marshal(smctxlist)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// write json for response
	rw.Write(d)
}
