package subscriberdatamanagement

// access and mobility subscription data retrieval service

import (
	"github.com/tommcclymont/TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Amfs is the struct for this handler
type Amfs struct {
	l *log.Logger
}

// AmsDataRetrieval creates the handler
func AmsDataRetrieval(l *log.Logger) *Amfs {
	return &Amfs{l}
}

// GetAMSData returns the AMS data
func (ams *Amfs) GetAMSData(rw http.ResponseWriter, r *http.Request) {

	// extract supi from request path
	vars := mux.Vars(r)
	supi := vars["supi"]

	// set specified headers
	rw.Header().Set("Cache-Control", "max-age:100000000")
	rw.Header().Set("Last-Modified", time.Now().Format(http.TimeFormat))
	rw.Header().Set("Etag", "ams")
	rw.Header().Set("Content-Type", "application/json")

	// get data and convert to json
	amslist := models.GetAMSData(supi)
	d, err := json.Marshal(amslist)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// write json for response
	rw.Write(d)
}
