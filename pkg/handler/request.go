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

const (
	requestHandlerPath = "/v1/request"
)

// Request represents simple HTTP resource
type Request struct {
	Meta    *RequestMetadata       `json:"meta,omitempty"`
	Headers map[string]interface{} `json:"head"`
	EnvVars map[string]interface{} `json:"envs"`
}

// RequestMetadata represents metadata of the request
type RequestMetadata struct {
	ID     string    `json:"id"`
	On     time.Time `json:"ts"`
	URI    string    `json:"uri"`
	Host   string    `json:"host"`
	Method string    `json:"method"`
}

func NewRequestHandler(logger *log.Logger) RequestHandler {
	return RequestHandler{
		logger: logger,
		Path:   requestHandlerPath,
	}
}

type RequestHandler struct {
	logger *log.Logger
	Path   string
}

func (h RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("serving: %+v", r)

	if r.URL.Path != h.Path {
		handleError(w, http.StatusNotFound, "Expected: %s, got:%s", h.Path, r.URL.Path)
		return
	}

	result := &Request{
		Meta:    getRequestMetadata(r),
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

	w.WriteHeader(http.StatusOK)
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
