package asm

var compareRegisterNumber = numberToRegisterEncoder{
	baseCode:            0x81,
	oneByteCode:         0x80,
	reg:                 0b111,
	regEqualsRM:         false,
	useNumberSize:       false,
	supports64BitNumber: false,
	useBaseCodeOffset:   false,
}

var compareRegisterNumber1B = numberToRegisterEncoder{
	baseCode:            0x83,
	oneByteCode:         compareRegisterNumber.oneByteCode,
	reg:                 compareRegisterNumber.reg,
	regEqualsRM:         compareRegisterNumber.regEqualsRM,
	useNumberSize:       true,
	supports64BitNumber: compareRegisterNumber.supports64BitNumber,
	useBaseCodeOffset:   compareRegisterNumber.useBaseCodeOffset,
}

// CompareRegisterNumber compares a register with a number.
func (a *Assembler) CompareRegisterNumber(registerName string, number uint64) {
	if registerName == "al" {
		a.WriteBytes(0x3c, byte(number))
		return
	}

	operandBitSize := bitsNeeded(int64(number))

	if registerName == "rax" && operandBitSize >= 16 && operandBitSize <= 32 {
		a.WriteBytes(0x48, 0x3d)
		a.WriteUint32(uint32(number))
		return
	}

	if registerName == "eax" && operandBitSize >= 16 && operandBitSize <= 32 {
		a.WriteBytes(0x3d)
		a.WriteUint32(uint32(number))
		return
	}

	if registerName == "ax" && operandBitSize == 16 {
		a.WriteBytes(0x66, 0x3d)
		a.WriteUint16(uint16(number))
		return
	}

	encoder := &compareRegisterNumber

	if operandBitSize == 8 {
		encoder = &compareRegisterNumber1B
	}

	a.numberToRegister(encoder, registerName, number)
}

// CompareRegisterRegister compares a register with a register.
func (a *Assembler) CompareRegisterRegister(registerNameA string, registerNameB string) {
	a.registerToRegister([]byte{0x39}, []byte{0x38}, registerNameA, registerNameB, false)
}
