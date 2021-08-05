package equipmentidentitycheck

// equipment status service

import (
	"github.com/tommcclymont/TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"
)

// EqStats is the struct for this handler
type EqStats struct {
	l *log.Logger
}

// EquipmentStatusRetrieval creates the handler
func EquipmentStatusRetrieval(l *log.Logger) *EqStats {
	return &EqStats{l}
}

func (EqStat *EqStats) GetEquipmentStatus(rw http.ResponseWriter, r *http.Request) {

	// extract values from request path
	pei := r.URL.Query().Get("pei")

	// set specified headers
	rw.Header().Set("Content-Type", "application/json")

	// get data and convert to json
	eirresponselist := models.GetEirResponseData(pei)
	d, err := json.Marshal(eirresponselist)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// write json for response
	rw.Write(d)
}
