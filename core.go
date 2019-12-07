package asm

import (
	"log"
	"math"

	"github.com/akyoto/asm/opcode"
)

// bitsNeeded tells you how many bits are needed to encode this number.
func bitsNeeded(number int64) int {
	switch {
	case number >= math.MinInt8 && number <= math.MaxInt8:
		return 8

	case number >= math.MinInt16 && number <= math.MaxInt16:
		return 16

	case number >= math.MinInt32 && number <= math.MaxInt32:
		return 32

	default:
		return 64
	}
}

// registerToRegister encodes an instruction that takes two register parameters.
// baseCode is used for most cases except for single-byte register where oneByteCode is used.
func (a *Assembler) registerToRegister(baseCode []byte, oneByteCode []byte, registerNameTo string, registerNameFrom string, reverseModRM bool) {
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

	switch registerTo.BitSize {
	case 8:
		baseCode = oneByteCode

	case 16:
		a.WriteBytes(0x66)

	case 64:
		w = 1
	}

	if r != 0 || b != 0 || w != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, r, 0, b))
	}

	_, _ = a.Write(baseCode)

	if reverseModRM {
		a.WriteBytes(opcode.ModRM(0b11, registerTo.BaseCodeOffset%8, registerFrom.BaseCodeOffset%8))
	} else {
		a.WriteBytes(opcode.ModRM(0b11, registerFrom.BaseCodeOffset%8, registerTo.BaseCodeOffset%8))
	}
}

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
