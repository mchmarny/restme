package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestHandler(t *testing.T) {
	h := NewHandler(testLogger)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/request", nil)
	h.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
