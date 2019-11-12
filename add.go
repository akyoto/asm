package asm

import (
	"log"

	"github.com/akyoto/asm/opcode"
)

// AddRegisterNumber adds a number to the given register.
func (a *Assembler) AddRegisterNumber(registerNameTo string, number uint64) {
	baseCode := byte(0x83)
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	if registerTo.BitSize == 8 {
		baseCode = 0x80

		// AL has a special instruction
		if registerNameTo == "al" {
			a.WriteBytes(0x04, byte(number))
			return
		}
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
	a.WriteBytes(opcode.ModRM(0b11, 0, registerTo.BaseCodeOffset%8))

	// Number
	buffer := []byte{byte(number)}
	_, _ = a.Write(buffer)
}

// AddRegisterRegister adds a register value into another register.
func (a *Assembler) AddRegisterRegister(registerNameTo string, registerNameFrom string) {
	a.registerToRegister(0x01, 0x00, registerNameTo, registerNameFrom)
}
