package auth

import (
	"testing"

	"github.com/mchmarny/restme/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestTokenAuthenticator(t *testing.T) {
	if err := createTestKeyFile(); err != nil {
		t.FailNow()
	}

	s, err := NewTokenAuthenticator(testKeyPath, log.Default())
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
