package rebrandly

import (
	"context"
)

// Account returns the current account metata
func (c *Controller) Account() (account Account, err error) {
	return c.AccountCtx(nil)
}

// AccountCtx returns the current account metata
func (c *Controller) AccountCtx(ctx context.Context) (account Account, err error) {
	url := *templateURL
	url.Path += "/account"
	err = c.request(ctx, "GET", url, nil, &account)
	return
}
