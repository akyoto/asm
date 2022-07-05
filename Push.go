package asm

import (
	"log"

	"github.com/akyoto/asm/opcode"
)

// PushRegister pushes the value inside the register onto the stack.
func (a *Assembler) PushRegister(registerName string) {
	a.pushPopRegister(0x50, registerName)
}

// PopRegister pops a value from the stack and saves it into the register.
func (a *Assembler) PopRegister(registerName string) {
	a.pushPopRegister(0x58, registerName)
}

// pushPopRegister encodes a push/pop instruction that only needs a register name.
func (a *Assembler) pushPopRegister(baseCode byte, registerNameTo string) {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	if registerTo.BitSize == 16 {
		a.WriteBytes(0x66)
	}

	if registerTo.BaseCodeOffset >= 8 {
		a.WriteBytes(opcode.REX(0, 0, 0, 1))
	}

	a.WriteBytes(baseCode + registerTo.BaseCodeOffset%8)
}
