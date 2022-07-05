package asm

// DivRegister divides the value in rax (usually) by the value in the specified register.
// This is a signed division.
// Quotient: AL AX EAX RAX
// Remainder: AH DX EDX RDX
func (a *Assembler) DivRegister(registerName string) {
	a.singleRegisterWithModRM(0xf7, 0xf6, registerName)
}
