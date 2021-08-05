package requests

// creates policy association data

import (
	"TM5GProject/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// creates and sends policy association data
func (c *Client) CreatePolAssoData(ctx context.Context, pdata models.PolicyAssociationRequest) (*models.PolicyAssociation, error) {

	pd, err := json.Marshal(pdata)
	if err != nil {
		return nil, err
	}
	upd := bytes.NewReader(pd)

	// prepare request
	reqValues := []interface{}{c.BaseURL}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/policies", reqValues...), upd)
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
		var errPolAssoData models.PolicyAssociation
		if err := json.NewDecoder(res.Body).Decode(&errPolAssoData); err != nil {
			return nil, err
		}

		return &errPolAssoData, err
	}

	// decode json and return successful response
	var successPolAssoData models.PolicyAssociation
	if err := json.NewDecoder(res.Body).Decode(&successPolAssoData); err != nil {
		return nil, err
	}

	return &successPolAssoData, nil
}
