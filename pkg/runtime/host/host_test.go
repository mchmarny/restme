package host

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHost(t *testing.T) {
	info := GetHostInfo()
	assert.NotNil(t, info)
	assert.NotEmpty(t, info.HostID)
	assert.NotEmpty(t, info.Hostname)
	assert.NotEmpty(t, info.OS)
	assert.NotEmpty(t, info.Platform)
	t.Logf("host: %+v", info)
}
