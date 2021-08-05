package eventexposure

// event exposure subscription service

import (
	"TM5GProject/models"
	"encoding/json"
	"log"
	"net/http"
)

// evsubs is the struct for this handler
type evsubs struct {
	l *log.Logger
}

// EventExposeSubscription creates the handler
func EventExposeSubscription(l *log.Logger) *evsubs {
	return &evsubs{l}
}

// SubscribeEventExpose creates the subscription and returns the response data
func (evsub evsubs) SubscribeEventExpose(rw http.ResponseWriter, r *http.Request) {

	var subscriptiondata models.AmfCreateEventSubscription
	err := json.NewDecoder(r.Body).Decode(&subscriptiondata)
	// err = json.Unmarshal(bd, &subscriptiondata)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// set specified headers
	rw.Header().Set("Location", "https://127.0.0.78:9090/npcf-evts/v2/subscriptions/dummysubid")
	rw.Header().Set("Content-Type", "application/json")

	subid := "dummysubid"
	subscriptiondata.SubscriptionId = subid

	// create subscription
	err = models.PostEventExposeSub(subid, subscriptiondata)
	if err != nil {
		http.Error(rw, "Subscription Id not found", http.StatusInternalServerError)
		return
	}

	// get response data and convert to json
	eventexposesublist := models.GetEventExposeSubData(subid)
	d, err := json.Marshal(eventexposesublist)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// write json for response
	rw.Write(d)
}
