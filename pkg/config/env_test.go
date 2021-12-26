package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	key := "test"
	expected := time.Now().Format(time.RFC3339Nano)
	os.Setenv(key, expected)
	actual := GetEnv(key, "")
	assert.Equal(t, expected, actual)
}
