package asm

import "log"

// PushRegister pushes the value inside the register onto the stack.
func (a *Assembler) PushRegister(registerName string) {
	a.encodeRegister(0x50, registerName)
}

// PopRegister pops a value from the stack and saves it into the register.
func (a *Assembler) PopRegister(registerName string) {
	a.encodeRegister(0x58, registerName)
}

// encodeRegister encodes an instruction that only needs a register name.
func (a *Assembler) encodeRegister(baseCode byte, registerNameTo string) {
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	if registerTo.BaseCodeOffset >= 8 {
		a.WriteBytes(REX(0, 0, 0, 1))
	}

	a.WriteBytes(baseCode + registerTo.BaseCodeOffset%8)
}
