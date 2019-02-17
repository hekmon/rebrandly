package rebrandly

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL           = "https://api.rebrandly.com/v1"
	jsonContentType   = "application/json"
	contentTypeHeader = "content-type"
)

func (c *Controller) request(ctx context.Context, verb, URI string, payload, answer interface{}) (err error) {
	var bodySource io.Reader
	// Create payload if necessary
	if payload != nil {
		var data []byte
		if data, err = json.Marshal(payload); err != nil {
			err = fmt.Errorf("can't marshall body data as JSON: %v", err)
			return
		}
		bodySource = bytes.NewReader(data)
	}
	// Create request
	req, err := http.NewRequest(verb, fmt.Sprintf("%s/%s", baseURL, URI), bodySource)
	if err != nil {
		err = fmt.Errorf("can't prepare the request: %v", err)
		return
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	// Set headers
	if payload != nil {
		req.Header.Set(contentTypeHeader, jsonContentType)
	}
	req.Header.Set("apikey", c.apiKey)
	// Execute
	resp, err := c.client.Do(req)
	if err != nil {
		err = fmt.Errorf("request execution error: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("request failed: %s", resp.Status)
		return
	}
	// Unmarshall JSON
	if ct := resp.Header.Get(contentTypeHeader); ct != jsonContentType {
		err = fmt.Errorf("response %s is invalid (expecting '%s'): %s", contentTypeHeader, jsonContentType, ct)
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(answer); err != nil {
		err = fmt.Errorf("unmarshalling response as JSON failed: %v", err)
	}
	return
}
