package asm

// MulRegisterNumber multiplies a register with a number.
func (a *Assembler) MulRegisterNumber(registerNameTo string, number uint64) {
	a.numberToRegister(0x6b, 0x6b, 0, true, true, false, false, registerNameTo, number)
}

// MulRegisterRegister multiplies a register with another register.
func (a *Assembler) MulRegisterRegister(registerNameTo string, registerNameFrom string) {
	a.registerToRegister([]byte{0x0f, 0xaf}, nil, registerNameTo, registerNameFrom, true)
}
