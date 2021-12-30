package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mchmarny/restme/pkg/config"
	"github.com/mchmarny/restme/pkg/echo"
	"github.com/mchmarny/restme/pkg/log"
	"github.com/mchmarny/restme/pkg/request"

	"github.com/gin-gonic/gin"
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

func makeRouter(logger *log.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(options)

	v1 := r.Group("/v1")
	{
		echoGroup := v1.Group("/echo")
		{
			echoService := echo.NewService(logger)
			echoGroup.POST("/message", echoService.MessageHandler)
		}

		reqGroup := v1.Group(("/request"))
		{
			reqService := request.NewService(logger)
			reqGroup.GET("/info", reqService.RequestHandler)
		}
	}

	// collect routes for index
	routes := []string{}
	routeInfo := r.Routes()
	for _, info := range routeInfo {
		routes = append(routes, fmt.Sprintf("%-7s %s", info.Method, info.Path))
	}

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"routes": routes,
		})
	})

	return r
}

// options midleware adds options headers.
func options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "POST,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {
	logger := log.New(appName, version)

	r := makeRouter(logger)

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

	if err := s.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown failed: %+v", err)
	}
}
