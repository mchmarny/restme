package resource

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/jaypipes/ghw"
	"golang.org/x/sys/unix"
)

const (
	limitMemResourceFile       = "/sys/fs/cgroup/memory/memory.limit_in_bytes"
	limitCPUPeriodResourceFile = "/sys/fs/cgroup/cpu/cpu.cfs_period_us"
	limitCPUQuotaResourceFile  = "/sys/fs/cgroup/cpu/cpu.cfs_quota_us"
)

// GetLimits retreaves limits based on cgroups
func GetLimits() *Info {
	limit := &Info{
		RAM: &Measurement{},
		CPU: &Measurement{},
	}

	// pod memory
	val, wr, ctx := getCGroupsFile(limitMemResourceFile)
	limit.RAM.Value = byteSize(uint64(val))
	limit.RAM.Context = fmt.Sprintf("%s, writable: %v, size: %s", ctx, wr, byteSize(uint64(val)))

	// pod cpu (calculated: quota / period)
	quotaVal, wr, quotaCtx := getCGroupsFile(limitCPUQuotaResourceFile)
	periodVal, _, _ := getCGroupsFile(limitCPUPeriodResourceFile)

	// always context
	limit.CPU.Context = fmt.Sprintf("%s, writable: %v", quotaCtx, wr)

	// value only if there is data
	if quotaVal > 0 && periodVal > 0 {
		limit.CPU.Value = quotaVal / periodVal
	}

	return limit
}

func getCGroupsFile(path string) (val float64, wr bool, info string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return 0, false, fmt.Sprintf("file not found: %s - %v", path, err)
	}

	file, err := os.Open(path)
	if err != nil {
		return 0, false, fmt.Sprintf("unable to open file: %s - %v", path, err)
	}

	bc, err := ioutil.ReadAll(file)
	if err != nil {
		return 0, false, fmt.Sprintf("unable to read file: %s - %v", path, err)
	}

	cs := strings.Trim(string(bc), "\n")

	ic, err := strconv.ParseFloat(cs, 64)
	if err != nil {
		return 0, false, fmt.Sprintf("non-numeric value: %s - %v", cs, err)
	}

	return ic, unix.Access(path, unix.W_OK) == nil, fmt.Sprintf("source: %s", path)
}

func getGPUInfo() (val int, info string) {
	gpu, err := ghw.GPU()
	if err != nil || gpu == nil || gpu.GraphicsCards == nil {
		return 0, fmt.Sprintf("error getting GPU info: %v", err)
	}

	num := len(gpu.GraphicsCards)
	infos := make([]string, 0)

	for _, card := range gpu.GraphicsCards {
		infos = append(infos, fmt.Sprintf("gpu[%d]: %s - %s",
			card.Index, card.DeviceInfo.Vendor.Name, card.DeviceInfo.Product.Name))
	}

	return num, strings.Join(infos, "; ")
}
