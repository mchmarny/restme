package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
)

func DefaultHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"request": getRequestMetadata(c),
		"routes": []string{
			"POST /v1/echo",
			"GET /v1/load/:duration (e.g. 5s)",
			"GET /v1/request",
			"GET /v1/resource",
		},
	})
}

func SetupRouter(name, version string, logger *log.Logger) *gin.Engine {
	gin.ForceConsoleColor()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(Options)

	r.GET("/", DefaultHandler)

	v1 := r.Group("/v1")
	{
		echoHandler := NewEchoHandler(logger)
		v1.POST("/echo", echoHandler.EchoHandler)

		v1.GET("/load/:duration", LoadHandler)
		v1.GET("/resource", ResourceHandler)
		v1.GET("/request", RequestHandler)
	}
	return r
}

// Options midleware adds options headers.
func Options(c *gin.Context) {
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
