package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
)

func NewHandler(logger *log.Logger) *Handler {
	gin.ForceConsoleColor()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(Options)

	h := &Handler{
		logger: logger,
		Engine: r,
	}

	r.GET("/", h.DefaultHandler)

	v1 := r.Group("/v1")
	v1.POST("/echo", h.EchoHandler)
	v1.GET("/load/:duration", h.LoadHandler)
	v1.GET("/resource", h.ResourceHandler)
	v1.GET("/request", h.RequestHandler)

	return h
}

type Handler struct {
	logger *log.Logger
	Engine *gin.Engine
}

func (h *Handler) DefaultHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"request": h.getRequestMetadata(c),
		"routes": []string{
			"POST /v1/echo",
			"GET /v1/load/:duration (e.g. 5s)",
			"GET /v1/request",
			"GET /v1/resource",
		},
	})
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
