package load

import (
	"testing"
	"time"
)

func TestHost(t *testing.T) {
	d, _ := time.ParseDuration("2s")
	c := MakeCPULoad(d)
	t.Logf("count: %+v", c)
}
