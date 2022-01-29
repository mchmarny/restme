package request

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/config"
	"github.com/mchmarny/restme/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandler(t *testing.T) {
	c, err := config.GetConfigFromFile("../../configs/unit.json")
	assert.NoError(t, err)
	assert.NotNil(t, c)

	s, err := NewService(log.Default(), c)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	r := gin.Default()
	r.GET("/", s.RequestHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", http.NoBody)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
