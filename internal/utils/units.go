package utils

import (
	"math"
)

// ByteSize represents the size of a value in bits.
type ByteSize int64

const (
	B  ByteSize = 1
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

var shortUnitMap = map[string]ByteSize{
	"B":  B,
	"KB": KB,
	"MB": MB,
	"GB": GB,
	"TB": TB,
	"PB": PB,
	"EB": EB,
}

// PrettyByteSize ...
func PrettyByteSize(b float64, size string) int64 {
	s, ok := shortUnitMap[size]
	if !ok {
		return 0
	}

	b = b * float64(s)
	b = math.Round(b * 10)

	return int64(b)
}
