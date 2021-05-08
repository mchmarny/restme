package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/kube"
)

// Resource represents simple Kube resource
type Resource struct {
	Request   *RequestMetadata   `json:"request,omitempty"`
	Host      *kube.HostInfo     `json:"host,omitempty"`
	Resources *kube.ResourceInfo `json:"resources,omitempty"`
	Limits    *kube.ResourceInfo `json:"limits,omitempty"`
}

func ResourceHandler(c *gin.Context) {
	result := &Resource{
		Request:   getRequestMetadata(c),
		Host:      kube.GetHostInfo(),
		Resources: kube.GetResourceInfo(),
		Limits:    kube.GetLimits(),
	}

	c.IndentedJSON(http.StatusOK, result)
}
