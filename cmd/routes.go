package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/config"
	"github.com/mchmarny/restme/pkg/echo"
	"github.com/mchmarny/restme/pkg/log"
	"github.com/mchmarny/restme/pkg/request"
	"github.com/pkg/errors"
)

func makeRouter(logger *log.Logger, conf *config.Config) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(options)

	v1 := r.Group("/api/v1")
	v1.Use(APITokenRequired(conf.Auth.Tokens)) // all API calls require token

	echoGroup := v1.Group("/echo")
	echoService := echo.NewService(logger)
	echoGroup.POST("/message", echoService.MessageHandler)

	reqGroup := v1.Group(("/request"))
	reqService, err := request.NewService(logger, conf)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create request service")
	}
	reqGroup.GET("/info", reqService.RequestHandler)

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

	return r, nil
}
