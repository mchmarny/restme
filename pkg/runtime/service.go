package runtime

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
	"github.com/mchmarny/restme/pkg/runtime/host"
	"github.com/mchmarny/restme/pkg/runtime/resource"
)

// Info represents runtime resources.
type Info struct {
	Request   gin.H          `json:"request,omitempty"`
	Host      *host.Info     `json:"host,omitempty"`
	Resources *resource.Info `json:"resources,omitempty"`
	Limits    *resource.Info `json:"limits,omitempty"`
}

// NewService creates new ResourceService instance.
func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

// ResourceService provides information about physical resources.
type Service struct {
	logger *log.Logger
}

func (s *Service) ResourceHandler(c *gin.Context) {
	result := &Info{
		Host:      host.GetHostInfo(),
		Resources: resource.GetResourceInfo(),
		Limits:    resource.GetLimits(),
	}

	c.IndentedJSON(http.StatusOK, result)
}
