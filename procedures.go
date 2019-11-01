package asm

import (
	"encoding/binary"
)

// AddLabel adds a label for the current instruction address.
func (a *Assembler) AddLabel(name string) {
	address := a.Len()
	a.Labels[name] = address

	// Fix all references to previously undefined labels
	for _, position := range a.undefinedCallLabels[name] {
		slice := a.code[position : position+4]
		binary.LittleEndian.PutUint32(slice, uint32(address-(position+4)))
	}

	delete(a.undefinedCallLabels, name)
}

// Call places the return address on the top of the stack and continues
// program flow at the new address. The address is relative to the next instruction.
func (a *Assembler) Call(label string) {
	a.WriteBytes(0xe8)
	pointerPosition := a.Len()
	absoluteAddress, exists := a.Labels[label]

	if !exists {
		a.undefinedCallLabels[label] = append(a.undefinedCallLabels[label], pointerPosition)
		a.WriteBytes(0, 0, 0, 0)
		return
	}

	relativeAddress := absoluteAddress - (pointerPosition + 4)
	_ = binary.Write(a, binary.LittleEndian, relativeAddress)
}

// Return transfers program control to a return address located on the top of the stack.
// The address is usually placed on the stack by a Call instruction.
func (a *Assembler) Return() {
	a.WriteBytes(0xc3)
}
