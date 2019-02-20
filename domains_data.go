package rebrandly

import (
	"time"
)

// DomainsFilters represents the filters used to make a Domains query.
// https://developers.rebrandly.com/docs/domains-list-endpoint
type DomainsFilters struct {
	Active   *bool        `urlQuery:"active"`
	Type     *DomainsType `urlQuery:"type"`
	OrderBy  *OrderBy     `urlQuery:"orderBy"`
	OrderDir *OrderDir    `urlQuery:"orderDir"`
	Limit    *int         `urlQuery:"limit"`
	Last     *string      `urlQuery:"last"`
}

// DomainsType represent a domain type
type DomainsType string

const (
	// DomainsTypeUser represents the "user" type for a DomainsType
	DomainsTypeUser DomainsType = "user"
	// DomainsTypeService represents the "service" type for DomainsType
	DomainsTypeService DomainsType = "service"
)

// DomainsCountFilters represents the filer usable within a DomainsCount request
type DomainsCountFilters struct {
	Active *bool        `urlQuery:"active"`
	Type   *DomainsType `urlQuery:"active"`
}

// Domains represents a list domains
type Domains []Domain

// Domain represents a single custom domain
type Domain struct {
	ID             string    `json:"id"`
	FullName       string    `json:"fullName"`
	TopLevelDomain string    `json:"topLevelDomain"`
	Level          int       `json:"level"`
	CreationDate   time.Time `json:"creationDate"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	OwnerID        string    `json:"ownerId"`
	Type           string    `json:"type"`
	Subdomains     int       `json:"subdomains"`
	Managed        bool      `json:"managed"`
	Status         struct {
		DNS        string `json:"dns"`
		Encryption string `json:"encryption"`
	} `json:"status"`
	HTTPS       bool      `json:"https"`
	Active      bool      `json:"active"`
	Clicks      int       `json:"clicks"`
	Sessions    int       `json:"sessions"`
	LastClickAt time.Time `json:"lastClickAt"`
	Correlation struct {
		Status string `json:"status"`
	} `json:"correlation"`
}

type domainCountQuery struct {
	Active bool   `json:"active"`
	Type   string `json:"type"`
}

type domainCountResponse struct {
	Count int `json:"count"`
}
