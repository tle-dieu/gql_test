package resolver

import requester "github.com/tle-dieu/ad_graphql_api/infrastructure/http/client"

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	HttpClient *requester.Client
}
