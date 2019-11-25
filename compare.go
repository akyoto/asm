package asm

// CompareRegisterNumber compares a register with a number.
func (a *Assembler) CompareRegisterNumber(registerName string, number uint64) {
	a.numberToRegisterSimple(0x83, 0x80, 0x3c, 0b111, false, registerName, number)
}

// CompareRegisterRegister compares a register with a register.
func (a *Assembler) CompareRegisterRegister(registerNameA string, registerNameB string) {
	a.registerToRegister([]byte{0x39}, []byte{0x38}, registerNameA, registerNameB, false)
}
