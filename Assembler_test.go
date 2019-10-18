package asm_test

import (
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestAssembler(t *testing.T) {
	a := asm.New()
	a.Print("Hello World\n")
	a.Exit(0)
	assert.NotEqual(t, len(a.Bytes()), 0)
}
