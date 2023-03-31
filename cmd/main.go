package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mchmarny/restme/pkg/config"
	"github.com/mchmarny/restme/pkg/log"
)

const (
	appName = "rester"

	serverShutdownWaitSeconds = 5
	serverTimeoutSeconds      = 300
	serverMaxHeaderBytes      = 20
)

var (
	version    = "v0.0.1-default"
	address    = config.GetEnv("ADDRESS", "127.0.0.1:8080")
	configPath = config.GetEnv("CONFIG", "")
)

func main() {
	c, err := config.GetConfigFromFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("error getting app config from %s: %v", configPath, err))
	}

	logger := log.New(appName, version, c.Log.Level, c.Log.JSON)

	r, err := makeRouter(logger, c)
	if err != nil {
		panic(fmt.Sprintf("error creating router: %v", err))
	}

	s := &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    serverTimeoutSeconds * time.Second,
		WriteTimeout:   serverTimeoutSeconds * time.Second,
		MaxHeaderBytes: 1 << serverMaxHeaderBytes,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("server error: %s\n", err)
		}
	}()
	logger.Infof("app server started at %s", address)

	<-done
	logger.Info("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownWaitSeconds*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Fatalf("server shutdown failed: %+v", err)
	}
}
