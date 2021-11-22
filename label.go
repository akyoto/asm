package asm

import "encoding/binary"

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
