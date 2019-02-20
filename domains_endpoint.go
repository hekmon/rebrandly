package rebrandly

import (
	"context"
	"fmt"
)

// DomainsGet returns the list of domains
func (c *Controller) DomainsGet(filters *DomainsQuery) (domains Domains, err error) {
	return c.DomainsGetCtx(nil, filters)
}

// DomainsGetCtx returns the list of domains
func (c *Controller) DomainsGetCtx(ctx context.Context, filters *DomainsQuery) (domains Domains, err error) {
	query, err := convertStructToURLQuery(filters)
	if err != nil {
		err = fmt.Errorf("can't convert filters to query params: %v", err)
		return
	}
	url := *templateURL
	url.Path += "/domains"
	url.RawQuery = query.Encode()
	err = c.request(ctx, "GET", url, nil, &domains)
	return
}

// DomainsGetByID return the domains details represented by id
func (c *Controller) DomainsGetByID(id string) (domain Domain, err error) {
	return c.DomainsGetByIDCtx(nil, id)
}

// DomainsGetByIDCtx return the domains details represented by id
func (c *Controller) DomainsGetByIDCtx(ctx context.Context, id string) (domain Domain, err error) {
	url := *templateURL
	url.Path += fmt.Sprintf("/domains/%s", id)
	err = c.request(ctx, "GET", url, nil, &domain)
	return
}

// DomainsCount returns the number of domains
func (c *Controller) DomainsCount(filters *DomainsCountQuery) (nbDomains int, err error) {
	return c.DomainsCountCtx(nil, filters)
}

// DomainsCountCtx returns the number of domains
func (c *Controller) DomainsCountCtx(ctx context.Context, filters *DomainsCountQuery) (nbDomains int, err error) {
	query, err := convertStructToURLQuery(filters)
	if err != nil {
		err = fmt.Errorf("can't convert filters to query params: %v", err)
		return
	}
	var resp domainCountResponse
	url := *templateURL
	url.Path += "/domains/count"
	url.RawQuery = query.Encode()
	if err = c.request(ctx, "GET", url, nil, &resp); err != nil {
		return
	}
	nbDomains = resp.Count
	return
}
