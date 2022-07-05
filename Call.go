package asm

// Call places the return address on the top of the stack and continues
// program flow at the new address. The address is relative to the next instruction.
func (a *Assembler) Call(label string) {
	a.jumps = append(a.jumps, jump{
		addressPosition: a.Position() + 1,
		toLabel:         label,
		nearCode:        []byte{0xe8},
	})

	a.WriteBytes(0xe8)
	a.WriteBytes(0, 0, 0, 0)
}

// Return transfers program control to a return address located on the top of the stack.
// The address is usually placed on the stack by a Call instruction.
func (a *Assembler) Return() {
	a.WriteBytes(0xc3)
}
