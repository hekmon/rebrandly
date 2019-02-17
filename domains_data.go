package rebrandly

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// DomainsQuery allows to configure a Domains query
type DomainsQuery struct {
	Active   *bool   `json:"active"`
	Type     *string `json:"type"`
	OrderBy  *string `json:"orderBy"`
	OrderDir *string `json:"orderDir"`
	Limit    *int    `json:"limit"`
	Last     *string `json:"last"`
}

// MarshalJSON only marshal as JSON instanciate fields of DomainsQuery
func (dq *DomainsQuery) MarshalJSON() (data []byte, err error) {
	refType := reflect.TypeOf(*dq)
	refValue := reflect.ValueOf(*dq)
	tmp := make(map[string]interface{}, refType.NumField())
	var (
		field reflect.Value
		key   string
	)
	for i := 0; i < refType.NumField(); i++ {
		field = refValue.Field(i)
		if field.IsNil() {
			continue
		}
		key = refType.Field(i).Tag.Get("json")
		switch typedValue := field.Elem().Interface().(type) {
		case bool:
			tmp[key] = typedValue
		case string:
			tmp[key] = typedValue
		case int:
			tmp[key] = typedValue
		default:
			err = fmt.Errorf("elem id %d with json key '%s' is not supported: %v", i, key, reflect.TypeOf(typedValue))
		}
	}
	return json.Marshal(tmp)
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
