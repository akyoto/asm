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
		{"rax", []byte{0x50}},
		{"ax", []byte{0x66, 0x50}},
		{"rcx", []byte{0x51}},
		{"cx", []byte{0x66, 0x51}},
		{"rdx", []byte{0x52}},
		{"dx", []byte{0x66, 0x52}},
		{"rbx", []byte{0x53}},
		{"bx", []byte{0x66, 0x53}},
		{"rsi", []byte{0x56}},
		{"si", []byte{0x66, 0x56}},
		{"rdi", []byte{0x57}},
		{"di", []byte{0x66, 0x57}},
		{"rsp", []byte{0x54}},
		{"sp", []byte{0x66, 0x54}},
		{"rbp", []byte{0x55}},
		{"bp", []byte{0x66, 0x55}},
		{"r8", []byte{0x41, 0x50}},
		{"r8w", []byte{0x66, 0x41, 0x50}},
		{"r9", []byte{0x41, 0x51}},
		{"r9w", []byte{0x66, 0x41, 0x51}},
		{"r10", []byte{0x41, 0x52}},
		{"r10w", []byte{0x66, 0x41, 0x52}},
		{"r11", []byte{0x41, 0x53}},
		{"r11w", []byte{0x66, 0x41, 0x53}},
		{"r12", []byte{0x41, 0x54}},
		{"r12w", []byte{0x66, 0x41, 0x54}},
		{"r13", []byte{0x41, 0x55}},
		{"r13w", []byte{0x66, 0x41, 0x55}},
		{"r14", []byte{0x41, 0x56}},
		{"r14w", []byte{0x66, 0x41, 0x56}},
		{"r15", []byte{0x41, 0x57}},
		{"r15w", []byte{0x66, 0x41, 0x57}},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		t.Logf("push %s", pattern.Register)
		a.PushRegister(pattern.Register)
		assert.DeepEqual(t, a.Code(), pattern.Code)
	}
}

func TestPopRegister(t *testing.T) {
	usagePatterns := []struct {
		Register string
		Code     []byte
	}{
		{"rax", []byte{0x58}},
		{"ax", []byte{0x66, 0x58}},
		{"rcx", []byte{0x59}},
		{"cx", []byte{0x66, 0x59}},
		{"rdx", []byte{0x5a}},
		{"dx", []byte{0x66, 0x5a}},
		{"rbx", []byte{0x5b}},
		{"bx", []byte{0x66, 0x5b}},
		{"rsi", []byte{0x5e}},
		{"si", []byte{0x66, 0x5e}},
		{"rdi", []byte{0x5f}},
		{"di", []byte{0x66, 0x5f}},
		{"rsp", []byte{0x5c}},
		{"sp", []byte{0x66, 0x5c}},
		{"rbp", []byte{0x5d}},
		{"bp", []byte{0x66, 0x5d}},
		{"r8", []byte{0x41, 0x58}},
		{"r8w", []byte{0x66, 0x41, 0x58}},
		{"r9", []byte{0x41, 0x59}},
		{"r9w", []byte{0x66, 0x41, 0x59}},
		{"r10", []byte{0x41, 0x5a}},
		{"r10w", []byte{0x66, 0x41, 0x5a}},
		{"r11", []byte{0x41, 0x5b}},
		{"r11w", []byte{0x66, 0x41, 0x5b}},
		{"r12", []byte{0x41, 0x5c}},
		{"r12w", []byte{0x66, 0x41, 0x5c}},
		{"r13", []byte{0x41, 0x5d}},
		{"r13w", []byte{0x66, 0x41, 0x5d}},
		{"r14", []byte{0x41, 0x5e}},
		{"r14w", []byte{0x66, 0x41, 0x5e}},
		{"r15", []byte{0x41, 0x5f}},
		{"r15w", []byte{0x66, 0x41, 0x5f}},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		t.Logf("pop %s", pattern.Register)
		a.PopRegister(pattern.Register)
		assert.DeepEqual(t, a.Code(), pattern.Code)
	}
}
