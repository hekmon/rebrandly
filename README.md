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

## Enpoints

### Basic endpoints

* GET    account
* GET    links
* GET    links/:id
* GET    links/count
* GET    links/new (god no)
* POST   links
* POST   links/:id
* DELETE links/:id
* GET    domains
* GET    domains/:id
* GET    domains/count

### Advanced endpoints

not planned, feel free to fork or PR !
