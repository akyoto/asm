package asm

import (
	"fmt"
	"log"

	"github.com/akyoto/asm/opcode"
)

type numberToRegisterEncoder struct {
	baseCode            byte
	oneByteCode         byte
	reg                 byte
	regEqualsRM         bool
	useNumberSize       bool
	supports64BitNumber bool
	useBaseCodeOffset   bool
}

// numberToRegister encodes an instruction with a register and a number parameter.
func (a *Assembler) numberToRegister(encoder *numberToRegisterEncoder, registerNameTo string, number uint64) uint32 {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	// We start with the assumption that the base code will be the default.
	baseCode := encoder.baseCode

	// If the target register is only 8 bits long,
	// most instructions use a different base code.
	if registerTo.BitSize == 8 {
		baseCode = encoder.oneByteCode
	}

	// If the target register is 16 bits long,
	// the base code is always prefixed with 0x66.
	if registerTo.BitSize == 16 {
		a.WriteBytes(0x66)
	}

	operandBitSize := bitsNeeded(int64(number))
	registerBitSize := registerTo.BitSize

	// Change 64-bit target registers like "rax" to "eax"
	// when the number doesn't need the full 64 bits
	// because the register will be zero-extended in 32-bit mode
	// and the instruction takes less bytes.
	if a.EnableOptimizer && encoder.supports64BitNumber && registerBitSize == 64 && operandBitSize < 64 {
		registerBitSize = 32
	}

	if operandBitSize > registerBitSize {
		panic(fmt.Errorf("Operand '%v' (%d bits) doesn't fit into register %s (%d bits)", number, operandBitSize, registerNameTo, registerBitSize))
	}

	// Indicates a 64-bit register.
	w := byte(0)

	if registerBitSize == 64 {
		w = 1
	}

	// Extension to the "reg" field in ModRM.
	r := byte(0)

	if encoder.regEqualsRM && registerTo.BaseCodeOffset >= 8 {
		r = 1
	}

	// Extension to the SIB index field.
	x := byte(0)

	// Are we accessing any of the 64-bit only registers (r8 up to r15)?
	// This is an extension to the "rm" field in ModRM or the SIB base.
	b := byte(0)

	if registerTo.BaseCodeOffset >= 8 {
		b = 1
	}

	// REX
	if w != 0 || b != 0 || x != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, r, x, b))
	}

	// If the encoding uses a base code offset,
	// encode the target register in the base code.
	if encoder.useBaseCodeOffset {
		baseCode += registerTo.BaseCodeOffset % 8
	}

	// Base code
	a.WriteBytes(baseCode)

	// If the encoding doesn't use a base code offset,
	// the target register is specified in the rm part
	// of the ModRM byte.
	if !encoder.useBaseCodeOffset {
		reg := encoder.reg
		rm := registerTo.BaseCodeOffset % 8

		if encoder.regEqualsRM {
			reg = rm
		}

		a.WriteBytes(opcode.ModRM(0b11, reg, rm))
	}

	bitSize := registerBitSize

	if encoder.useNumberSize {
		bitSize = operandBitSize
	}

	// If the number needs 64 bit but the instruction doesn't
	// support 64-bit immediate values, set the bit size to 32.
	if !encoder.supports64BitNumber && bitSize == 64 {
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
