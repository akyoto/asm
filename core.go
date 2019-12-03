package asm

import (
	"log"
	"math"

	"github.com/akyoto/asm/opcode"
)

// bitsNeeded tells you how many bits are needed to encode this number.
func bitsNeeded(number uint64) int {
	switch {
	case number <= math.MaxUint8:
		return 8

	case number <= math.MaxUint16:
		return 16

	case number <= math.MaxUint32:
		return 32

	default:
		return 64
	}
}

// numberToRegister encodes an instruction with a register and a number parameter.
func (a *Assembler) numberToRegister(baseCode byte, oneByteCode byte, reg byte, regEqualsRM bool, useNumberSize bool, supports64BitNumber bool, useBaseCodeOffset bool, registerNameTo string, number uint64) uint32 {
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

	operandBitSize := bitsNeeded(number)
	registerBitSize := registerTo.BitSize

	if a.EnableOptimizer && supports64BitNumber && registerBitSize == 64 && operandBitSize < 64 {
		registerBitSize = 32
	}

	if operandBitSize > registerBitSize {
		log.Printf("Operand '%v' (%d bits) doesn't fit into register %s (%d bits)", number, operandBitSize, registerNameTo, registerBitSize)
	}

	bitSize := registerBitSize

	// 64-bit register
	w := byte(0)

	if bitSize == 64 {
		w = 1
	}

	r := byte(0)

	if regEqualsRM && registerTo.BaseCodeOffset >= 8 {
		r = 1
	}

	// Are we accessing any of the 64-bit only registers (r8 up to r15)?
	b := byte(0)

	if registerTo.BaseCodeOffset >= 8 {
		b = 1
	}

	// REX
	if w != 0 || b != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, r, 0, b))
	}

	if useBaseCodeOffset {
		baseCode += registerTo.BaseCodeOffset % 8
	}

	// Base code
	a.WriteBytes(baseCode)

	if !useBaseCodeOffset {
		rm := registerTo.BaseCodeOffset % 8

		if regEqualsRM {
			reg = rm
		}

		a.WriteBytes(opcode.ModRM(0b11, reg, rm))
	}

	if useNumberSize {
		bitSize = operandBitSize
	}

	if !supports64BitNumber && bitSize == 64 {
		bitSize = 32
	}

	numberPos := a.Len()

	// Number
	switch bitSize {
	case 64:
		a.WriteUint64(number)

	case 32:
		a.WriteUint32(uint32(number))

	case 16:
		a.WriteUint16(uint16(number))

	case 8:
		a.WriteBytes(byte(number))
	}

	return numberPos
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
