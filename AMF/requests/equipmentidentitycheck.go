package requests

// creates equipment status request

import (
	"github.com/tommcclymont/TM5GProject/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// structure for specified optional parameters in EIR status request
type EIROptParams struct {
	Pei               string
	Supi              string
	Gpsi              string
	SupportedFeatures string
}

// creates and sends equipment status retrieval request
func (c *Client) GetEquipmentStatus(ctx context.Context, EIROptParams *EIROptParams) (*models.EirResponseData, error) {

	// prepare request
	reqValues := []interface{}{c.BaseURL}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/equipment-status", reqValues...), nil)
	if err != nil {
		return nil, err
	}

	// add context
	req = req.WithContext(ctx)

	// set specified query parameters
	q := req.URL.Query()

	if EIROptParams.Pei != "" {
		q.Add("pei", fmt.Sprintf("%s", EIROptParams.Pei))
	}

	if EIROptParams.Supi != "" {
		q.Add("supi", fmt.Sprintf("%s", EIROptParams.Supi))
	}

	if EIROptParams.Gpsi != "" {
		q.Add("gpsi", fmt.Sprintf("%s", EIROptParams.Gpsi))
	}

	if EIROptParams.SupportedFeatures != "" {
		q.Add("supported-features", fmt.Sprintf("%s", EIROptParams.SupportedFeatures))
	}

	req.URL.RawQuery = q.Encode()

	// send request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// error handling of response
	if res.StatusCode == 400 || res.StatusCode == 404 || res.StatusCode == 500 || res.StatusCode == 503 {
		var errEIRData models.EirResponseData
		if err := json.NewDecoder(res.Body).Decode(&errEIRData); err != nil {
			return nil, err
		}

		return &errEIRData, err
	}

	// decode json and return successful response
	var successEIRData models.EirResponseData
	if err := json.NewDecoder(res.Body).Decode(&successEIRData); err != nil {
		return nil, err
	}

	return &successEIRData, nil
}
