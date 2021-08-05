package requests

// creates request to update amf access registration data

import (
	"TM5GProject/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// creates and sends amf access registration data
func (c *Client) UpdateAmfRegData(ctx context.Context, ueid string, pdata models.Amf3gppAccessRegistration) error {

	pd, err := json.Marshal(pdata)
	if err != nil {
		return err
	}
	upd := bytes.NewReader(pd)

	// prepare request
	reqValues := []interface{}{c.BaseURL, ueid}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/registrations/amf-3gpp-access", reqValues...), upd)
	if err != nil {
		return err
	}

	// add context
	req = req.WithContext(ctx)

	// send request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// error handling of response
	if res.StatusCode == 204 || res.StatusCode == 400 || res.StatusCode == 404 || res.StatusCode == 500 || res.StatusCode == 503 {
		return err
	}

	return nil
}
