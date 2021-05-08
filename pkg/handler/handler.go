package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
)

func DefaultHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		c.FullPath(): "not implemented",
	})
}

func SetupRouter(logger *log.Logger) *gin.Engine {
	gin.ForceConsoleColor()
	r := gin.Default()

	r.Any("/", DefaultHandler)

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
