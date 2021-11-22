package asm

import (
	"github.com/akyoto/asm/sections"
)

var (
	moveRegisterNumber = numberToRegisterEncoder{
		baseCode:            0xb8,
		oneByteCode:         0xb0,
		reg:                 0b000,
		useNumberSize:       false,
		supports64BitNumber: true,
		useBaseCodeOffset:   true,
	}

	moveRegisterRegister = registerToRegisterEncoder{
		baseCode:    []byte{0x89},
		oneByteCode: []byte{0x88},
	}
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

// MoveRegisterRegister moves a register value into another register.
func (a *Assembler) MoveRegisterRegister(registerNameTo string, registerNameFrom string) {
	a.registerToRegister(&moveRegisterRegister, registerNameTo, registerNameFrom)
}
