package rebrandly

import "time"

// LinksFilters represents all filters usable on Links queries
type LinksFilters struct {
	DomainID       *string   `urlQuery:"domain.id"`
	DomainFullName *string   `urlQuery:"domain.fullName"`
	SlashTag       *string   `urlQuery:"slashtag"`
	CreatorID      *string   `urlQuery:"creator.id"`
	OrderBy        *OrderBy  `urlQuery:"orderBy"`
	OrderDir       *OrderDir `urlQuery:"orderDir"`
	Limit          *int      `urlQuery:"limit"`
	Last           *string   `urlQuery:"last"`
}

// Link represents a registered rebrandly link
type Link struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Slashtag      string    `json:"slashtag"`
	Destination   string    `json:"destination"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Status        string    `json:"status"`
	Clicks        int       `json:"clicks"`
	Sessions      int       `json:"sessions"`
	LastClickDate time.Time `json:"lastClickDate"`
	LastClickAt   time.Time `json:"lastClickAt"`
	IsPublic      bool      `json:"isPublic"`
	ShortURL      string    `json:"shortUrl"`
	DomainID      string    `json:"domainId"`
	DomainName    string    `json:"domainName"`
	Domain        struct {
		ID       string `json:"id"`
		Ref      string `json:"ref"`
		FullName string `json:"fullName"`
		Active   bool   `json:"active"`
	} `json:"domain"`
	HTTPS     bool `json:"https"`
	Favourite bool `json:"favourite"`
	Creator   struct {
		ID        string `json:"id"`
		FullName  string `json:"fullName"`
		AvatarURL string `json:"avatarUrl"`
	} `json:"creator"`
	Integrated bool `json:"integrated"`
}

// Links represents a collection of Link
type Links []Link
