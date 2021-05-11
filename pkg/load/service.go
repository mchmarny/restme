package load

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/load/cpu"
	"github.com/mchmarny/restme/pkg/log"
)

// Info represents simple HTTP load result
type Info struct {
	Request gin.H           `json:"request,omitempty"`
	Result  *cpu.LoadResult `json:"result,omitempty"`
}

// NewLoadService creates new LoadService instance.
func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

// LoadService creates resource load.
type Service struct {
	logger *log.Logger
}

func (s *Service) CPULoadHandler(c *gin.Context) {
	durStr := c.Param("duration")
	duration, err := time.ParseDuration(durStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid duration parameter": durStr})
		return
	}

	result := &Info{
		Request: gin.H{
			"time":     time.Now().UTC(),
			"duration": duration,
		},
		Result: cpu.MakeCPULoad(duration),
	}

	c.IndentedJSON(http.StatusOK, result)
}
