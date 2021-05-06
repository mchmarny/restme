package kube

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// NodeInfo represents host info data
type NodeInfo struct {
	ID        string        `json:"hostId,omitempty"`
	BootTime  time.Time     `json:"bootTs,omitempty"`
	OS        string        `json:"os,omitempty"`
	Resources *ResourceInfo `json:"resources,omitempty"`
}

// GetNodeInfo retreaves node info
func GetNodeInfo() *NodeInfo {
	node := &NodeInfo{
		Resources: &ResourceInfo{
			RAM: &Measurement{},
			CPU: &Measurement{},
		},
	}

	// host
	info, err := host.Info()
	if err == nil {
		node.ID = info.HostID
		node.BootTime = time.Unix(int64(info.BootTime), 0)
		node.OS = info.OS
	}

	// vm
	vm, err := mem.VirtualMemory()
	if err == nil {
		node.Resources.RAM.Value = float64(vm.Total)
		node.Resources.RAM.Context = fmt.Sprintf(
			"Source: OS process status, Size: %s", byteSize(vm.Total))
	}

	// cpu
	count, err := cpu.Counts(true)
	if err == nil {
		node.Resources.CPU.Value = float64(count)
		node.Resources.CPU.Context = "Source: OS process status"
	}

	// gpu
	gpuCount, deviceInfo := getGPUInfo()
	if gpuCount > 0 {
		node.Resources.GPU = &Measurement{
			Value:   float64(gpuCount),
			Context: deviceInfo,
		}
	}

	return node
}
