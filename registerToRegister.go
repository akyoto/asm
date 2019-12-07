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

	w := byte(0)
	r := byte(0)
	b := byte(0)

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

	if w != 0 || r != 0 || b != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, r, 0, b))
	}

	_, _ = a.Write(baseCode)

	if encoder.reverseModRM {
		a.WriteBytes(opcode.ModRM(0b11, registerTo.BaseCodeOffset%8, registerFrom.BaseCodeOffset%8))
	} else {
		a.WriteBytes(opcode.ModRM(0b11, registerFrom.BaseCodeOffset%8, registerTo.BaseCodeOffset%8))
	}
}
