package kube

import (
	"strconv"
	"strings"
)

type ByteSize uint64

const (
	// const for file sizes
	B  ByteSize = 1
	KB          = B << 10
	MB          = KB << 10
	GB          = MB << 10
	TB          = GB << 10
	PB          = TB << 10
	EB          = PB << 10
)

// byteSize returns a human-readable byte string of the form 10M, 12.5K, and so forth.
// The following units are available:
//	T: Terabyte
//	G: Gigabyte
//	M: Megabyte
//	K: Kilobyte
//	B: Byte
// The unit that results in the smallest number greater than or equal to 1 is always chosen.
func byteSize(bytes uint64) string {
	unit := ""
	value := float64(bytes)

	switch {
	case bytes >= uint64(TB):
		unit = "T"
		value /= float64(TB)
	case bytes >= uint64(GB):
		unit = "G"
		value /= float64(GB)
	case bytes >= uint64(MB):
		unit = "M"
		value /= float64(MB)
	case bytes >= uint64(KB):
		unit = "K"
		value /= float64(KB)
	case bytes >= uint64(B):
		unit = "B"
	case bytes == 0:
		return "0"
	}

	result := strconv.FormatFloat(value, 'f', 1, 64)
	result = strings.TrimSuffix(result, ".0")
	return result + unit
}
