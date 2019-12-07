package asm

import (
	"log"

	"github.com/akyoto/asm/opcode"
)

// singleRegisterWithModRM encodes an instruction that takes one register parameter encoded via ModRM.
func (a *Assembler) singleRegisterWithModRM(baseCode byte, oneByteCode byte, registerNameTo string) {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	w := byte(0)
	b := byte(0)

	if registerTo.BaseCodeOffset >= 8 {
		b = 1
	}

	switch registerTo.BitSize {
	case 8:
		baseCode = oneByteCode

	case 16:
		a.WriteBytes(0x66)

	case 64:
		w = 1
	}

	if b != 0 || w != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, 0, 0, b))
	}

	a.WriteBytes(baseCode)
	a.WriteBytes(opcode.ModRM(0b11, 0b111, registerTo.BaseCodeOffset%8))
}
