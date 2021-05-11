package cpu

import (
	"testing"
	"time"
)

func TestCPULoad(t *testing.T) {
	d, _ := time.ParseDuration("2s")
	c := MakeCPULoad(d)
	t.Logf("count: %+v", c)
}
