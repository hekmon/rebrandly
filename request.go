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
	userAgentHeaderName   = "user-agent"
	apikeyHeaderName      = "apikey"
	workspaceHeaderName   = "workspace"
	debug                 = false
	pipeDreamDebugging    = "https://en6uhmjiqbxd5.x.pipedream.net/v1"
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
		bodySource = bytes.NewReader(data)
	}
	// Create request
	var URL string
	if debug {
		URL = fmt.Sprintf("%s/%s", pipeDreamDebugging, URI)
	} else {
		URL = fmt.Sprintf("%s/%s", baseURL, URI)
	}
	req, err := http.NewRequest(verb, URL, bodySource)
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
	if ua := c.GetUserAgent(); ua != "" {
		req.Header.Set(userAgentHeaderName, ua)
	}
	req.Header.Set(apikeyHeaderName, c.apiKey)
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
