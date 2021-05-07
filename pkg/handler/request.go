package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

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
	ID     string    `json:"id"`
	On     time.Time `json:"time"`
	URI    string    `json:"uri"`
	Host   string    `json:"host"`
	Method string    `json:"method"`
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	result := &Request{
		Request: getRequestMetadata(r),
		Headers: make(map[string]interface{}),
		EnvVars: make(map[string]interface{}),
	}

	// env vars
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		result.EnvVars[pair[0]] = pair[1]
	}

	// headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for i, h := range headers {
			if len(headers) > 1 {
				result.Headers[fmt.Sprintf("%s[%d]", name, i)] = h
			} else {
				result.Headers[name] = h
			}
		}
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(result); err != nil {
		handleError(w, http.StatusInternalServerError, "Error processing request: %v", err)
		return
	}
}

func getRequestMetadata(r *http.Request) *RequestMetadata {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error while getting id: %v\n", err)
	}

	return &RequestMetadata{
		ID:     id.String(),
		On:     time.Now(),
		URI:    r.RequestURI,
		Host:   r.Host,
		Method: r.Method,
	}
}
