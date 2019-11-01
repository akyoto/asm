package asm

import (
	"encoding/binary"
	"log"
)

// MoveRegisterNumber moves a number into the given register.
func (a *Assembler) MoveRegisterNumber(registerNameTo string, num interface{}) {
	baseCode := byte(0xb8)
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
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
		a.WriteBytes(REX(w, 0, 0, b))
	}

	// Base code
	a.WriteBytes(baseCode + registerTo.BaseCodeOffset%8)

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
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	registerFrom, exists := registers[registerNameFrom]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameFrom)
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

		a.WriteBytes(REX(1, r, 0, b))
	}

	a.WriteBytes(baseCode)
	a.WriteBytes(ModRM(0b11, registerFrom.BaseCodeOffset, registerTo.BaseCodeOffset))
}
