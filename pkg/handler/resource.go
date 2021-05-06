package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mchmarny/restme/pkg/kube"
)

const (
	resourceHandlerPath = "/v1/resource"
)

func NewResourceHandler(logger *log.Logger) ResourceHandler {
	return ResourceHandler{
		logger: logger,
		Path:   resourceHandlerPath,
	}
}

type ResourceHandler struct {
	logger *log.Logger
	Path   string
}

// Resource represents simple Kube resource
type Resource struct {
	Request   *RequestMetadata   `json:"request,omitempty"`
	Host      *kube.HostInfo     `json:"host,omitempty"`
	Resources *kube.ResourceInfo `json:"resources,omitempty"`
	Limits    *kube.ResourceInfo `json:"limits,omitempty"`
}

func (h ResourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("serving: %+v", r)

	if r.URL.Path != h.Path {
		handleError(w, http.StatusNotFound, "Expected: %s, got:%s", h.Path, r.URL.Path)
		return
	}

	result := &Resource{
		Request:   getRequestMetadata(r),
		Host:      kube.GetHostInfo(),
		Resources: kube.GetResourceInfo(),
		Limits:    kube.GetLimits(),
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(result); err != nil {
		handleError(w, http.StatusInternalServerError, "Error processing request: %v", err)
		return
	}
}
