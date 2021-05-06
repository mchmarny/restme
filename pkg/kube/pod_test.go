package kube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPod(t *testing.T) {
	info := GetPodInfo()
	assert.NotNil(t, info)
	assert.NotEmpty(t, info.Hostname)
	assert.NotNil(t, info.Limits)
	assert.NotNil(t, info.Limits.CPU)
	assert.NotNil(t, info.Limits.RAM)
}
