package requests

// creates authenticate ue requests

import (
	"github.com/tommcclymont/TM5GProject/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// creates and sends ue authentication request
func (c *Client) AuthenticateUe(ctx context.Context, pdata models.AuthenticationInfo) (*models.UeAuthenticationCtx, error) {

	pd, err := json.Marshal(pdata)
	if err != nil {
		return nil, err
	}
	upd := bytes.NewReader(pd)

	// prepare request
	reqValues := []interface{}{c.BaseURL}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ue-authentications", reqValues...), upd)
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
		var errUeAuthenticateData models.UeAuthenticationCtx
		if err := json.NewDecoder(res.Body).Decode(&errUeAuthenticateData); err != nil {
			return nil, err
		}

		return &errUeAuthenticateData, err
	}

	// decode json and return successful response
	var successUeAuthenticateData models.UeAuthenticationCtx
	if err := json.NewDecoder(res.Body).Decode(&successUeAuthenticateData); err != nil {
		return nil, err
	}

	return &successUeAuthenticateData, nil
}
