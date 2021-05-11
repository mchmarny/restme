package echo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestEchoHandler(t *testing.T) {
	s := NewService(log.Default())
	assert.NotNil(t, s)

	r := gin.Default()
	r.POST("/", s.MessageHandler)

	m := Message{
		On:      time.Now().Unix(),
		Message: "test",
	}

	b, err := json.Marshal(m)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var m2 Message
	err = json.NewDecoder(w.Result().Body).Decode(&m2)
	assert.Nilf(t, err, "error decoding body: %v", err)

	assert.Equal(t, m, m2)
}
