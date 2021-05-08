package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	assert.Equal(t, "1.1T", byteSize(1234567890000))
	assert.Equal(t, "1.1G", byteSize(1234567890))
	assert.Equal(t, "1.2M", byteSize(1234567))
	assert.Equal(t, "1.2K", byteSize(1234))
	assert.Equal(t, "1B", byteSize(1))
}
