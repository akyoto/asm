package asm

import (
	"bytes"
	"encoding/binary"

	"github.com/akyoto/asm/sections"
	"github.com/akyoto/asm/syscall"
	"github.com/akyoto/asm/utils"
)

type Assembler struct {
	bytes.Buffer
	Strings        *sections.Strings
	StringPointers []utils.Pointer
}

func New() *Assembler {
	return &Assembler{
		Strings: sections.NewStrings(),
	}
}

func (a *Assembler) AddString(msg string) int64 {
	address := a.Strings.Add(msg)

	a.StringPointers = append(a.StringPointers, utils.Pointer{
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

// MoveRegisterNumber moves a number into the given register.
func (a *Assembler) MoveRegisterNumber(registerNameTo string, num interface{}) {
	baseCode := byte(0xb8)
	registerTo, exists := registers[registerNameTo]

	if !exists {
		panic("Unknown register name: " + registerNameTo)
	}

	// 64-bit operand
	w := byte(0)

	switch num.(type) {
	case string, int64:
		w = 1
	}

	// Register extension
	b := byte(0)

	if registerTo.BaseCodeOffset >= 8 {
		b = 1
	}

	// REX
	if b != 0 || w != 0 {
		a.WriteByte(REX(w, 0, 0, b))
	}

	// Base code
	a.WriteByte(baseCode + registerTo.BaseCodeOffset%8)

	// Number
	switch v := num.(type) {
	case string:
		_ = binary.Write(a, binary.LittleEndian, a.AddString(v))
	default:
		_ = binary.Write(a, binary.LittleEndian, num)
	}
}

// MoveRegisterRegister moves a register value into another register.
func (a *Assembler) MoveRegisterRegister(registerNameTo string, registerNameFrom string) {
	baseCode := byte(0x89)
	registerTo, exists := registers[registerNameTo]

	if !exists {
		panic("Unknown register name: " + registerNameTo)
	}

	registerFrom, exists := registers[registerNameFrom]

	if !exists {
		panic("Unknown register name: " + registerNameFrom)
	}

	if registerTo.BitSize == 64 {
		r := byte(0)
		b := byte(0)

		if registerFrom.BaseCodeOffset >= 8 {
			r = 1
		}

		if registerTo.BaseCodeOffset >= 8 {
			b = 1
		}

		a.WriteByte(REX(1, r, 0, b))
	}

	a.WriteByte(baseCode)
	a.WriteByte(ModRM(0b11, registerFrom.BaseCodeOffset, registerTo.BaseCodeOffset))
}

// PushRegister pushes the value inside the register onto the stack.
func (a *Assembler) PushRegister(registerName string) {
	a.encodeRegister(0x50, registerName)
}

// PopRegister pops a value from the stack and saves it into the register.
func (a *Assembler) PopRegister(registerName string) {
	a.encodeRegister(0x58, registerName)
}

// encodeRegister encodes an instruction that only needs a register name.
func (a *Assembler) encodeRegister(baseCode byte, registerNameTo string) {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		panic("Unknown register name: " + registerNameTo)
	}

	if registerTo.BaseCodeOffset >= 8 {
		a.WriteByte(REX(0, 0, 0, 1))
	}

	a.WriteByte(baseCode + registerTo.BaseCodeOffset%8)
}

func (a *Assembler) Syscall(parameters ...interface{}) {
	for count, parameter := range parameters {
		a.MoveRegisterNumber(syscall.Registers[count], parameter)
	}

	a.WriteBytes(0x0f, 0x05)
}

func (a *Assembler) Println(msg string) {
	a.Print(msg + "\n")
}

func (a *Assembler) Print(msg string) {
	a.Syscall(syscall.Write, int32(1), msg, int32(len(msg)))
}

func (a *Assembler) Open(fileName string) {
	a.Syscall(syscall.Open, 2, fileName, int32(0102), int32(0666))
}

func (a *Assembler) Exit(code int32) {
	a.Syscall(syscall.Exit, code)
}
