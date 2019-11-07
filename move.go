package asm

import (
	"encoding/binary"
	"log"
	"math"

	"github.com/akyoto/asm/opcode"
	"github.com/akyoto/asm/sections"
)

// MoveRegisterAddress moves an address into the given register.
func (a *Assembler) MoveRegisterAddress(registerNameTo string, address uint32) {
	addressPosition := a.MoveRegisterNumber(registerNameTo, uint64(address))

	a.StringPointers = append(a.StringPointers, sections.Pointer{
		Address:  address,
		Position: addressPosition,
	})
}

// MoveRegisterNumber moves a number into the given register.
func (a *Assembler) MoveRegisterNumber(registerNameTo string, number uint64) uint32 {
	baseCode := byte(0xb8)
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	if registerTo.BitSize == 8 {
		baseCode = 0xb0
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

// MoveRegisterRegister moves a register value into another register.
func (a *Assembler) MoveRegisterRegister(registerNameTo string, registerNameFrom string) {
	baseCode := byte(0x89)
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
		baseCode = 0x88

	case 16:
		a.WriteBytes(0x66)

	case 64:
		w = 1
	}

	if r != 0 || b != 0 || w != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, r, 0, b))
	}

	a.WriteBytes(baseCode)
	a.WriteBytes(opcode.ModRM(0b11, registerFrom.BaseCodeOffset%8, registerTo.BaseCodeOffset%8))
}
