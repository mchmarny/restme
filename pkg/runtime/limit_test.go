package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPod(t *testing.T) {
	info := GetLimits()
	assert.NotNil(t, info)
	assert.NotNil(t, info.CPU)
	assert.NotNil(t, info.RAM)
	t.Logf("pod: %+v", info)
}
