package kube

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// Measurement represents int measurement
type Measurement struct {
	Value   interface{} `json:"value,omitempty"`
	Context string      `json:"context,omitempty"`
}

// ResourceInfo represents node cpu
type ResourceInfo struct {
	RAM *Measurement `json:"ram,omitempty"`
	CPU *Measurement `json:"cpu,omitempty"`
	GPU *Measurement `json:"gpu,omitempty"`
}

// String returns the JSON serialized representation of the object
func (h *ResourceInfo) String() string {
	s, _ := json.Marshal(h)
	return string(s)
}

// GetResourceInfo retreaves node info
func GetResourceInfo() *ResourceInfo {
	resource := &ResourceInfo{
		RAM: &Measurement{},
		CPU: &Measurement{},
	}

	// vm
	vm, err := mem.VirtualMemory()
	if err == nil {
		resource.RAM.Value = byteSize(vm.Total)
		resource.RAM.Context = fmt.Sprintf(
			"Source: OS process status, Size: %s", byteSize(vm.Total))
	}

	// cpu
	count, err := cpu.Counts(true)
	if err == nil {
		resource.CPU.Value = count
		resource.CPU.Context = "Source: OS process status"
	}

	// gpu
	gpuCount, deviceInfo := getGPUInfo()
	if gpuCount > 0 {
		resource.GPU = &Measurement{
			Value:   gpuCount,
			Context: deviceInfo,
		}
	}

	return resource
}
