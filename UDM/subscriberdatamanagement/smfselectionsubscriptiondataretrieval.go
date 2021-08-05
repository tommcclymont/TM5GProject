package subscriberdatamanagement

// smf selection subscription data retrieval service

import (
	"github.com/tommcclymont/TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Smfsub is the struct for this handler
type Smfsub struct {
	l *log.Logger
}

// SmfSubDataRetrieval creates the handler
func SmfSubDataRetrieval(l *log.Logger) *Smfsub {
	return &Smfsub{l}
}

// GetSMFSubData returns the SMF selection subscription data
func (smfsub *Smfsub) GetSMFSubData(rw http.ResponseWriter, r *http.Request) {

	// extract supi from request path
	vars := mux.Vars(r)
	supi := vars["supi"]

	// set specified headers
	rw.Header().Set("Cache-Control", "max-age:100000000")
	rw.Header().Set("Last-Modified", time.Now().Format(http.TimeFormat))
	rw.Header().Set("Etag", "smfsub")
	rw.Header().Set("Content-Type", "application/json")

	// get data and convert to json
	smfsublist := models.GetSMFSubData(supi)
	d, err := json.Marshal(smfsublist)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// write json for response
	rw.Write(d)
}
