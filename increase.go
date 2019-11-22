package asm

import (
	"log"

	"github.com/akyoto/asm/opcode"
)

// IncreaseRegister increases the register value by 1.
func (a *Assembler) IncreaseRegister(registerName string) {
	a.incDecRegister(0xff, 0xfe, 0, registerName)
}

// DecreaseRegister decreases the register value by 1.
func (a *Assembler) DecreaseRegister(registerName string) {
	a.incDecRegister(0xff, 0xfe, 1, registerName)
}

// incDecRegister encodes an inc/dec instruction that only needs a register name.
func (a *Assembler) incDecRegister(baseCode byte, oneByteCode byte, reg byte, registerNameTo string) {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	if registerTo.BitSize == 8 {
		baseCode = oneByteCode
	}

	if registerTo.BitSize == 16 {
		a.WriteBytes(0x66)
	}

	bitSize := registerTo.BitSize

	// 64-bit register
	w := byte(0)

	if bitSize == 64 {
		w = 1
	}

	// Are we accessing any of the 64-bit only registers (r8 up to r15)?
	b := byte(0)

	if registerTo.BaseCodeOffset >= 8 {
		b = 1
	}

	// REX
	if w != 0 || b != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, 0, 0, b))
	}

	// Base code
	a.WriteBytes(baseCode)

	a.WriteBytes(opcode.ModRM(0b11, reg, registerTo.BaseCodeOffset%8))
}
