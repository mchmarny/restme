package kube

// Measurement represents int measurement
type Measurement struct {
	Value   float64 `json:"value,omitempty"`
	Context string  `json:"context,omitempty"`
}

// ResourceInfo represents node cpu
type ResourceInfo struct {
	RAM *Measurement `json:"ram,omitempty"`
	CPU *Measurement `json:"cpu,omitempty"`
	GPU *Measurement `json:"gpu,omitempty"`
}
