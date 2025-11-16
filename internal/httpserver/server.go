package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/HyprDrive/HyprDrive-server/internal/config"
	"github.com/HyprDrive/HyprDrive-server/internal/logger"
)

// Server describes the API server with graceful shutdown capabilities.
type Server struct {
	cfg    *config.Config
	logger *logger.Logger
	http   *http.Server
}

// New creates a new HTTP server with default routes.
func New(cfg *config.Config, log *logger.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.HTTPPort),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &Server{cfg: cfg, logger: log, http: srv}
}

// Start begins listening for incoming HTTP traffic.
func (s *Server) Start() error {
	s.logger.Printf("starting http server on %s", s.http.Addr)
	return s.http.ListenAndServe()
}

// Shutdown gracefully stops the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Println("shutting down http server")
	return s.http.Shutdown(ctx)
}
