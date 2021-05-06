package load

import (
	"testing"
	"time"
)

func TestHost(t *testing.T) {
	d, _ := time.ParseDuration("3s")
	c := MakeCPULoad(d)
	t.Logf("count: %+v", c)
}
