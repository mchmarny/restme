package request

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mchmarny/restme/pkg/log"
)

// NewRequestService creates new RequestService instance.
func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

// Service provides object representing the inbound HTTP request.
type Service struct {
	logger *log.Logger
}

// RequestHandler handles the inbound requests.
func (s *Service) RequestHandler(c *gin.Context) {
	result := &Response{
		Request: s.getRequestMetadata(c),
		Headers: make(map[string]interface{}),
		EnvVars: make(map[string]interface{}),
	}

	// env vars
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		result.EnvVars[pair[0]] = pair[1]
	}

	// headers
	for name, headers := range c.Request.Header {
		name = strings.ToLower(name)
		for i, h := range headers {
			if len(headers) > 1 {
				result.Headers[fmt.Sprintf("%s[%d]", name, i)] = h
			} else {
				result.Headers[name] = h
			}
		}
	}

	c.IndentedJSON(http.StatusOK, result)
}

func (s *Service) getRequestMetadata(c *gin.Context) gin.H {
	id, err := uuid.NewUUID()
	if err != nil {
		s.logger.Errorf("Error while getting id: %v\n", err)
	}

	return gin.H{
		"id":       id.String(),
		"time":     time.Now().UTC(),
		"version":  s.logger.GetAppVersion(),
		"path":     c.Request.URL.EscapedPath(),
		"protocol": c.Request.Proto,
		"host":     c.Request.Host,
		"method":   c.Request.Method,
	}
}
