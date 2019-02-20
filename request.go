package rebrandly

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	pipeDreamDebugging    = "https://en6uhmjiqbxd5.x.pipedream.net"
)

var (
	templateURL      *url.URL
	templateDebugURL *url.URL
)

func init() {
	var err error
	if templateURL, err = url.Parse(baseURL); err != nil {
		panic(fmt.Sprintf("can't parse baseURL '%s' as url.URL: %v", baseURL, err))
	}
	if templateDebugURL, err = url.Parse(pipeDreamDebugging); err != nil {
		panic(fmt.Sprintf("can't parse pipeDreamDebugging '%s' as url.URL: %v", pipeDreamDebugging, err))
	}
}

func (c *Controller) request(ctx context.Context, verb string, URL url.URL, payload, answer interface{}, supportedErrors []int) (err error) {
	var bodySource io.Reader
	// Create payload if necessary
	payloadUsable := payload != nil && (reflect.ValueOf(payload).Kind() != reflect.Ptr || !reflect.ValueOf(payload).IsNil())
	if payloadUsable {
		var data []byte
		if data, err = json.Marshal(payload); err != nil {
			err = fmt.Errorf("can't marshall body data as JSON: %v", err)
			return
		}
		bodySource = bytes.NewReader(data)
	}
	// Create request
	if debug {
		URL.Scheme = templateDebugURL.Scheme
		URL.Host = templateDebugURL.Host
	}
	req, err := http.NewRequest(verb, URL.String(), bodySource)
	if err != nil {
		err = fmt.Errorf("can't prepare the request: %v", err)
		return
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	// Set headers
	if payloadUsable {
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
		if isErrorSupported(resp.StatusCode, supportedErrors) {
			var extendedError Error
			if err = json.NewDecoder(resp.Body).Decode(&extendedError); err != nil {
				err = fmt.Errorf("unmarshalling response error as JSON failed: %v", err)
			} else {
				err = extendedError
			}
		} else {
			err = fmt.Errorf("request failed: %s", resp.Status)
		}
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

func isErrorSupported(statusCode int, supported []int) bool {
	for _, code := range supported {
		if statusCode == code {
			return true
		}
	}
	return false
}
