package pcfrequests

// creates event exposure subscription

import (
	"TM5GProject/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// creates and sends event exposure subscription data
func (c *Client) CreateEventExposeSub(ctx context.Context, pdata models.AmfCreateEventSubscription) (*models.AmfCreatedEventSubscription, error) {

	pd, err := json.Marshal(pdata)
	if err != nil {
		return nil, err
	}
	upd := bytes.NewReader(pd)

	// prepare request
	reqValues := []interface{}{c.BaseURL}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/subscriptions", reqValues...), upd)
	if err != nil {
		return nil, err
	}

	// add context
	req = req.WithContext(ctx)

	// send request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// error handling of response
	if res.StatusCode == 204 || res.StatusCode == 400 || res.StatusCode == 404 || res.StatusCode == 500 || res.StatusCode == 503 {
		var errEventExposeSubData models.AmfCreatedEventSubscription
		if err := json.NewDecoder(res.Body).Decode(&errEventExposeSubData); err != nil {
			return nil, err
		}

		return &errEventExposeSubData, err
	}

	// decode json and return successful response
	var successEventExposeSubData models.AmfCreatedEventSubscription
	if err := json.NewDecoder(res.Body).Decode(&successEventExposeSubData); err != nil {
		return nil, err
	}

	return &successEventExposeSubData, nil
}
