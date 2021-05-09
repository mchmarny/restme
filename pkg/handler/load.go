package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/load"
)

// Load represents simple HTTP load result
type LoadResult struct {
	Request gin.H               `json:"request,omitempty"`
	Result  *load.CPULoadResult `json:"result,omitempty"`
}

func (h *Handler) LoadHandler(c *gin.Context) {
	durStr := c.Param("duration")
	duration, err := time.ParseDuration(durStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid duration parameter": durStr})
		return
	}

	result := &LoadResult{
		Request: h.getRequestMetadata(c),
		Result:  load.MakeCPULoad(duration),
	}

	c.IndentedJSON(http.StatusOK, result)
}
