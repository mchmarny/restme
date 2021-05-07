package load

import (
	"encoding/json"
	"runtime"
	"sync"
	"time"
)

type CPULoadResult struct {
	Cores      int    `json:"cores,omitempty"`
	Start      int64  `json:"start,omitempty"`
	End        int64  `json:"end,omitempty"`
	Operations int64  `json:"operations,omitempty"`
	Duration   string `json:"duration,omitempty"`
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

	var counter int64

	start := time.Now()
	lock := sync.RWMutex{}
	countCh := make(chan bool, cores)

	// load for each test
	for i := 0; i < cores; i++ {
		go func() {
			runtime.LockOSThread()
			for {
				countCh <- true
			}
		}()
	}

	// count the operations
	for {
		<-countCh
		lock.Lock()
		counter++
		lock.Unlock()

		if time.Since(start) >= duration {
			break
		}
	}

	return &CPULoadResult{
		Start:      start.Unix(),
		End:        time.Now().Unix(),
		Duration:   duration.String(),
		Cores:      cores,
		Operations: counter,
	}
}
