package rebrandly

import (
	"context"
	"fmt"
)

// LinksGet returns the list of links
func (c *Controller) LinksGet(filters *LinksFilters) (links Links, err error) {
	return c.LinksGetCtx(nil, filters)
}

// LinksGetCtx returns the list of links
func (c *Controller) LinksGetCtx(ctx context.Context, filters *LinksFilters) (links Links, err error) {
	query, err := convertStructToURLQuery(filters)
	if err != nil {
		err = fmt.Errorf("can't convert filters to query params: %v", err)
		return
	}
	url := *templateURL
	url.Path += "/links"
	url.RawQuery = query.Encode()
	err = c.request(ctx, "GET", url, nil, &links)
	return
}
