package rebrandly

import (
	"context"
)

func (c *Controller) Workspaces() (workspaces []Workspace, err error) {
	url := *templateURL
	url.Path += "/workspaces"
	err = c.request(context.Background(), "GET", url, nil, &workspaces, []int{})
	return
}
