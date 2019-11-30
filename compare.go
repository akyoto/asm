package asm

// CompareRegisterNumber compares a register with a number.
func (a *Assembler) CompareRegisterNumber(registerName string, number uint64) {
	if registerName == "al" {
		a.WriteBytes(0x3c, byte(number))
		return
	}

	operandBitSize := bitsNeeded(number)

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

	baseCode := byte(0x81)
	useNumberSize := false

	if bitsNeeded(number) == 8 {
		baseCode = 0x83
		useNumberSize = true
	}

	a.numberToRegister(baseCode, 0x80, 0b111, false, useNumberSize, false, false, registerName, number)
}

// CompareRegisterRegister compares a register with a register.
func (a *Assembler) CompareRegisterRegister(registerNameA string, registerNameB string) {
	a.registerToRegister([]byte{0x39}, []byte{0x38}, registerNameA, registerNameB, false)
}
