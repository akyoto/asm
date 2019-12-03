package asm

// SignExtendToDX doubles the size of the register by sign-extending to DX/EDX/RDX.
func (a *Assembler) SignExtendToDX(registerName string) {
	switch registerName {
	case "ax":
		a.WriteBytes(0x66, 0x99)

	case "eax":
		a.WriteBytes(0x99)

	case "rax":
		a.WriteBytes(0x48, 0x99)

	default:
		panic("Not implemented")
	}
}
