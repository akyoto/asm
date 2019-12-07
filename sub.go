package asm

var subRegisterNumber = numberToRegisterEncoder{
	baseCode:            0x83,
	oneByteCode:         0x80,
	reg:                 0b101,
	regEqualsRM:         false,
	useNumberSize:       true,
	supports64BitNumber: false,
	useBaseCodeOffset:   false,
}

// SubRegisterNumber subtracts a number from a register.
func (a *Assembler) SubRegisterNumber(registerNameTo string, number uint64) {
	if registerNameTo == "al" {
		a.WriteBytes(0x2c, byte(number))
		return
	}

	a.numberToRegister(&subRegisterNumber, registerNameTo, number)
}

// SubRegisterRegister subtracts a register value from another register.
func (a *Assembler) SubRegisterRegister(registerNameTo string, registerNameFrom string) {
	a.registerToRegister([]byte{0x29}, []byte{0x28}, registerNameTo, registerNameFrom, false)
}
