package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResource(t *testing.T) {
	info := GetResourceInfo()
	assert.NotNil(t, info)
	assert.NotNil(t, info.CPU)
	assert.NotNil(t, info.RAM)
	assert.NotEmpty(t, info.CPU.Context)
	assert.NotZero(t, info.CPU.Value)
	assert.NotEmpty(t, info.RAM.Context)
	assert.NotZero(t, info.RAM.Value)
	t.Logf("host: %+v", info)
}
