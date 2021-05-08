package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mchmarny/restme/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandler(t *testing.T) {
	router := SetupRouter(log.New("TestRequestHandler"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/request", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
