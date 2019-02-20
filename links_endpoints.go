package rebrandly

import (
	"context"
	"errors"
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

// LinksCount returns the number of links
func (c *Controller) LinksCount(filters *LinksCountFilters) (nbLinks int, err error) {
	return c.LinksCountCtx(nil, filters)
}

// LinksCountCtx returns the number of links
func (c *Controller) LinksCountCtx(ctx context.Context, filters *LinksCountFilters) (nbLinks int, err error) {
	query, err := convertStructToURLQuery(filters)
	if err != nil {
		err = fmt.Errorf("can't convert filters to query params: %v", err)
		return
	}
	var resp domainCountResponse
	url := *templateURL
	url.Path += "/links/count"
	url.RawQuery = query.Encode()
	if err = c.request(ctx, "GET", url, nil, &resp); err != nil {
		return
	}
	nbLinks = resp.Count
	return
}

// Links creation with GET won't supported as this is pure evil to create a ressource with GET

// LinksCreate creates a link
func (c *Controller) LinksCreate(payload LinkCreationPayload) (link Link, err error) {
	return c.LinksCreateCtx(nil, payload)
}

// LinksCreateCtx creates a link
func (c *Controller) LinksCreateCtx(ctx context.Context, payload LinkCreationPayload) (link Link, err error) {
	if payload.Destination == "" {
		err = errors.New("destination can't be empty")
		return
	}
	url := *templateURL
	url.Path += "/links"
	err = c.request(ctx, "POST", url, payload, &link)
	return
}
