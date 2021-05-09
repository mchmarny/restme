package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	On      int64  `json:"on"`
	Message string `json:"msg"`
}

func (h *Handler) EchoHandler(c *gin.Context) {
	var m message
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Debugf("message: %v", m)

	c.IndentedJSON(http.StatusOK, m)
}
