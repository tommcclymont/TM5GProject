package oldamfrequests

// creates request to release OldAMF's sm context

import (
	"github.com/tommcclymont/TM5GProject/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// creates and sends sm context release request
func (c *Client) ReleaseSmContext(ctx context.Context, smcontextref string, pdata models.SmContextRelease) error {

	pd, err := json.Marshal(pdata)
	if err != nil {
		return err
	}
	upd := bytes.NewReader(pd)

	// prepare request
	reqValues := []interface{}{c.BaseURL, smcontextref}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/sm-contexts/%s/release", reqValues...), upd)
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

	return nil
}
