package asm

import (
	"log"

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
	return a.numberToRegister(&moveRegisterNumber, registerNameTo, number)
}

var moveRegisterNumber = numberToRegisterEncoder{
	baseCode:            0xb8,
	oneByteCode:         0xb0,
	reg:                 0b000,
	useNumberSize:       false,
	supports64BitNumber: true,
	useBaseCodeOffset:   true,
}

// MoveRegisterRegister moves a register value into another register.
func (a *Assembler) MoveRegisterRegister(registerNameTo string, registerNameFrom string) {
	a.registerToRegister(&moveRegisterRegister, registerNameTo, registerNameFrom)
}

var moveRegisterRegister = registerToRegisterEncoder{
	baseCode:    []byte{0x89},
	oneByteCode: []byte{0x88},
}

// MoveMemoryNumber moves a number into the memory address included in the given register.
func (a *Assembler) MoveMemoryNumber(registerNameTo string, byteCount int, number uint64) {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	baseCode := byte(0xc7)

	switch byteCount {
	case 2:
		a.WriteBytes(0x66)

	case 1:
		baseCode = 0xc6
	}

	// REX prefix
	w := byte(0) // Indicates a 64-bit register.
	r := byte(0) // Extension to the "reg" field in ModRM.
	x := byte(0) // Extension to the SIB index field.
	b := byte(0) // Extension to the "rm" field in ModRM or the SIB base (r8 up to r15 use this).

	if byteCount == 8 {
		w = 1
	}

	if registerTo.BaseCodeOffset >= 8 {
		b = 1
	}

	if w != 0 || b != 0 || x != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, r, x, b))
	}

	// Base code
	a.WriteBytes(baseCode)

	// ModRM
	if registerNameTo == "rsp" || registerNameTo == "r12" {
		a.WriteBytes(0x04, 0x24)
	} else if registerNameTo == "rbp" {
		a.WriteBytes(opcode.ModRM(0b01, 0, 0b101), 0)
	} else if registerNameTo == "r13" {
		a.WriteBytes(0x45, 0x00)
	} else {
		a.WriteBytes(opcode.ModRM(0b00, 0, registerTo.BaseCodeOffset%8))
	}

	// Number
	switch byteCount {
	case 8, 4:
		a.WriteUint32(uint32(number))

	case 2:
		a.WriteUint16(uint16(number))

	case 1:
		a.WriteBytes(byte(number))
	}
}
