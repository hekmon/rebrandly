package rebrandly

import (
	"net/http"
	"sync"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// Controller abstracts the access to the cpg API
type Controller struct {
	apiKey          string
	workspace       string
	workspaceAccess sync.RWMutex
	useragent       string
	useragentAccess sync.RWMutex
	client          *http.Client
}

// New returns an initialized and ready to use Controller
func New(APIKey string) *Controller {
	return &Controller{
		apiKey: APIKey,
		client: cleanhttp.DefaultPooledClient(),
	}
}

// ClearWorkspace allows to reset the controller on the default workspace
func (c *Controller) ClearWorkspace() {
	c.SetWorkspace("")
}

// GetWorkspace returns the current workspace on which the controller work against.
// Empty means default workspace.
func (c *Controller) GetWorkspace() (workspace string) {
	c.workspaceAccess.RLock()
	workspace = c.workspace
	c.workspaceAccess.RUnlock()
	return
}

// SetWorkspace changes the workspace on which the controller will work against.
func (c *Controller) SetWorkspace(workspace string) {
	c.workspaceAccess.Lock()
	c.workspace = workspace
	c.workspaceAccess.Unlock()
}

// SetUserAgent allows tu customize the user agent used by the controller when performing http requests.
func (c *Controller) SetUserAgent(ua string) {
	c.useragentAccess.Lock()
	c.useragent = ua
	c.useragentAccess.Unlock()
}

// GetUserAgent returns the current user agent configured. If empty, golang default ua will be used.
func (c *Controller) GetUserAgent() (ua string) {
	c.useragentAccess.RLock()
	ua = c.useragent
	c.useragentAccess.RUnlock()
	return
}
