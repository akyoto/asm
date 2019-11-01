package asm_test

import (
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestPushRegister(t *testing.T) {
	usagePatterns := []struct {
		Register string
		Code     []byte
	}{
		{
			"rbp",
			[]byte{0x55},
		},
		{
			"r9",
			[]byte{0x41, 0x51},
		},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		a.PushRegister(pattern.Register)
		assert.DeepEqual(t, a.Bytes(), pattern.Code)
	}
}

func TestPopRegister(t *testing.T) {
	usagePatterns := []struct {
		Register string
		Code     []byte
	}{
		{
			"rbp",
			[]byte{0x5d},
		},
		{
			"r9",
			[]byte{0x41, 0x59},
		},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		a.PopRegister(pattern.Register)
		assert.DeepEqual(t, a.Bytes(), pattern.Code)
	}
}
