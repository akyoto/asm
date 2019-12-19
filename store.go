package asm

import (
	"log"

	"github.com/akyoto/asm/opcode"
)

// StoreNumber stores a number into the memory address included in the given register.
func (a *Assembler) StoreNumber(registerNameTo string, offset byte, byteCount byte, number uint64) {
	a.store(0xc7, 0xc6, registerNameTo, offset, byteCount, "")

	// Number
	switch byteCount {
	case 8, 4:
		a.WriteUint32(uint32(number))

	case 2:
		a.WriteUint16(uint16(number))

	case 1:
		a.WriteBytes(byte(number))
	}
}

// StoreRegister stores the contents of a register into the memory address included in the given register.
func (a *Assembler) StoreRegister(registerNameTo string, offset byte, byteCount byte, registerNameFrom string) {
	a.store(0x89, 0x88, registerNameTo, offset, byteCount, registerNameFrom)
}

// store is the core function for memory store instructions.
func (a *Assembler) store(baseCode byte, oneByteCode byte, registerNameTo string, offset byte, byteCount byte, registerNameFrom string) {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	registerFrom := registers[registerNameFrom]

	if registerFrom != nil && registerFrom.MustHaveREX {
		a.WriteBytes(0x40)
	}

	switch byteCount {
	case 2:
		a.WriteBytes(0x66)

	case 1:
		baseCode = oneByteCode
	}

	// REX prefix
	w := byte(0) // Indicates a 64-bit register.
	r := byte(0) // Extension to the "reg" field in ModRM.
	x := byte(0) // Extension to the SIB index field.
	b := byte(0) // Extension to the "rm" field in ModRM or the SIB base (r8 up to r15 use this).

	if byteCount == 8 {
		w = 1
	}

	if registerFrom != nil && registerFrom.BaseCodeOffset >= 8 {
		r = 1
	}

	if registerTo.BaseCodeOffset >= 8 {
		b = 1
	}

	if w != 0 || b != 0 || x != 0 || registerTo.MustHaveREX {
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

	reg := byte(0)

	if registerFrom != nil {
		reg = registerFrom.BaseCodeOffset % 8
	}

	if hasOffset {
		a.WriteBytes(opcode.ModRM(0b01, reg, registerTo.BaseCodeOffset%8))
	} else {
		a.WriteBytes(opcode.ModRM(0b00, reg, registerTo.BaseCodeOffset%8))
	}

	// rsp and r12 always need an SIB byte
	if registerNameTo == "rsp" || registerNameTo == "r12" {
		a.WriteBytes(opcode.SIB(0b00, 0b100, 0b100))
	}

	if hasOffset {
		a.WriteBytes(offset)
	}
}
