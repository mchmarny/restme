package runtime

import (
	"encoding/json"

	"github.com/shirou/gopsutil/host"
)

// HostInfo represents the info about this host.
// This is a copy of the gopsutil struct to align JSON names
type HostInfo struct {
	Hostname             string `json:"hostname,omitempty"`
	Uptime               uint64 `json:"uptime,omitempty"`
	BootTime             uint64 `json:"boot_time,omitempty"`
	Processes            uint64 `json:"processes,omitempty"`           // number of processes
	OS                   string `json:"os,omitempty"`                  // ex: freebsd, linux
	Platform             string `json:"platform,omitempty"`            // ex: ubuntu, linuxmint
	PlatformFamily       string `json:"platform_family,omitempty"`     // ex: debian, rhel
	PlatformVersion      string `json:"platform_version,omitempty"`    // version of the complete OS
	KernelVersion        string `json:"kernel_version,omitempty"`      // version of the OS kernel
	KernelArchitecture   string `json:"kernel_architecture,omitempty"` // native cpu architecture
	VirtualizationSystem string `json:"virtualization_system,omitempty"`
	VirtualizationRole   string `json:"virtualization_role,omitempty"` // guest or host
	HostID               string `json:"host_id,omitempty"`             // ex: uuid
	Error                error  `json:"error,omitempty"`               // ex: uuid
}

// String returns the JSON serialized representation of the object
func (h *HostInfo) String() string {
	s, _ := json.Marshal(h)
	return string(s)
}

// GetHostInfo retreaves node info
func GetHostInfo() *HostInfo {
	h := &HostInfo{}
	info, err := host.Info()
	if err != nil {
		h.Error = err
		return h
	}

	h.OS = info.OS
	h.Hostname = info.Hostname
	h.HostID = info.HostID
	h.Uptime = info.Uptime
	h.BootTime = info.BootTime
	h.Processes = info.Procs
	h.Platform = info.Platform
	h.PlatformFamily = info.PlatformFamily
	h.PlatformVersion = info.PlatformVersion
	h.KernelVersion = info.KernelVersion
	h.KernelArchitecture = info.KernelArch
	h.VirtualizationSystem = info.VirtualizationSystem
	h.VirtualizationRole = info.VirtualizationRole

	return h
}
