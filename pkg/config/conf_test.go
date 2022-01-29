package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	_, err := GetConfigFromFile("../file/not/exist.json")
	assert.Error(t, err)

	c, err := GetConfigFromFile("../../configs/unit.json")
	assert.NoError(t, err)
	assert.NotNil(t, c)

	assert.NotNil(t, c.Auth)
	assert.NotEmpty(t, c.Auth.Tokens)

	assert.NotNil(t, c.Log)
	assert.NotEmpty(t, c.Log.Level)

	assert.NotNil(t, c.IP)
}
