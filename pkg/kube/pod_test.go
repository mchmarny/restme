package kube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPod(t *testing.T) {
	info := GetPodInfo()
	assert.NotNil(t, info)
	assert.NotNil(t, info.Host)
	assert.NotNil(t, info.Limits)
	assert.NotNil(t, info.Limits.CPU)
	assert.NotNil(t, info.Limits.RAM)
}
