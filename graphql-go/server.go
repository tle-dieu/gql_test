package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	graphql_go "github.com/tle-dieu/gql_test/graphql-go"
	"github.com/tle-dieu/gql_test/pkg/db/mysql"
)

const defaultPort = "8080"

func parseSchema(path string, resolver interface{}) *graphql.Schema {
	bstr, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	schemaString := string(bstr)
	parsedSchema, err := graphql.ParseSchema(
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
		Schema: parseSchema("./graphql-go/schema.graphqls", &graphql_go.AdResolver{Db: mysqlClient}),
	})
	port := "8080"
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
