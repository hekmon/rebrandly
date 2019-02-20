# Rebrandly

[![GoDoc](https://godoc.org/github.com/hekmon/rebrandly?status.svg)](https://godoc.org/github.com/hekmon/rebrandly) [![Go Report Card](https://goreportcard.com/badge/github.com/hekmon/rebrandly)](https://goreportcard.com/report/github.com/hekmon/rebrandly)

Rebrandly API golang bindings.

## Quickstart

```golang
// Init
rebapi := rebrandly.New(apiKey)
// rebapi.SetWorkspace("workspace")
// rebapi.SetUserAgent("whatever")

// Create
link, err := rebapi.LinksCreate(rebrandly.LinkCreationPayload{
    Destination: "https://github.com/hekmon/rebrandly",
    SlashTag:    "rebrandly-golang",
    Title:       "Rebrandly Golang bindings library",
    // Domain:      &domain,
})
```

Each endpoint supports golang contextes, just prefix Ctx to it: `rebapi.LinksCreateCtx(ctx, rebrandly.LinkCreationPayload{...`

## Enpoints

### Basic endpoints

* GET    account - [binding doc](https://godoc.org/github.com/hekmon/rebrandly#Controller.Account) | [api doc](https://developers.rebrandly.com/docs/account-details-endpoint)
* GET    links - [binding doc](https://godoc.org/github.com/hekmon/rebrandly#Controller.LinksGet) | [api doc](https://developers.rebrandly.com/docs/links-list-endpoint)
* GET    links/:id - [binding doc](https://godoc.org/github.com/hekmon/rebrandly#Controller.LinksGetByID) | [api doc](https://developers.rebrandly.com/docs/links-detail-endpoint)
* GET    links/count - [binding doc](https://godoc.org/github.com/hekmon/rebrandly#Controller.LinksCount) | [api doc](https://developers.rebrandly.com/docs/counting-links-endpoint)
* GET    links/new (god no)
* POST   links - [binding doc](https://godoc.org/github.com/hekmon/rebrandly#Controller.LinksCreate) | [api doc](https://developers.rebrandly.com/docs/create-link-endpoint)
* POST   links/:id - [binding doc](https://godoc.org/github.com/hekmon/rebrandly#Controller.LinksUpdate) | [api doc](https://developers.rebrandly.com/docs/update-link-endpoint)
* DELETE links/:id - [binding doc](https://godoc.org/github.com/hekmon/rebrandly#Controller.LinksDelete) | [api doc](https://developers.rebrandly.com/docs/delete-link-endpoint)
* GET    domains - [binding doc](https://godoc.org/github.com/hekmon/rebrandly#Controller.DomainsGet) | [api doc](https://developers.rebrandly.com/docs/domains-list-endpoint)
* GET    domains/:id - [binding doc](https://godoc.org/github.com/hekmon/rebrandly#Controller.DomainsGetByID) | [api doc](https://developers.rebrandly.com/docs/domain-details-endpoint)
* GET    domains/count - [binding doc](https://godoc.org/github.com/hekmon/rebrandly#Controller.DomainsCount) | [api doc](https://developers.rebrandly.com/docs/count-domains-endpoint)

### Advanced endpoints

not planned, feel free to fork or PR !
