package config

import (
	"context"
	"fmt"

	"github.com/etf1/go-config"
	"github.com/etf1/go-config/env"
	"github.com/etf1/go-config/prefix"
)

// Config of HTTP server.
type Config struct {
	HTTPPort string `config:"HTTP_PORT"`

	MetricsConfig
}

// MetricsConfig defines the number of top used request values that we would monitor
type MetricsConfig struct {
	TopUsedRequest int `config:"TOP_USED_REQUEST"`
}

// NewConfig method initializes a server's config
func NewConfig(ctx context.Context) *Config {
	cfg := &Config{
		HTTPPort: ":8001",
		MetricsConfig: MetricsConfig{
			TopUsedRequest: 5,
		},
	}

	loader := config.NewDefaultConfigLoader().AppendBackends(
		prefix.NewBackend("", env.NewBackend()),
	)

	loader.LoadOrFatal(ctx, cfg)
	fmt.Println(config.TableString(cfg))

	return cfg
}
