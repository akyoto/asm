package asm

var (
	mulRegisterNumber = numberToRegisterEncoder{
		baseCode:            0x6b,
		oneByteCode:         0x6b,
		regRMEqual:          true,
		useNumberSize:       true,
		supports64BitNumber: false,
		useBaseCodeOffset:   false,
	}

	mulRegisterRegister = registerToRegisterEncoder{
		baseCode:     []byte{0x0f, 0xaf},
		oneByteCode:  nil,
		reverseModRM: true,
	}
)

// MulRegisterNumber multiplies a register with a number.
func (a *Assembler) MulRegisterNumber(registerNameTo string, number uint64) {
	a.numberToRegister(&mulRegisterNumber, registerNameTo, number)
}

// MulRegisterRegister multiplies a register with another register.
func (a *Assembler) MulRegisterRegister(registerNameTo string, registerNameFrom string) {
	a.registerToRegister(&mulRegisterRegister, registerNameTo, registerNameFrom)
}
