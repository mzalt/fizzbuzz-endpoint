package service

import (
	"os"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/middleware"
)

// GetLogger intializes once and returns logger
func (container *Container) GetLogger() logger.LoggerInterface {
	container.loggerOnce.Do(func() {
		container.logger = logger.NewLogger(container.getLoggerHandler())
	})

	return container.logger
}

func (container *Container) getLoggerHandler() logger.HandlerInterface {
	logFormatter := formatter.NewDefaultFormatter(formatter.WithColor(true), formatter.WithContext(true))
	h := handler.Stream(os.Stdout, logFormatter)

	return container.getLoggerHandlerMiddleware().Decorate(h)
}

func (container *Container) getLoggerHandlerMiddleware() logger.Middlewares {
	return logger.MiddlewareStack(
		middleware.Placeholder(),
		middleware.MinLevelFilter(logger.InfoLevel),
	)
}
