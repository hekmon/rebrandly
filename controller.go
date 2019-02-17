package rebrandly

import (
	"net/http"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// Controller abstracts the access to the cpg API
type Controller struct {
	apiKey string
	client *http.Client
}

// New returns an initialized and ready to use Controller
func New(APIKey string) *Controller {
	return &Controller{
		apiKey: APIKey,
		client: cleanhttp.DefaultPooledClient(),
	}
}
