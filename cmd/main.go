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
	shutdownWaitSeconds = 5
)

var (
	logger  = log.New(os.Stdout, "", 0)
	address = config.GetEnv("ADDRESS", ":8080")
)

func main() {
	mux := http.NewServeMux()

	// default
	mux.HandleFunc("/", handler.DefaultHandler)

	// echo
	echoHandler := handler.NewEchoHandler(logger)
	mux.Handle(echoHandler.Path, echoHandler)

	// resource
	resourceHandler := handler.NewResourceHandler(logger)
	mux.Handle(resourceHandler.Path, resourceHandler)

	// request
	requestHandler := handler.NewRequestHandler(logger)
	mux.Handle(requestHandler.Path, requestHandler)

	s := &http.Server{
		Addr:     address,
		Handler:  mux,
		ErrorLog: logger,
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
