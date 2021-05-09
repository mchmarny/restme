package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mchmarny/restme/pkg/log"
	"github.com/stretchr/testify/assert"
)

const (
	testLoggerName    = "test"
	testLoggerVersion = "v0.0.1-test"
)

var (
	testLogger = log.New(testLoggerName, testLoggerVersion)
)

func TestDefaultHandler(t *testing.T) {
	h := NewHandler(testLogger)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	h.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
