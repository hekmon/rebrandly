package rebrandly

import (
	"time"
)

// DomainsQuery allows to configure a Domains query.
// https://developers.rebrandly.com/docs/domains-list-endpoint
type DomainsQuery struct {
	Active   *bool                 `urlQuery:"active"`
	Type     *DomainsType          `urlQuery:"type"`
	OrderBy  *DomainsQueryOrderBy  `urlQuery:"orderBy"`
	OrderDir *DomainsQueryOrderDir `urlQuery:"orderDir"`
	Limit    *int                  `urlQuery:"limit"`
	Last     *string               `urlQuery:"last"`
}

// DomainsType represent a domain type
type DomainsType string

const (
	// DomainsQueryTypeUser represents the "user" type for a DomainsType
	DomainsQueryTypeUser DomainsType = "user"
	// DomainsQueryTypeService represents the "service" type for DomainsType
	DomainsQueryTypeService DomainsType = "service"
)

// DomainsQueryOrderBy represent a given ordering for a DomainsQuery.
type DomainsQueryOrderBy string

const (
	// DomainsQueryOrderByCreatedAt represents the "createdAt" ordering for a DomainsQueryOrderBy
	DomainsQueryOrderByCreatedAt DomainsQueryOrderBy = "createdAt"
	// DomainsQueryOrderByUpdatedAt represents the "updatedAt" ordering for a DomainsQueryOrderBy
	DomainsQueryOrderByUpdatedAt DomainsQueryOrderBy = "updatedAt"
	// DomainsQueryOrderByFullName represents the "fullName" ordering for a DomainsQueryOrderBy
	DomainsQueryOrderByFullName DomainsQueryOrderBy = "fullName"
)

// DomainsQueryOrderDir represents the sorting direction for a DomainsQuery
type DomainsQueryOrderDir string

const (
	// DomainsQueryOrderDirDesc represents the "desc" sorting direction for a DomainsQueryOrderDir
	DomainsQueryOrderDirDesc DomainsQueryOrderDir = "desc"
	// DomainsQueryOrderDirAsc represents the "asc" sorting direction for a DomainsQueryOrderDir
	DomainsQueryOrderDirAsc DomainsQueryOrderDir = "asc"
)

// DomainsCountQuery represents the filer usable within a DomainsCount request
type DomainsCountQuery struct {
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
