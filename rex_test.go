package asm_test

import (
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestREX(t *testing.T) {
	rex := asm.MakeREXPrefix(1, 0, 0, 0)
	assert.Equal(t, rex, byte(0x48))
}
