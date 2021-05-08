package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
)

func NewEchoHandler(logger *log.Logger) EchoHandler {
	return EchoHandler{
		logger: logger,
	}
}

type EchoHandler struct {
	logger *log.Logger
}

type message struct {
	On      int64  `json:"on"`
	Message string `json:"msg"`
}

func (h EchoHandler) EchoHandler(c *gin.Context) {
	var m message
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Debugf("message: %v", m)

	c.IndentedJSON(http.StatusOK, m)
}
