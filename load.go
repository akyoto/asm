package asm

import (
	"log"

	"github.com/akyoto/asm/opcode"
)

// LoadRegister loads from memory into a register.
func (a *Assembler) LoadRegister(registerNameTo string, registerNameFrom string, offset byte, byteCount byte) {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	registerFrom, exists := registers[registerNameFrom]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameFrom)
	}

	baseCode := byte(0x8b)

	switch byteCount {
	case 2:
		a.WriteBytes(0x66)

	case 1:
		baseCode = 0x8a
	}

	// REX prefix
	w := byte(0) // Indicates a 64-bit register.
	r := byte(0) // Extension to the "reg" field in ModRM.
	x := byte(0) // Extension to the SIB index field.
	b := byte(0) // Extension to the "rm" field in ModRM or the SIB base (r8 up to r15 use this).

	if byteCount == 8 {
		w = 1
	}

	if registerTo.BaseCodeOffset >= 8 {
		b = 1
	}

	if registerFrom.BaseCodeOffset >= 8 {
		r = 1
	}

	if w != 0 || r != 0 || b != 0 || x != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, r, x, b))
	}

	// Base code
	a.WriteBytes(baseCode)

	// ModRM
	hasOffset := offset != 0

	// rbp and r13 always have an offset
	if registerNameTo == "rbp" || registerNameTo == "r13" {
		hasOffset = true
	}

	if hasOffset {
		a.WriteBytes(opcode.ModRM(0b01, registerTo.BaseCodeOffset%8, registerFrom.BaseCodeOffset%8))
	} else {
		a.WriteBytes(opcode.ModRM(0b00, registerTo.BaseCodeOffset%8, registerFrom.BaseCodeOffset%8))
	}

	// rsp always need an SIB byte
	if registerNameTo == "rsp" || registerNameTo == "esp" || registerNameTo == "sp" || registerNameTo == "spl" {
		a.WriteBytes(opcode.SIB(0b00, 0b100, 0b100))
	}

	// r12 always need an SIB byte
	if registerNameTo == "r12" || registerNameTo == "r12d" || registerNameTo == "r12w" || registerNameTo == "r12b" {
		a.WriteBytes(opcode.SIB(0b00, 0b100, 0b100))
	}

	if hasOffset {
		a.WriteBytes(offset)
	}
}
