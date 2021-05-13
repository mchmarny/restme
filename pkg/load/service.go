package load

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/httputil"
	"github.com/mchmarny/restme/pkg/load/cpu"
	"github.com/mchmarny/restme/pkg/log"
	"github.com/pkg/errors"
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

// CPULoadHandler godoc
// @Summary generates CPU load
// @Description CPU load
// @Tags load,cpu
// @Accept json
// @Produce json
// @Param duration path string true "Time Duration"
// @Success 200 {string} string "info"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /load/cpu/{duration} [get]
func (s *Service) CPULoadHandler(c *gin.Context) {
	durStr := c.Param("duration")
	duration, err := time.ParseDuration(durStr)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, errors.Errorf("Invalid duration parameter: %s", durStr))
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
