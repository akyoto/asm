package elf_test

import (
	"os"
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
	"github.com/akyoto/assert"
)

func TestHelloWorld(t *testing.T) {
	defer os.Remove("test.out")

	a := asm.New()
	a.Println("Hello World")
	a.Exit(0)

	err := elf.New(a).WriteToFile("test.out")
	assert.Nil(t, err)
}
