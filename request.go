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
	// Create payload if necessary
	bodySource, err := createBodyReader(payload)
	if err != nil {
		return
	}
	// Create request
	if debug {
		URL.Scheme = templateDebugURL.Scheme
		URL.Host = templateDebugURL.Host
	}
	req, err := http.NewRequestWithContext(ctx, verb, URL.String(), bodySource)
	if err != nil {
		err = fmt.Errorf("can't prepare the request: %v", err)
		return
	}
	// Set headers
	c.setRequestHeaders(req)
	// Execute
	resp, err := c.client.Do(req)
	if err != nil {
		err = fmt.Errorf("request execution error: %v", err)
		return
	}
	defer resp.Body.Close()
	// Handles http & api errors
	if err = handleHTTPErrors(resp, supportedErrors); err != nil {
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

func createBodyReader(payload interface{}) (bodySource io.Reader, err error) {
	if !isPayloadUsable(payload) {
		return
	}
	var data []byte
	if data, err = json.Marshal(payload); err != nil {
		err = fmt.Errorf("can't marshall body data as JSON: %v", err)
		return
	}
	bodySource = bytes.NewReader(data)
	return
}

func isPayloadUsable(payload interface{}) bool {
	return payload != nil && (reflect.ValueOf(payload).Kind() != reflect.Ptr || !reflect.ValueOf(payload).IsNil())
}

func (c *Controller) setRequestHeaders(req *http.Request) {
	if req == nil {
		return
	}
	if req.Body != nil && req.Body != http.NoBody {
		req.Header.Set(contentTypeHeaderName, jsonContentType)
	}
	if workspace := c.GetWorkspace(); workspace != "" {
		req.Header.Set(workspaceHeaderName, workspace)
	}
	if ua := c.GetUserAgent(); ua != "" {
		req.Header.Set(userAgentHeaderName, ua)
	}
	req.Header.Set(apikeyHeaderName, c.apiKey)
}

func handleHTTPErrors(resp *http.Response, supportedErrors []int) (err error) {
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
