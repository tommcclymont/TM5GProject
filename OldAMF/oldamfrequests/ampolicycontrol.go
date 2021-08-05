package oldamfrequests

// delete old amf policy association

import (
	"context"
	"fmt"
	"net/http"
)

// creates and sends policy association deletion request
func (c *Client) DeletePolAsso(ctx context.Context, polassoid string) error {

	// prepare request
	reqValues := []interface{}{c.BaseURL, polassoid}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/sm-contexts/%s/release", reqValues...), nil)
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
