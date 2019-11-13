package asm

import (
	"encoding/binary"
	"log"
	"math"

	"github.com/akyoto/asm/opcode"
)

// numberToRegister encodes an instruction with a register and a number parameter.
func (a *Assembler) numberToRegister(baseCode byte, oneByteCode byte, registerNameTo string, number uint64) uint32 {
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

	operandBitSize := 0

	switch {
	case number <= math.MaxUint8:
		operandBitSize = 8

	case number <= math.MaxUint16:
		operandBitSize = 16

	case number <= math.MaxUint32:
		operandBitSize = 32

	default:
		operandBitSize = 64
	}

	registerBitSize := registerTo.BitSize

	if a.EnableOptimizer && registerBitSize == 64 && operandBitSize < 64 {
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
	a.WriteBytes(baseCode + registerTo.BaseCodeOffset%8)
	addressPosition := a.Len()

	// Number
	var buffer []byte

	switch bitSize {
	case 64:
		buffer = make([]byte, 8)
		binary.LittleEndian.PutUint64(buffer, number)

	case 32:
		buffer = make([]byte, 4)
		binary.LittleEndian.PutUint32(buffer, uint32(number))

	case 16:
		buffer = make([]byte, 2)
		binary.LittleEndian.PutUint16(buffer, uint16(number))

	case 8:
		buffer = []byte{byte(number)}
	}

	_, _ = a.Write(buffer)
	return addressPosition
}

// numberToRegisterSimple encodes an instruction with a register and a number parameter.
func (a *Assembler) numberToRegisterSimple(baseCode byte, oneByteCode byte, alCode byte, reg byte, registerNameTo string, number uint64) {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	if registerTo.BitSize == 8 {
		baseCode = oneByteCode

		// AL has a special instruction
		if registerNameTo == "al" {
			a.WriteBytes(alCode, byte(number))
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
	a.WriteBytes(opcode.ModRM(0b11, reg, registerTo.BaseCodeOffset%8))

	// Number
	buffer := []byte{byte(number)}
	_, _ = a.Write(buffer)
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
