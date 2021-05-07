package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mchmarny/restme/pkg/config"
	"github.com/mchmarny/restme/pkg/handler"
)

const (
	shutdownWaitSeconds  = 5
	serverTimeoutSeconds = 600
)

var (
	logger  = log.New(os.Stdout, "", 0)
	address = config.GetEnv("ADDRESS", ":8080")
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.DefaultHandler)
	mux.HandleFunc("/v1/load", handler.LoadHandler)
	mux.HandleFunc("/v1/resource", handler.ResourceHandler)
	mux.HandleFunc("/v1/request", handler.RequestHandler)

	// echo
	echoHandler := handler.NewEchoHandler(logger)
	mux.Handle("/v1/echo", echoHandler)

	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * serverTimeoutSeconds,
		WriteTimeout: time.Second * serverTimeoutSeconds,
		IdleTimeout:  time.Second * serverTimeoutSeconds,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()
	logger.Println("server started")

	<-done
	logger.Println("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownWaitSeconds*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown failed: %+v", err)
	}
}
