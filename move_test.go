package asm_test

import (
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestMoveRegisterNumber(t *testing.T) {
	usagePatterns := []struct {
		Destination string
		Number      int32
		Code        []byte
	}{
		{
			"rax",
			int32(0xff),
			[]byte{0xb8, 0xff, 0, 0, 0},
		},
		{
			"rcx",
			int32(0xff),
			[]byte{0xb8 + 1, 0xff, 0, 0, 0},
		},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		a.MoveRegisterNumber(pattern.Destination, pattern.Number)
		assert.DeepEqual(t, a.Bytes(), pattern.Code)
	}
}

func TestMoveRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination string
		Source      string
		Code        []byte
	}{
		{
			"rax",
			"rax",
			[]byte{0x48, 0x89, 0xc0},
		},
		{
			"rcx",
			"r9",
			[]byte{0x4c, 0x89, 0xc9},
		},
		{
			"r9",
			"rcx",
			[]byte{0x49, 0x89, 0xc9},
		},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		a.MoveRegisterRegister(pattern.Destination, pattern.Source)
		assert.DeepEqual(t, a.Bytes(), pattern.Code)
	}
}
