package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gol4ng/logger"
)

// Server is the struct for the HTTP server.
type Server struct {
	httpServer *http.Server
	logger     logger.LoggerInterface
}

// ListenAndServe to handle requests
func (s *Server) ListenAndServe() {
	s.logger.Info("starting http server")

	err := s.httpServer.ListenAndServe()
	if err != nil {
		s.logger.Error("failed to start http server : %err%", logger.Error("err", err))
	}
}

// Shutdown stops the HTTP server
func (s *Server) Shutdown() error {
	if s.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return s.httpServer.Shutdown(ctx)
	}
	return nil
}

// NewServer method initializes a new HTTP server instance
func NewServer(addr string, handler http.Handler, l logger.LoggerInterface) *Server {
	return &Server{
		httpServer: &http.Server{Addr: addr, Handler: handler},
		logger:     l,
	}
}
