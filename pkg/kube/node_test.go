package kube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	info := GetNodeInfo()
	assert.NotNil(t, info)
	assert.NotEmpty(t, info.ID)
	assert.NotEmpty(t, info.OS)
	assert.NotNil(t, info.Resources)
	assert.NotNil(t, info.Resources.CPU)
	assert.NotNil(t, info.Resources.RAM)
	assert.NotEmpty(t, info.Resources.CPU.Context)
	assert.NotZero(t, info.Resources.CPU.Value)
	assert.NotEmpty(t, info.Resources.RAM.Context)
	assert.NotZero(t, info.Resources.RAM.Value)
}
