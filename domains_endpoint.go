package rebrandly

import "context"

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
