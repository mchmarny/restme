package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mchmarny/restme/pkg/config"
	"github.com/mchmarny/restme/pkg/handler"
	"github.com/mchmarny/restme/pkg/log"
)

const (
	appName = "rester"

	serverShutdownWaitSeconds = 5
	serverTimeoutSeconds      = 300
	serverMaxHeaderBytes      = 20
)

var (
	version = "v0.0.1-default"
	address = config.GetEnv("ADDRESS", ":8080")
)

func main() {
	logger := log.New(appName)

	s := &http.Server{
		Addr:           address,
		Handler:        handler.SetupRouter(appName, version, logger),
		ReadTimeout:    serverTimeoutSeconds * time.Second,
		WriteTimeout:   serverTimeoutSeconds * time.Second,
		MaxHeaderBytes: 1 << serverMaxHeaderBytes,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("error: %s\n", err)
		}
	}()
	logger.Info("server started")

	<-done
	logger.Info("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownWaitSeconds*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown failed: %+v", err)
	}
}
