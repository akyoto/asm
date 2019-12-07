package asm

// AddRegisterNumber adds a number to the given register.
func (a *Assembler) AddRegisterNumber(registerNameTo string, number uint64) {
	if registerNameTo == "al" {
		a.WriteBytes(0x04, byte(number))
		return
	}

	a.numberToRegister(&addRegisterNumber, registerNameTo, number)
}

var addRegisterNumber = numberToRegisterEncoder{
	baseCode:            0x83,
	oneByteCode:         0x80,
	reg:                 0b000,
	useNumberSize:       true,
	supports64BitNumber: false,
	useBaseCodeOffset:   false,
}

// AddRegisterRegister adds a register value into another register.
func (a *Assembler) AddRegisterRegister(registerNameTo string, registerNameFrom string) {
	a.registerToRegister(&addRegisterRegister, registerNameTo, registerNameFrom)
}

var addRegisterRegister = registerToRegisterEncoder{
	baseCode:    []byte{0x01},
	oneByteCode: []byte{0x00},
}
