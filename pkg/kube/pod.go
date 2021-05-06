package kube

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/host"

	"golang.org/x/sys/unix"

	"github.com/jaypipes/ghw"
)

// PodInfo represents pod info
type PodInfo struct {
	Hostname string        `json:"hostname,omitempty"`
	Limits   *ResourceInfo `json:"limits,omitempty"`
}

// GetPodInfo retreaves pod info
func GetPodInfo() *PodInfo {
	pod := &PodInfo{
		Limits: &ResourceInfo{
			RAM: &Measurement{},
			CPU: &Measurement{},
		},
	}

	// host
	info, err := host.Info()
	if err == nil {
		pod.Hostname = info.Hostname
	}

	// pod memory
	val, wr, ctx := getCGroupsFile(limitMemResourceFile)
	pod.Limits.RAM.Value = val
	pod.Limits.RAM.Context = fmt.Sprintf("%s, Writable: %v, Size: %s", ctx, wr, byteSize(uint64(val)))

	// pod cpu (calculated: quota / period)
	quotaVal, wr, quotaCtx := getCGroupsFile(limitCPUQuotaResourceFile)
	periodVal, _, _ := getCGroupsFile(limitCPUPeriodResourceFile)

	// always context
	pod.Limits.CPU.Context = fmt.Sprintf("%s, Writable: %v", quotaCtx, wr)

	// value only if there is data
	if quotaVal > 0 && periodVal > 0 {
		pod.Limits.CPU.Value = quotaVal / periodVal
	}

	return pod
}

func getCGroupsFile(path string) (val float64, wr bool, info string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("File[%s]: not found: %v", path, err)
		return 0, false, fmt.Sprintf("File not found: %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening file[%s]: %v", path, err)
		return 0, false, fmt.Sprintf("Unable to open file: %s", path)
	}

	bc, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file[%s]: %v", path, err)
		return 0, false, fmt.Sprintf("Unable to read file: %s", path)
	}

	cs := strings.Trim(string(bc), "\n")

	ic, err := strconv.ParseFloat(cs, 64)
	if err != nil {
		log.Printf("Error parsing content[%s]: %v", path, err)
		return 0, false, fmt.Sprintf("Non-numeric value: %s", cs)
	}

	log.Printf("File[%s]: %v = %f", path, cs, ic)

	return ic, unix.Access(path, unix.W_OK) == nil, fmt.Sprintf("Source: %s", path)
}

func getGPUInfo() (val int, info string) {
	gpu, err := ghw.GPU()
	if err != nil || gpu == nil || gpu.GraphicsCards == nil {
		log.Printf("Error getting GPU info: %v", err)
		return 0, ""
	}

	num := len(gpu.GraphicsCards)
	infos := make([]string, 0)

	for _, card := range gpu.GraphicsCards {
		log.Printf(" %v\n", card)
		infos = append(infos, fmt.Sprintf("gpu[%d]: %s - %s",
			card.Index, card.DeviceInfo.Vendor.Name, card.DeviceInfo.Product.Name))
	}

	return num, strings.Join(infos, "; ")
}
