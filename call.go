package asm

import "encoding/binary"

// Call places the return address on the top of the stack and continues
// program flow at the new address. The address is relative to the next instruction.
func (a *Assembler) Call(label string) {
	a.WriteBytes(0xe8)
	pointerPosition := a.Len()
	absoluteAddress, exists := a.Labels[label]

	if !exists {
		a.undefinedJumpLabels[label] = append(a.undefinedJumpLabels[label], jumpPointer{pointerPosition, 4})
		a.WriteBytes(0, 0, 0, 0)
		return
	}

	relativeAddress := int32(absoluteAddress - (pointerPosition + 4))
	_ = binary.Write(a, binary.LittleEndian, relativeAddress)
}

// Return transfers program control to a return address located on the top of the stack.
// The address is usually placed on the stack by a Call instruction.
func (a *Assembler) Return() {
	a.WriteBytes(0xc3)
}
