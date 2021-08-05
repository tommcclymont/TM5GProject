package requests

// transfer ue context from old amf to new amf

import (
	"TM5GProject/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// creates and sends ue context data retrieval request
func (c *Client) TransferUeContext(ctx context.Context, uecontextid string, pdata models.UeContextTransferReq) (*models.UeContextTransferRsp, error) {

	pd, err := json.Marshal(pdata)
	if err != nil {
		return nil, err
	}
	upd := bytes.NewReader(pd)

	// prepare request
	reqValues := []interface{}{c.BaseURL, uecontextid}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ue-contexts/%s/transfer", reqValues...), upd)
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
		var errUeContextData models.UeContextTransferRsp
		if err := json.NewDecoder(res.Body).Decode(&errUeContextData); err != nil {
			return nil, err
		}

		return &errUeContextData, err
	}

	// decode json and return successful response
	var successUeContextData models.UeContextTransferRsp
	if err := json.NewDecoder(res.Body).Decode(&successUeContextData); err != nil {
		return nil, err
	}

	return &successUeContextData, nil
}
