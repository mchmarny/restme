package log

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHost(t *testing.T) {
	t.Run("create new logger instance", func(t *testing.T) {
		logger := Default()
		assert.NotNil(t, logger)
	})
	t.Run("test runtime level", func(t *testing.T) {
		logger := New("test", "v0.0.1v", "info", false)
		assert.Equal(t, logrus.InfoLevel.String(), logger.GetLevel().String())
	})
}
