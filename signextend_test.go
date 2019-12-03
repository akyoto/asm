package asm_test

import (
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestSignExtendToDX(t *testing.T) {
	usagePatterns := []struct {
		Register string
		Code     []byte
	}{
		{"rax", []byte{0x48, 0x99}},
		{"eax", []byte{0x99}},
		{"ax", []byte{0x66, 0x99}},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		t.Logf("cdq %s", pattern.Register)
		a.SignExtendToDX(pattern.Register)
		assert.DeepEqual(t, a.Bytes(), pattern.Code)
	}
}
