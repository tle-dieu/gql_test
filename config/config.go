package config

import (
	"context"
	"fmt"

	"github.com/etf1/go-config"
	"github.com/etf1/go-config/env"
	"github.com/heetch/confita/backend/flags"
)

// configuration for extract-ads app
type ExtractAds struct {
	HTTPClientPort int    `config:"http_client_port"`
	HTTPClientHost string `config:"http_client_host"`
	GraphqlPort    int    `config:"graphql_port"`
}

// NewExtractAds creates a new ExtractAds configuration from env vars
func NewExtractAds() *ExtractAds {
	// create default config
	cfg := &ExtractAds{
		HTTPClientPort: 8080,
		HTTPClientHost: "localhost",
		GraphqlPort:    8081,
	}

	// load from .env and flags
	loader := config.NewConfigLoader(
		env.NewBackend(),
		flags.NewBackend(),
	)

	loader.LoadOrFatal(context.Background(), cfg)
	fmt.Println(config.TableString(cfg))

	return cfg
}
