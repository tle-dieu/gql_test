package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	graphql_go "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/tle-dieu/gql_test/graph"
	"github.com/tle-dieu/gql_test/pkg/db/mysql"
)

const defaultPort = "8080"

func parseSchema(path string, resolver interface{}) *graphql_go.Schema {
	bstr, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	schemaString := string(bstr)
	parsedSchema, err := graphql_go.ParseSchema(
		schemaString,
		resolver,
	)
	if err != nil {
		log.Fatal(err)
	}
	return parsedSchema
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	mysqlClient := mysql.NewMySQLClient()
	mysqlClient.Migrate()
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", &relay.Handler{
		Schema: parseSchema("./graph/schema.graphql", &graph.Resolver{Db: mysqlClient}),
	})
	port := "8080"
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}