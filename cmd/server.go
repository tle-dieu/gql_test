package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tle-dieu/ad_graphql_api/config"
	"github.com/tle-dieu/ad_graphql_api/infrastructure/graph/generated"
	"github.com/tle-dieu/ad_graphql_api/infrastructure/graph/resolver"
	requester "github.com/tle-dieu/ad_graphql_api/infrastructure/http/client"
)

func NewHTTPClient(host string, port int) *requester.Client {
	httpClient := &http.Client{}
	httpRequester := requester.NewClient(fmt.Sprintf("http://%s:%d", host, port), httpClient)
	return httpRequester
}

func main() {
	// load config
	cfg := config.NewExtractAds()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolver.Resolver{
			HTTPClient: NewHTTPClient(cfg.HTTPClientHost, cfg.HTTPClientPort),
		}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", cfg.GraphqlPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.GraphqlPort), nil))
}
