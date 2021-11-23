package service

import (
	"net/http"

	httpSrv "github.com/fizzbuzz-api/internal/server/http"
	"github.com/fizzbuzz-api/internal/server/http/handler"

	"github.com/gorilla/mux"
)

// GetHTTPServer return HTTP server.
func (container *Container) GetHTTPServer() *httpSrv.Server {
	return httpSrv.NewServer(
		container.Config().HTTPPort,
		container.getHTTPHandler(),
		container.GetLogger(),
	)
}

func (container *Container) getHTTPHandler() http.Handler {
	r := mux.NewRouter()
	r.Use(container.MonitoringMiddleware)
	r.Path("/fizzbuzz-api/fizzbuzz/").Methods("GET").HandlerFunc(handler.MyFizzBuzz(container.GetLogger()))
	r.Path("/fizzbuzz-api/fizzbuzz/metrics").HandlerFunc(handler.PrintMetrics(container.GetLogger(), container.GetMetric()))

	return r
}
