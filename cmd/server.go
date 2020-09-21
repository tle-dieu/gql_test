package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tle-dieu/ad_graphql_api/infrastructure/graph/generated"
	"github.com/tle-dieu/ad_graphql_api/infrastructure/graph/resolver"
	requester "github.com/tle-dieu/ad_graphql_api/infrastructure/http/client"
)

const defaultPort = "8081"
const apiHost = "127.0.0.1"
const apiPort = "8080"

func NewHTTPClient() *requester.Client {
	httpClient := &http.Client{}
	httpRequester := requester.NewClient(fmt.Sprintf("http://%s:%s", apiHost, apiPort), httpClient)
	return httpRequester
}

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolver.Resolver{
			HttpClient: NewHTTPClient(),
		}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
