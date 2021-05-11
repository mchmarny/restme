package echo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
)

// NewService creates a new EchoService instance.
func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

// Service provides message echo service.
type Service struct {
	logger *log.Logger
}

// MessageHandler handles the inbound messages.
func (s *Service) MessageHandler(c *gin.Context) {
	var m Message
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.logger.Debugf("message: %v", m)

	c.IndentedJSON(http.StatusOK, m)
}
