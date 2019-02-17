package rebrandly

import "time"

// Domains represents a list domains
type Domains []Domain

// Domain represents a single custom domain
type Domain struct {
	ID             string    `json:"id"`
	FullName       string    `json:"fullName"`
	TopLevelDomain string    `json:"topLevelDomain"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Type           string    `json:"type"`
	Active         bool      `json:"active"`
}

type domainCountQuery struct {
	Active bool   `json:"active"`
	Type   string `json:"type"`
}

type domainCountResponse struct {
	Count int `json:"count"`
}
