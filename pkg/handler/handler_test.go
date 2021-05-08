package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mchmarny/restme/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestDefaultHandler(t *testing.T) {
	router := SetupRouter(log.New("DefaultHandler", true))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
