package asm

import (
	"bytes"
	"encoding/binary"
)

type Assembler struct {
	bytes.Buffer
	strings *stringTable
}

func New() *Assembler {
	return &Assembler{
		strings: newStringTable(),
	}
}

func (a *Assembler) WriteBytes(someBytes ...byte) {
	for _, b := range someBytes {
		a.WriteByte(b)
	}
}

func (a *Assembler) Mov(registerName string, num interface{}) {
	switch registerName {
	case "rax":
		a.WriteByte(0xb8)
	case "rdi":
		a.WriteByte(0xbf)
	case "rsi":
		a.WriteBytes(0x48, 0xbe)
	case "rdx":
		a.WriteByte(0xba)
	}

	_ = binary.Write(a, binary.LittleEndian, num)
}

func (a *Assembler) Syscall() {
	a.WriteBytes(0x0f, 0x05)
}

func (a *Assembler) Print(msg string) {
	a.Mov("rax", int32(1))
	a.Mov("rdi", int32(1))
	a.Mov("rsi", a.strings.Add(msg))
	a.Mov("rdx", int32(len(msg)))
	a.Syscall()
}

func (a *Assembler) Exit(code int32) {
	a.Mov("rax", int32(60))
	a.Mov("rdi", code)
	a.Syscall()
}
