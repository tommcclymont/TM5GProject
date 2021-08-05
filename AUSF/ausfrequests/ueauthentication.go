package ausfrequests

// creates authentication info request (AUSF to UDM)

import (
	"TM5GProject/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// creates and sends authentication info retrieval request
func (c *Client) AuthenticateUe(ctx context.Context, supi string, pdata models.AuthenticationInfoRequest) (*models.AuthenticationInfoResult, error) {

	pd, err := json.Marshal(pdata)
	if err != nil {
		return nil, err
	}
	upd := bytes.NewReader(pd)

	// prepare request
	reqValues := []interface{}{c.BaseURL, supi}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/security-information/generate-auth-data", reqValues...), upd)
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
		var errUeAuthenticationData models.AuthenticationInfoResult
		if err := json.NewDecoder(res.Body).Decode(&errUeAuthenticationData); err != nil {
			return nil, err
		}

		return &errUeAuthenticationData, err
	}

	// decode json and return successful response
	var successUeAuthenticationData models.AuthenticationInfoResult
	if err := json.NewDecoder(res.Body).Decode(&successUeAuthenticationData); err != nil {
		return nil, err
	}

	return &successUeAuthenticationData, nil
}
