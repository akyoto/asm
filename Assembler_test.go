package asm_test

import (
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestExit(t *testing.T) {
	a := asm.New()
	a.Exit(0)
	assert.DeepEqual(t, a.Bytes(), []byte{0xb8, 0x3c, 0x00, 0x00, 0x00, 0xbf, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x05})
}

func TestHelloWorld(t *testing.T) {
	a := asm.New()
	a.Println("Hello World")
	a.Exit(0)

	assert.NotNil(t, a.Bytes())
	assert.Equal(t, a.Strings.Count(), 1)
}

func TestProcedures(t *testing.T) {
	a := asm.New()
	a.Call("hello")
	a.Call("niceday")
	a.Call("exit")

	a.AddLabel("hello")
	a.Println("Hello World")
	a.Return()

	a.AddLabel("niceday")
	a.Println("Nice day, isn't it?")
	a.Return()

	a.AddLabel("exit")
	a.Exit(0)

	assert.NotNil(t, a.Bytes())
	assert.Equal(t, a.Strings.Count(), 2)
	assert.Equal(t, len(a.Labels), 3)
}

func TestPushRegister(t *testing.T) {
	a := asm.New()
	a.PushRegister("rbp")
	assert.DeepEqual(t, a.Bytes(), []byte{0x55})
}

func TestPushRegister2(t *testing.T) {
	a := asm.New()
	a.PushRegister("r9")
	assert.DeepEqual(t, a.Bytes(), []byte{0x41, 0x51})
}

func TestPopRegister(t *testing.T) {
	a := asm.New()
	a.PopRegister("rbp")
	assert.DeepEqual(t, a.Bytes(), []byte{0x5d})
}

func TestPopRegister2(t *testing.T) {
	a := asm.New()
	a.PopRegister("r9")
	assert.DeepEqual(t, a.Bytes(), []byte{0x41, 0x59})
}
