package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
	"github.com/stretchr/testify/assert"
)

func getTestRouter() *gin.Engine {
	return SetupRouter("test", "v0.0.1-test", log.New("test"))
}

func TestDefaultHandler(t *testing.T) {
	router := getTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
