package service

import (
	"sync"

	"github.com/fizzbuzz-api/config"
	"github.com/fizzbuzz-api/internal/metric"

	"github.com/gol4ng/logger"
)

// Container struct
type Container struct {
	cfg *config.Config

	logger     logger.LoggerInterface
	loggerOnce sync.Once

	metric *metric.Metric
}

// NewContainer initializes a new container instance.
func NewContainer(cfg *config.Config) *Container {
	return &Container{
		cfg: cfg,
	}
}

// Config returns server's config
func (c *Container) Config() *config.Config {
	return c.cfg
}
