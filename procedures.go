package asm

import (
	"encoding/binary"
)

// AddLabelAt adds a label for the current instruction address.
func (a *Assembler) AddLabel(name string) {
	a.AddLabelAt(name, a.Len())
}

// AddLabelAt adds a label for the given address.
func (a *Assembler) AddLabelAt(name string, address uint32) {
	a.Labels[name] = address

	// Fix all references to previously undefined labels
	for _, pointer := range a.undefinedJumpLabels[name] {
		slice := a.code[pointer.Position : pointer.Position+uint32(pointer.Size)]
		offset := address - (pointer.Position + uint32(pointer.Size))

		switch pointer.Size {
		case 4:
			binary.LittleEndian.PutUint32(slice, offset)

		case 2:
			binary.LittleEndian.PutUint16(slice, uint16(offset))

		case 1:
			slice[0] = byte(offset)
		}
	}

	delete(a.undefinedJumpLabels, name)
}

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

// Jump continues program flow at the new address.
// The address is relative to the next instruction.
func (a *Assembler) Jump(label string) {
	// TODO: Make this work for any type of jump (currently 1-byte only).
	a.WriteBytes(0xeb)
	pointerPosition := a.Len()
	absoluteAddress, exists := a.Labels[label]

	if !exists {
		a.undefinedJumpLabels[label] = append(a.undefinedJumpLabels[label], jumpPointer{pointerPosition, 1})
		a.WriteBytes(0)
		return
	}

	relativeAddress := int8(absoluteAddress - (pointerPosition + 1))
	_ = binary.Write(a, binary.LittleEndian, relativeAddress)
}

// Return transfers program control to a return address located on the top of the stack.
// The address is usually placed on the stack by a Call instruction.
func (a *Assembler) Return() {
	a.WriteBytes(0xc3)
}
