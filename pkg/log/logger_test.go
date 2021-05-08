package log

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHost(t *testing.T) {
	t.Run("create new logger instance", func(t *testing.T) {
		logger := New("test")
		assert.NotNil(t, logger)
	})
	t.Run("test runtime level", func(t *testing.T) {
		os.Unsetenv(LogLevelEnvVar)
		logger := New("test")
		assert.Equal(t, logrus.InfoLevel.String(), logger.GetLevel().String())
	})
	t.Run("test debug level", func(t *testing.T) {
		os.Setenv(LogLevelEnvVar, "debug")
		lev := os.Getenv(LogLevelEnvVar)
		assert.Equal(t, "debug", lev)
		logger := New("test")
		assert.Equal(t, logrus.DebugLevel.String(), logger.GetLevel().String())
	})
}
