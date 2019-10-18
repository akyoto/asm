package asm

import (
	"bytes"
	"encoding/binary"

	"github.com/akyoto/asm/stringtable"
)

type Assembler struct {
	bytes.Buffer
	StringTable     *stringtable.StringTable
	SectionPointers []Pointer
}

func New() *Assembler {
	return &Assembler{
		StringTable: stringtable.New(),
	}
}

func (a *Assembler) AddString(msg string) int64 {
	address := a.StringTable.Add(msg)

	a.SectionPointers = append(a.SectionPointers, Pointer{
		Address:  address,
		Position: a.Len(),
	})

	return address
}

func (a *Assembler) WriteBytes(someBytes ...byte) {
	for _, b := range someBytes {
		a.WriteByte(b)
	}
}

func (a *Assembler) Mov(registerName string, num interface{}) {
	baseCode := byte(0xb8)
	registerID, exists := registerIDs[registerName]

	if !exists {
		panic("Unknown register name: " + registerName)
	}

	switch num.(type) {
	case string, int64:
		a.WriteByte(REX(1, 0, 0, 0))
	}

	a.WriteByte(baseCode + registerID)

	switch v := num.(type) {
	case string:
		_ = binary.Write(a, binary.LittleEndian, a.AddString(v))
	default:
		_ = binary.Write(a, binary.LittleEndian, num)
	}
}

func (a *Assembler) Syscall() {
	a.WriteBytes(0x0f, 0x05)
}

func (a *Assembler) Println(msg string) {
	a.Print(msg + "\n")
}

func (a *Assembler) Print(msg string) {
	a.Mov("rax", int32(1))
	a.Mov("rdi", int32(1))
	a.Mov("rsi", msg)
	a.Mov("rdx", int32(len(msg)))
	a.Syscall()
}

func (a *Assembler) Exit(code int32) {
	a.Mov("rax", int32(60))
	a.Mov("rdi", code)
	a.Syscall()
}
