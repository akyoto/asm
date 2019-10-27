package asm_test

import (
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestModRM(t *testing.T) {
	modRM := asm.ModRM(0b11, 0b10, 0b1)
	assert.Equal(t, modRM, byte(0b11010001))
}
