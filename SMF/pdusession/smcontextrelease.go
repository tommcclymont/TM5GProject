package pdusession

// sm context release service

import (
	"github.com/tommcclymont/TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// smctxrels is the struct for this handler
type smctxrels struct {
	l *log.Logger
}

// SmContextRelease creates the handler
func SmContextRelease(l *log.Logger) *smctxrels {
	return &smctxrels{l}
}

func (smctxrel smctxrels) ReleaseSmContext(rw http.ResponseWriter, r *http.Request) {

	var releasedata models.SmContextRelease
	err := json.NewDecoder(r.Body).Decode(&releasedata)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// set specified headers
	rw.Header().Set("Content-Type", "application/json")

	// extract sm context ref from request path
	vars := mux.Vars(r)
	smcontextref := vars["smcontextref"]
	releasedata.SmContextRef = smcontextref

	// create sm context release request data
	err = models.PostSmCtxRel(smcontextref, releasedata)
	if err != nil {
		http.Error(rw, "SM Context Ref not found", http.StatusInternalServerError)
		return
	}

	// release sm context
	err = models.ReleaseSmCtx(smcontextref)
	if err != nil {
		http.Error(rw, "SM Context Ref not found", http.StatusInternalServerError)
		return
	}
}
