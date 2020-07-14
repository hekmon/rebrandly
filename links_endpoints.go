package rebrandly

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// LinksGet returns the list of links
func (c *Controller) LinksGet(filters *LinksFilters) (links Links, err error) {
	return c.LinksGetCtx(context.Background(), filters)
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
	err = c.request(ctx, "GET", url, nil, &links, []int{http.StatusNotFound})
	return
}

// LinksGetAll returns the list of links
func (c *Controller) LinksGetAll(filters *LinksFilters) (links Links, err error) {
	return c.LinksGetAllCtx(context.Background(), filters)
}

// LinksGetAllCtx recursively retreive all the links (not limited by max)
func (c *Controller) LinksGetAllCtx(ctx context.Context, filters *LinksFilters) (links Links, err error) {
	query, err := convertStructToURLQuery(filters)
	if err != nil {
		err = fmt.Errorf("can't convert filters to query params: %v", err)
		return
	}
	url := *templateURL
	url.Path += "/links"
	// Loop thru all pages
	var page int
	for {
		page++
		if page != 1 {
			query.Set("last", links[len(links)-1].ID)
		}
		url.RawQuery = query.Encode()
		var pageLinks Links
		if err = c.request(ctx, "GET", url, nil, &pageLinks, []int{http.StatusNotFound}); err != nil {
			err = fmt.Errorf("error at page %d: %w", page, err)
			return
		}
		links = append(links, pageLinks...)
		if len(pageLinks) == 0 {
			break
		}
	}
	return
}

// LinksGetByID returns the link details of link id.
func (c *Controller) LinksGetByID(id string) (link Link, err error) {
	return c.LinksGetByIDCtx(context.Background(), id)
}

// LinksGetByIDCtx returns the link details of link id.
func (c *Controller) LinksGetByIDCtx(ctx context.Context, id string) (link Link, err error) {
	url := *templateURL
	url.Path += fmt.Sprintf("/links/%s", id)
	err = c.request(ctx, "GET", url, nil, &link, []int{http.StatusNotFound})
	return
}

// LinksCount returns the number of links
func (c *Controller) LinksCount(filters *LinksCountFilters) (nbLinks int, err error) {
	return c.LinksCountCtx(context.Background(), filters)
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
	if err = c.request(ctx, "GET", url, nil, &resp, nil); err != nil {
		return
	}
	nbLinks = resp.Count
	return
}

// Links creation with GET won't supported as this is pure evil to create a ressource with GET

// LinksCreate creates a link
func (c *Controller) LinksCreate(payload LinkCreationPayload) (link Link, err error) {
	return c.LinksCreateCtx(context.Background(), payload)
}

// LinksCreateCtx creates a link
func (c *Controller) LinksCreateCtx(ctx context.Context, payload LinkCreationPayload) (link Link, err error) {
	if payload.Destination == "" {
		err = errors.New("destination can't be empty")
		return
	}
	url := *templateURL
	url.Path += "/links"
	err = c.request(ctx, "POST", url, payload, &link, []int{http.StatusForbidden, http.StatusNotFound})
	return
}

// LinksUpdate allows to update a link title and destination by its id.
func (c *Controller) LinksUpdate(id string, payload LinkUpdatePayload) (link Link, err error) {
	return c.LinksUpdateCtx(context.Background(), id, payload)
}

// LinksUpdateCtx allows to update a link title and destination by its id.
func (c *Controller) LinksUpdateCtx(ctx context.Context, id string, payload LinkUpdatePayload) (link Link, err error) {
	if id == "" {
		err = errors.New("id can't be empty")
		return
	}
	if payload.Title == "" {
		err = errors.New("title can't be empty")
		return
	}
	if payload.Destination == "" {
		err = errors.New("destination can't be empty")
		return
	}
	url := *templateURL
	url.Path += fmt.Sprintf("/links/%s", id)
	err = c.request(ctx, "POST", url, payload, &link, []int{http.StatusForbidden, http.StatusNotFound})
	return
}

// LinksDelete deletes a link identified by id.
func (c *Controller) LinksDelete(id string) (link Link, err error) {
	return c.LinksDeleteCtx(context.Background(), id)
}

// LinksDeleteCtx deletes a link identified by id.
func (c *Controller) LinksDeleteCtx(ctx context.Context, id string) (link Link, err error) {
	if id == "" {
		err = errors.New("id can't be empty")
		return
	}
	url := *templateURL
	url.Path += fmt.Sprintf("/links/%s", id)
	err = c.request(ctx, "DELETE", url, nil, &link, []int{http.StatusNotFound})
	return
}
