package log

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const (
	testLoggerName    = "test"
	testLoggerVersion = "v0.0.1-test"
)

func TestHost(t *testing.T) {
	t.Run("create new logger instance", func(t *testing.T) {
		logger := New(testLoggerName, testLoggerVersion)
		assert.NotNil(t, logger)
	})
	t.Run("test metadata", func(t *testing.T) {
		os.Unsetenv(LogLevelEnvVar)
		logger := New(testLoggerName, testLoggerVersion)
		assert.Equal(t, testLoggerName, logger.Name)
		assert.Equal(t, testLoggerVersion, logger.Version)
	})
	t.Run("test runtime level", func(t *testing.T) {
		os.Unsetenv(LogLevelEnvVar)
		logger := New(testLoggerName, testLoggerVersion)
		assert.Equal(t, logrus.InfoLevel.String(), logger.GetLevel().String())
	})
	t.Run("test debug level", func(t *testing.T) {
		os.Setenv(LogLevelEnvVar, "debug")
		lev := os.Getenv(LogLevelEnvVar)
		assert.Equal(t, "debug", lev)
		logger := New(testLoggerName, testLoggerVersion)
		assert.Equal(t, logrus.DebugLevel.String(), logger.GetLevel().String())
	})
}
