package asm_test

import (
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestIncreaseRegister(t *testing.T) {
	usagePatterns := []struct {
		Register string
		Code     []byte
	}{
		{"rax", []byte{0x48, 0xff, 0xc0}},
		{"eax", []byte{0xff, 0xc0}},
		{"ax", []byte{0x66, 0xff, 0xc0}},
		{"rcx", []byte{0x48, 0xff, 0xc1}},
		{"ecx", []byte{0xff, 0xc1}},
		{"cx", []byte{0x66, 0xff, 0xc1}},
		{"rdx", []byte{0x48, 0xff, 0xc2}},
		{"edx", []byte{0xff, 0xc2}},
		{"dx", []byte{0x66, 0xff, 0xc2}},
		{"rbx", []byte{0x48, 0xff, 0xc3}},
		{"ebx", []byte{0xff, 0xc3}},
		{"bx", []byte{0x66, 0xff, 0xc3}},
		{"rsi", []byte{0x48, 0xff, 0xc6}},
		{"esi", []byte{0xff, 0xc6}},
		{"si", []byte{0x66, 0xff, 0xc6}},
		{"rdi", []byte{0x48, 0xff, 0xc7}},
		{"edi", []byte{0xff, 0xc7}},
		{"di", []byte{0x66, 0xff, 0xc7}},
		{"rsp", []byte{0x48, 0xff, 0xc4}},
		{"esp", []byte{0xff, 0xc4}},
		{"sp", []byte{0x66, 0xff, 0xc4}},
		{"rbp", []byte{0x48, 0xff, 0xc5}},
		{"ebp", []byte{0xff, 0xc5}},
		{"bp", []byte{0x66, 0xff, 0xc5}},
		{"r8", []byte{0x49, 0xff, 0xc0}},
		{"r8d", []byte{0x41, 0xff, 0xc0}},
		{"r8w", []byte{0x66, 0x41, 0xff, 0xc0}},
		{"r9", []byte{0x49, 0xff, 0xc1}},
		{"r9d", []byte{0x41, 0xff, 0xc1}},
		{"r9w", []byte{0x66, 0x41, 0xff, 0xc1}},
		{"r10", []byte{0x49, 0xff, 0xc2}},
		{"r10d", []byte{0x41, 0xff, 0xc2}},
		{"r10w", []byte{0x66, 0x41, 0xff, 0xc2}},
		{"r11", []byte{0x49, 0xff, 0xc3}},
		{"r11d", []byte{0x41, 0xff, 0xc3}},
		{"r11w", []byte{0x66, 0x41, 0xff, 0xc3}},
		{"r12", []byte{0x49, 0xff, 0xc4}},
		{"r12d", []byte{0x41, 0xff, 0xc4}},
		{"r12w", []byte{0x66, 0x41, 0xff, 0xc4}},
		{"r13", []byte{0x49, 0xff, 0xc5}},
		{"r13d", []byte{0x41, 0xff, 0xc5}},
		{"r13w", []byte{0x66, 0x41, 0xff, 0xc5}},
		{"r14", []byte{0x49, 0xff, 0xc6}},
		{"r14d", []byte{0x41, 0xff, 0xc6}},
		{"r14w", []byte{0x66, 0x41, 0xff, 0xc6}},
		{"r15", []byte{0x49, 0xff, 0xc7}},
		{"r15d", []byte{0x41, 0xff, 0xc7}},
		{"r15w", []byte{0x66, 0x41, 0xff, 0xc7}},

		// 1-byte registers
		{"al", []byte{0xfe, 0xc0}},
		{"cl", []byte{0xfe, 0xc1}},
		{"dl", []byte{0xfe, 0xc2}},
		{"bl", []byte{0xfe, 0xc3}},
		{"r8b", []byte{0x41, 0xfe, 0xc0}},
		{"r9b", []byte{0x41, 0xfe, 0xc1}},
		{"r10b", []byte{0x41, 0xfe, 0xc2}},
		{"r11b", []byte{0x41, 0xfe, 0xc3}},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		t.Logf("inc %s", pattern.Register)
		a.IncreaseRegister(pattern.Register)
		assert.DeepEqual(t, a.Code(), pattern.Code)
	}
}

func TestDecreaseRegister(t *testing.T) {
	usagePatterns := []struct {
		Register string
		Code     []byte
	}{
		{"rax", []byte{0x48, 0xff, 0xc8}},
		{"eax", []byte{0xff, 0xc8}},
		{"ax", []byte{0x66, 0xff, 0xc8}},
		{"rcx", []byte{0x48, 0xff, 0xc9}},
		{"ecx", []byte{0xff, 0xc9}},
		{"cx", []byte{0x66, 0xff, 0xc9}},
		{"rdx", []byte{0x48, 0xff, 0xca}},
		{"edx", []byte{0xff, 0xca}},
		{"dx", []byte{0x66, 0xff, 0xca}},
		{"rbx", []byte{0x48, 0xff, 0xcb}},
		{"ebx", []byte{0xff, 0xcb}},
		{"bx", []byte{0x66, 0xff, 0xcb}},
		{"rsi", []byte{0x48, 0xff, 0xce}},
		{"esi", []byte{0xff, 0xce}},
		{"si", []byte{0x66, 0xff, 0xce}},
		{"rdi", []byte{0x48, 0xff, 0xcf}},
		{"edi", []byte{0xff, 0xcf}},
		{"di", []byte{0x66, 0xff, 0xcf}},
		{"rsp", []byte{0x48, 0xff, 0xcc}},
		{"esp", []byte{0xff, 0xcc}},
		{"sp", []byte{0x66, 0xff, 0xcc}},
		{"rbp", []byte{0x48, 0xff, 0xcd}},
		{"ebp", []byte{0xff, 0xcd}},
		{"bp", []byte{0x66, 0xff, 0xcd}},
		{"r8", []byte{0x49, 0xff, 0xc8}},
		{"r8d", []byte{0x41, 0xff, 0xc8}},
		{"r8w", []byte{0x66, 0x41, 0xff, 0xc8}},
		{"r9", []byte{0x49, 0xff, 0xc9}},
		{"r9d", []byte{0x41, 0xff, 0xc9}},
		{"r9w", []byte{0x66, 0x41, 0xff, 0xc9}},
		{"r10", []byte{0x49, 0xff, 0xca}},
		{"r10d", []byte{0x41, 0xff, 0xca}},
		{"r10w", []byte{0x66, 0x41, 0xff, 0xca}},
		{"r11", []byte{0x49, 0xff, 0xcb}},
		{"r11d", []byte{0x41, 0xff, 0xcb}},
		{"r11w", []byte{0x66, 0x41, 0xff, 0xcb}},
		{"r12", []byte{0x49, 0xff, 0xcc}},
		{"r12d", []byte{0x41, 0xff, 0xcc}},
		{"r12w", []byte{0x66, 0x41, 0xff, 0xcc}},
		{"r13", []byte{0x49, 0xff, 0xcd}},
		{"r13d", []byte{0x41, 0xff, 0xcd}},
		{"r13w", []byte{0x66, 0x41, 0xff, 0xcd}},
		{"r14", []byte{0x49, 0xff, 0xce}},
		{"r14d", []byte{0x41, 0xff, 0xce}},
		{"r14w", []byte{0x66, 0x41, 0xff, 0xce}},
		{"r15", []byte{0x49, 0xff, 0xcf}},
		{"r15d", []byte{0x41, 0xff, 0xcf}},
		{"r15w", []byte{0x66, 0x41, 0xff, 0xcf}},

		// 1-byte registers
		{"al", []byte{0xfe, 0xc8}},
		{"cl", []byte{0xfe, 0xc9}},
		{"dl", []byte{0xfe, 0xca}},
		{"bl", []byte{0xfe, 0xcb}},
		{"r8b", []byte{0x41, 0xfe, 0xc8}},
		{"r9b", []byte{0x41, 0xfe, 0xc9}},
		{"r10b", []byte{0x41, 0xfe, 0xca}},
		{"r11b", []byte{0x41, 0xfe, 0xcb}},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		t.Logf("dec %s", pattern.Register)
		a.DecreaseRegister(pattern.Register)
		assert.DeepEqual(t, a.Code(), pattern.Code)
	}
}
