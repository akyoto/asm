package asm

// CompareRegisterNumber compares a register with a number.
func (a *Assembler) CompareRegisterNumber(registerName string, number uint64) {
	a.numberToRegisterSimple(0x83, 0x80, 0x3c, 0b111, false, registerName, number)
}
