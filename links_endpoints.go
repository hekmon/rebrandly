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

// LinksGetByID returns the link details of link id.
func (c *Controller) LinksGetByID(id string) (link Link, err error) {
	return c.LinksGetByIDCtx(nil, id)
}

// LinksGetByIDCtx returns the link details of link id.
func (c *Controller) LinksGetByIDCtx(ctx context.Context, id string) (link Link, err error) {
	url := *templateURL
	url.Path += fmt.Sprintf("/links/%s", id)
	err = c.request(ctx, "GET", url, nil, &link)
	return
}
