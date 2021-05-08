package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Request represents simple HTTP resource
type Request struct {
	Request *RequestMetadata       `json:"request,omitempty"`
	Headers map[string]interface{} `json:"headers"`
	EnvVars map[string]interface{} `json:"env_vars"`
}

// RequestMetadata represents metadata of the request
type RequestMetadata struct {
	ID       string    `json:"id"`
	On       time.Time `json:"time"`
	Path     string    `json:"path"`
	Protocol string    `json:"protocol"`
	Host     string    `json:"host"`
	Method   string    `json:"method"`
}

func RequestHandler(c *gin.Context) {
	result := &Request{
		Request: getRequestMetadata(c),
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

func getRequestMetadata(c *gin.Context) *RequestMetadata {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error while getting id: %v\n", err)
	}

	return &RequestMetadata{
		ID:       id.String(),
		On:       time.Now().UTC(),
		Path:     c.FullPath(),
		Protocol: c.Request.Proto,
		Host:     c.Request.Host,
		Method:   c.Request.Method,
	}
}
