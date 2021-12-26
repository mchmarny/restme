package log

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHost(t *testing.T) {
	t.Run("create new logger instance", func(t *testing.T) {
		logger := Default()
		assert.NotNil(t, logger)
	})
	t.Run("test metadata", func(t *testing.T) {
		os.Unsetenv(LogLevelEnvVar)
		appName := "test1"
		appVersion := "v0.0.1-test"
		logger := New(appName, appVersion)
		assert.Equal(t, appName, logger.GetAppName())
		assert.Equal(t, appVersion, logger.GetAppVersion())
	})
	t.Run("test runtime level", func(t *testing.T) {
		os.Unsetenv(LogLevelEnvVar)
		logger := Default()
		assert.Equal(t, logrus.InfoLevel.String(), logger.GetLevel().String())
	})
	t.Run("test debug level", func(t *testing.T) {
		os.Setenv(LogLevelEnvVar, "debug")
		lev := os.Getenv(LogLevelEnvVar)
		assert.Equal(t, "debug", lev)
		logger := Default()
		assert.Equal(t, logrus.DebugLevel.String(), logger.GetLevel().String())
	})
}
