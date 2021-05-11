package load

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestCPULoad(t *testing.T) {
	d, _ := time.ParseDuration("2s")
	c := MakeCPULoad(d)
	t.Logf("count: %+v", c)
}

func TestCPULoadHandler(t *testing.T) {
	s := NewService(log.Default())
	assert.NotNil(t, s)

	r := gin.Default()
	r.GET("/:duration", s.CPULoadHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/3s", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
