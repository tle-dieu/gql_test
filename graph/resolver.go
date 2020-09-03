package graph

//go:generate ./gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "github.com/tle-dieu/gql_test/pkg/db/mysql"

type Resolver struct {
	Db *mysql.ClientMySQL
}
