package echo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/httputil"
	"github.com/mchmarny/restme/pkg/log"
	"github.com/pkg/errors"
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

// MessageHandler responds with the inbound message.
func (s *Service) MessageHandler(c *gin.Context) {
	var m Message
	if err := c.ShouldBindJSON(&m); err != nil {
		httputil.NewError(c, http.StatusBadRequest, errors.Wrap(err, "Invalid message format"))
		return
	}
	s.logger.Debugf("message: %v", m)
	c.IndentedJSON(http.StatusOK, m)
}
