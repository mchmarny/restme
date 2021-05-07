package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mchmarny/restme/pkg/kube"
)

// Resource represents simple Kube resource
type Resource struct {
	Request   *RequestMetadata   `json:"request,omitempty"`
	Host      *kube.HostInfo     `json:"host,omitempty"`
	Resources *kube.ResourceInfo `json:"resources,omitempty"`
	Limits    *kube.ResourceInfo `json:"limits,omitempty"`
}

func ResourceHandler(w http.ResponseWriter, r *http.Request) {
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
