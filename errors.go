package rebrandly

// Error represents an error returned by rebrandly
type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Source  string `json:"source"`
	Errors  []struct {
		Code     string `json:"code"`
		Property string `json:"property"`
		Message  string `json:"message"`
		Verbose  string `json:"verbose"`
	} `json:"errors"`
}

func (e Error) Error() string {
	return e.Message
}
