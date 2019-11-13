package asm

// MulRegisterRegister multiplies a register with another register.
func (a *Assembler) MulRegisterRegister(registerNameTo string, registerNameFrom string) {
	a.registerToRegister([]byte{0x0f, 0xaf}, nil, registerNameTo, registerNameFrom, true)
}
