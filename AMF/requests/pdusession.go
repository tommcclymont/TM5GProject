package requests

// create sm context for pdu session

import (
	"github.com/tommcclymont/TM5GProject/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// creates and sends sm context data
func (c *Client) CreateSmContext(ctx context.Context, pdata models.SmContextCreate) (*models.SmContextCreated, error) {

	pd, err := json.Marshal(pdata)
	if err != nil {
		return nil, err
	}
	upd := bytes.NewReader(pd)

	// prepare request
	reqValues := []interface{}{c.BaseURL}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/sm-contexts", reqValues...), upd)
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
		var errSmContextData models.SmContextCreated
		if err := json.NewDecoder(res.Body).Decode(&errSmContextData); err != nil {
			return nil, err
		}

		return &errSmContextData, err
	}

	// decode json and return successful response
	var successSmContextData models.SmContextCreated
	if err := json.NewDecoder(res.Body).Decode(&successSmContextData); err != nil {
		return nil, err
	}

	return &successSmContextData, nil
}
