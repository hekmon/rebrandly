package rebrandly

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

const (
	baseURL               = "https://api.rebrandly.com/v1"
	jsonContentType       = "application/json; charset=utf-8"
	contentTypeHeaderName = "content-type"
	apikeyHeaderName      = "apikey"
	workspaceHeaderName   = "workspace"
	debug                 = false
)

func (c *Controller) request(ctx context.Context, verb, URI string, payload, answer interface{}) (err error) {
	var bodySource io.Reader
	// Create payload if necessary
	if isPayloadUseable(payload) {
		var data []byte
		if data, err = json.Marshal(payload); err != nil {
			err = fmt.Errorf("can't marshall body data as JSON: %v", err)
			return
		}
		if debug {
			fmt.Println("Rebrandly API| payload:", string(data))
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
	if isPayloadUseable(payload) {
		req.Header.Set(contentTypeHeaderName, jsonContentType)
	}
	if workspace := c.GetWorkspace(); workspace != "" {
		req.Header.Set(workspaceHeaderName, workspace)
	}
	req.Header.Set(apikeyHeaderName, c.apiKey)
	// Last chance to inspect
	if debug {
		fmt.Printf("Rebrandly API| request:%+v\n", req)
	}
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
	if ct := resp.Header.Get(contentTypeHeaderName); ct != jsonContentType {
		err = fmt.Errorf("response %s is invalid (expecting '%s'): %s", contentTypeHeaderName, jsonContentType, ct)
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(answer); err != nil {
		err = fmt.Errorf("unmarshalling response as JSON failed: %v", err)
	}
	return
}

func isPayloadUseable(payload interface{}) bool {
	return payload != nil && (reflect.ValueOf(payload).Kind() != reflect.Ptr || !reflect.ValueOf(payload).IsNil())
}
