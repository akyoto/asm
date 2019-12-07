package asm

var mulRegisterNumber = numberToRegisterEncoder{
	baseCode:            0x6b,
	oneByteCode:         0x6b,
	reg:                 0b000,
	regEqualsRM:         true,
	useNumberSize:       true,
	supports64BitNumber: false,
	useBaseCodeOffset:   false,
}

// MulRegisterNumber multiplies a register with a number.
func (a *Assembler) MulRegisterNumber(registerNameTo string, number uint64) {
	a.numberToRegister(&mulRegisterNumber, registerNameTo, number)
}

// MulRegisterRegister multiplies a register with another register.
func (a *Assembler) MulRegisterRegister(registerNameTo string, registerNameFrom string) {
	a.registerToRegister([]byte{0x0f, 0xaf}, nil, registerNameTo, registerNameFrom, true)
}
