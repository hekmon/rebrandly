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
	err = c.request(ctx, "GET", "domains", filters, &domains)
	return
}

// DomainsGetByID return the domains details represented by id
func (c *Controller) DomainsGetByID(id string) (domain Domain, err error) {
	return c.DomainsGetByIDCtx(nil, id)
}

// DomainsGetByIDCtx return the domains details represented by id
func (c *Controller) DomainsGetByIDCtx(ctx context.Context, id string) (domain Domain, err error) {
	err = c.request(ctx, "GET", fmt.Sprintf("domains/%s", id), nil, &domain)
	return
}

// DomainsCount returns the number of domains
func (c *Controller) DomainsCount(active bool, domainType string) (nbDomains int, err error) {
	return c.DomainsCountCtx(nil, active, domainType)
}

// DomainsCountCtx returns the number of domains
func (c *Controller) DomainsCountCtx(ctx context.Context, active bool, domainType string) (nbDomains int, err error) {
	var resp domainCountResponse
	if err = c.request(ctx, "GET", "domains/count", domainCountQuery{
		Active: active,
		Type:   domainType,
	}, &resp); err != nil {
		return
	}
	nbDomains = resp.Count
	return
}
