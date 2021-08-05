package subscriberdatamanagement

// ue context in smf data retrieval service

import (
	"github.com/tommcclymont/TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// uesmf is the struct for this handler
type uesmf struct {
	l *log.Logger
}

// UeCtxSmfDataRetrieval creates the handler
func UeCtxSmfDataRetrieval(l *log.Logger) *uesmf {
	return &uesmf{l}
}

// GetUeCtxSmfData returns the ue context in smf data
func (uesmf *uesmf) GetUeCtxSmfData(rw http.ResponseWriter, r *http.Request) {

	// extract supi from request path
	vars := mux.Vars(r)
	supi := vars["supi"]

	rw.Header().Set("Content-Type", "application/json")

	// get data and convert to json
	uectxsmflist := models.GetUeCtxSmfData(supi)
	d, err := json.Marshal(uectxsmflist)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// write json for response
	rw.Write(d)
}
