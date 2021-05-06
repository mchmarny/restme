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
	Meta *RequestMetadata `json:"meta,omitempty"`
	Node *kube.NodeInfo   `json:"node,omitempty"`
	Pod  *kube.PodInfo    `json:"pod,omitempty"`
}

func (h ResourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("serving: %+v", r)

	if r.URL.Path != h.Path {
		handleError(w, http.StatusNotFound, "Expected: %s, got:%s", h.Path, r.URL.Path)
		return
	}

	result := &Resource{
		Meta: getRequestMetadata(r),
		Node: kube.GetNodeInfo(),
		Pod:  kube.GetPodInfo(),
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(result); err != nil {
		handleError(w, http.StatusInternalServerError, "Error processing request: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
