package asm

import (
	"log"

	"github.com/akyoto/asm/opcode"
)

type registerToRegisterEncoder struct {
	baseCode     []byte
	oneByteCode  []byte
	reverseModRM bool
}

// registerToRegister encodes an instruction that takes two register parameters.
// baseCode is used for most cases except for single-byte register where oneByteCode is used.
func (a *Assembler) registerToRegister(encoder *registerToRegisterEncoder, registerNameTo string, registerNameFrom string) {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	registerFrom, exists := registers[registerNameFrom]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameFrom)
	}

	// REX prefix
	w := byte(0) // Indicates a 64-bit register.
	r := byte(0) // Extension to the "reg" field in ModRM.
	x := byte(0) // Extension to the SIB index field.
	b := byte(0) // Extension to the "rm" field in ModRM or the SIB base (r8 up to r15 use this).

	if registerFrom.BaseCodeOffset >= 8 {
		r = 1
	}

	if registerTo.BaseCodeOffset >= 8 {
		b = 1
	}

	baseCode := encoder.baseCode

	switch registerTo.BitSize {
	case 8:
		baseCode = encoder.oneByteCode

	case 16:
		a.WriteBytes(0x66)

	case 64:
		w = 1
	}

	if w != 0 || r != 0 || b != 0 || x != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, r, x, b))
	}

	_, _ = a.Write(baseCode)

	if encoder.reverseModRM {
		a.WriteBytes(opcode.ModRM(0b11, registerTo.BaseCodeOffset%8, registerFrom.BaseCodeOffset%8))
	} else {
		a.WriteBytes(opcode.ModRM(0b11, registerFrom.BaseCodeOffset%8, registerTo.BaseCodeOffset%8))
	}
}
