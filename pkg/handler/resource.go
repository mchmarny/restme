package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/runtime"
)

// Resource represents simple Kube resource
type Resource struct {
	Request   gin.H                 `json:"request,omitempty"`
	Host      *runtime.HostInfo     `json:"host,omitempty"`
	Resources *runtime.ResourceInfo `json:"resources,omitempty"`
	Limits    *runtime.ResourceInfo `json:"limits,omitempty"`
}

func (h *Handler) ResourceHandler(c *gin.Context) {
	result := &Resource{
		Request:   h.getRequestMetadata(c),
		Host:      runtime.GetHostInfo(),
		Resources: runtime.GetResourceInfo(),
		Limits:    runtime.GetLimits(),
	}

	c.IndentedJSON(http.StatusOK, result)
}
