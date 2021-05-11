package runtime

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
)

// Info represents runtime resources.
type Info struct {
	Request   gin.H         `json:"request,omitempty"`
	Host      *HostInfo     `json:"host,omitempty"`
	Resources *ResourceInfo `json:"resources,omitempty"`
	Limits    *ResourceInfo `json:"limits,omitempty"`
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
		Host:      GetHostInfo(),
		Resources: GetResourceInfo(),
		Limits:    GetLimits(),
	}

	c.IndentedJSON(http.StatusOK, result)
}
