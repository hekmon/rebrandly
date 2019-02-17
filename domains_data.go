package rebrandly

import "time"

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
