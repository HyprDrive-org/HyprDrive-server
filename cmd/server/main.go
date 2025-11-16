package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HyprDrive/HyprDrive-server/internal/config"
	"github.com/HyprDrive/HyprDrive-server/internal/httpserver"
	"github.com/HyprDrive/HyprDrive-server/internal/logger"
)

func main() {
	log := logger.New()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	srv := httpserver.New(cfg, log)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("server stopped: %v", err)
		}
	}()

	shutdownCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-shutdownCtx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
		os.Exit(1)
	}
}
