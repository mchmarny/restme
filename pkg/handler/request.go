package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Request represents simple HTTP resource
type Request struct {
	Request gin.H                  `json:"request,omitempty"`
	Headers map[string]interface{} `json:"headers"`
	EnvVars map[string]interface{} `json:"env_vars"`
}

func (h *Handler) RequestHandler(c *gin.Context) {
	result := &Request{
		Request: h.getRequestMetadata(c),
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

func (h *Handler) getRequestMetadata(c *gin.Context) gin.H {
	id, err := uuid.NewUUID()
	if err != nil {
		h.logger.Errorf("Error while getting id: %v\n", err)
	}

	return gin.H{
		"id":       id.String(),
		"time":     time.Now().UTC(),
		"version":  h.logger.Version,
		"path":     c.FullPath(),
		"protocol": c.Request.Proto,
		"host":     c.Request.Host,
		"method":   c.Request.Method,
	}
}
