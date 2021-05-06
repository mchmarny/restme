package load

import (
	"encoding/json"
	"runtime"
	"sync"
	"time"
)

type CPULoadResult struct {
	CPUs       int          `json:"cpus,omitempty"`
	Operations int64        `json:"operations,omitempty"`
	Duration   string       `json:"duration,omitempty"`
	lock       sync.RWMutex `json:"-"`
}

func (r *CPULoadResult) add() {
	r.lock.Lock()
	r.Operations++
	r.lock.Unlock()
}

// String returns the JSON serialized representation of the object
func (r *CPULoadResult) String() string {
	s, _ := json.Marshal(r)
	return string(s)
}

// MakeCPULoad creates CPU load for specified duration.
func MakeCPULoad(duration time.Duration) *CPULoadResult {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	result := &CPULoadResult{
		Duration: duration.String(),
		CPUs:     cores,
		lock:     sync.RWMutex{},
	}

	for i := 0; i < cores; i++ {
		go func() {
			runtime.LockOSThread()
			for {
				result.add()
			}
		}()
	}

	// how long
	time.Sleep(duration)

	return result
}
