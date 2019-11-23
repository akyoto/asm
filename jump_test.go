package asm_test

import (
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestJumps(t *testing.T) {
	usagePatterns := []struct {
		JumpCondition string
		Code          []byte
	}{
		{"jl", []byte{0x7c, 0xfe}},
		{"jle", []byte{0x7e, 0xfe}},
		{"jg", []byte{0x7f, 0xfe}},
		{"jge", []byte{0x7d, 0xfe}},
		{"je", []byte{0x74, 0xfe}},
		{"jne", []byte{0x75, 0xfe}},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		label := "label"
		a.AddLabel(label)
		t.Logf("%s %s", pattern.JumpCondition, label)

		switch pattern.JumpCondition {
		case "jl":
			a.JumpIfLess(label)
		case "jle":
			a.JumpIfLessOrEqual(label)
		case "jg":
			a.JumpIfGreater(label)
		case "jge":
			a.JumpIfGreaterOrEqual(label)
		case "je":
			a.JumpIfEqual(label)
		case "jne":
			a.JumpIfNotEqual(label)
		}

		assert.DeepEqual(t, a.Bytes(), pattern.Code)
	}
}
