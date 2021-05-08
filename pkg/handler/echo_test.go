package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEchoHandler(t *testing.T) {
	router := getTestRouter()

	m := message{
		On:      time.Now().Unix(),
		Message: "test",
	}

	b, err := json.Marshal(m)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/echo", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var m2 message
	err = json.NewDecoder(w.Result().Body).Decode(&m2)
	assert.Nilf(t, err, "error decoding body: %v", err)

	assert.Equal(t, m, m2)
}
