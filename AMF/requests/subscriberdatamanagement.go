package requests

// creates subscriber data management requests from AMF to UDM

import (
	"github.com/tommcclymont/TM5GProject/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// structure for specified optional parameters in AM data request
type AMOptParams struct {
	IfNoneMatch       *string
	IfModifiedSince   *string
	SupportedFeatures *string
	PlmnId            *string
	AdjacentPlmns     *[]string
}

// structure for specified optional parameters in smf data request
type SMFOptParams struct {
	IfNoneMatch       *string
	IfModifiedSince   *string
	SupportedFeatures *string
	PlmnId            *string
}

// structure for specified optional parameters in smf data request
type UeCtxSmfOptParams struct {
	SupportedFeatures *string
}

// creates and sends access and mobility subscription data retrieval request
func (c *Client) GetAMData(ctx context.Context, supi string, AMOptParams *AMOptParams) (*models.AccessAndMobilitySubscriptionData, error) {

	// prepare request
	reqValues := []interface{}{c.BaseURL, supi}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/am-data", reqValues...), nil)
	if err != nil {
		return nil, err
	}

	// add context
	req = req.WithContext(ctx)

	// set specified header parameters
	if AMOptParams.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", fmt.Sprintf("%s", *AMOptParams.IfNoneMatch))
	}

	if AMOptParams.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", fmt.Sprintf("%s", *AMOptParams.IfModifiedSince))
	}

	if AMOptParams.IfModifiedSince != nil {
		req.Header.Set("Content-Type", fmt.Sprintf("application/json"))
	}

	// set specified query parameters
	q := req.URL.Query()

	if AMOptParams.SupportedFeatures != nil {
		q.Add("supported-features", fmt.Sprintf("%s", *AMOptParams.SupportedFeatures))
	}

	if AMOptParams.PlmnId != nil {
		q.Add("plmn-id", fmt.Sprintf("%s", *AMOptParams.PlmnId))
	}

	if AMOptParams.AdjacentPlmns != nil {
		q.Add("adjacentPlmns", fmt.Sprintf("%s", *AMOptParams.AdjacentPlmns))
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
		var errAMData models.AccessAndMobilitySubscriptionData
		if err := json.NewDecoder(res.Body).Decode(&errAMData); err != nil {
			return nil, err
		}

		return &errAMData, err
	}

	// decode json and return successful response
	var successAMData models.AccessAndMobilitySubscriptionData
	if err := json.NewDecoder(res.Body).Decode(&successAMData); err != nil {
		return nil, err
	}

	return &successAMData, nil
}

// creates and sends smf selection subscription data retrieval request
func (c *Client) GetSMFSelSubData(ctx context.Context, supi string, SMFOptParams *SMFOptParams) (*models.SmfSelectionSubscriptionData, error) {

	// prepare request
	reqValues := []interface{}{c.BaseURL, supi}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/smf-select-data", reqValues...), nil)
	if err != nil {
		return nil, err
	}

	// add context
	req = req.WithContext(ctx)

	// set specified header parameters
	if SMFOptParams.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", fmt.Sprintf("%s", *SMFOptParams.IfNoneMatch))
	}

	if SMFOptParams.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", fmt.Sprintf("%s", *SMFOptParams.IfModifiedSince))
	}

	if SMFOptParams.IfModifiedSince != nil {
		req.Header.Set("Content-Type", fmt.Sprintf("application/json"))
	}

	// set specified query parameters
	q := req.URL.Query()

	if SMFOptParams.SupportedFeatures != nil {
		q.Add("supported-features", fmt.Sprintf("%s", *SMFOptParams.SupportedFeatures))
	}

	if SMFOptParams.PlmnId != nil {
		q.Add("plmn-id", fmt.Sprintf("%s", *SMFOptParams.PlmnId))
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
		var errSMFData models.SmfSelectionSubscriptionData
		if err := json.NewDecoder(res.Body).Decode(&errSMFData); err != nil {
			return nil, err
		}

		return &errSMFData, err
	}

	// decode json and return successful response
	var successSMFData models.SmfSelectionSubscriptionData
	if err := json.NewDecoder(res.Body).Decode(&successSMFData); err != nil {
		return nil, err
	}

	return &successSMFData, nil
}

// creates and sends ue context in smf data retrieval request
func (c *Client) GetUeSmfData(ctx context.Context, supi string, UeCtxSmfOptParams *UeCtxSmfOptParams) (*models.UeContextInSmfData, error) {

	// prepare request
	reqValues := []interface{}{c.BaseURL, supi}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/ue-context-in-smf-data", reqValues...), nil)
	if err != nil {
		return nil, err
	}

	// add context
	req = req.WithContext(ctx)

	// set specified query parameters
	q := req.URL.Query()

	if UeCtxSmfOptParams.SupportedFeatures != nil {
		q.Add("supported-features", fmt.Sprintf("%s", *UeCtxSmfOptParams.SupportedFeatures))
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
		var errUeSmfData models.UeContextInSmfData
		if err := json.NewDecoder(res.Body).Decode(&errUeSmfData); err != nil {
			return nil, err
		}

		return &errUeSmfData, err
	}

	// decode json and return successful response
	var successUeSmfData models.UeContextInSmfData
	if err := json.NewDecoder(res.Body).Decode(&successUeSmfData); err != nil {
		return nil, err
	}

	return &successUeSmfData, nil
}
