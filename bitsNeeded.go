package asm

import (
	"math"
)

// bitsNeeded tells you how many bits are needed to encode this number.
func bitsNeeded(number int64) int {
	switch {
	case number >= math.MinInt8 && number <= math.MaxInt8:
		return 8

	case number >= math.MinInt16 && number <= math.MaxInt16:
		return 16

	case number >= math.MinInt32 && number <= math.MaxInt32:
		return 32

	default:
		return 64
	}
}
