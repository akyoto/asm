package opcode_test

import (
	"testing"

	"github.com/akyoto/asm/opcode"
	"github.com/akyoto/assert"
)

func TestModRM(t *testing.T) {
	modRM := opcode.ModRM(0b11, 0b10, 0b1)
	assert.Equal(t, modRM, byte(0b11010001))
}
