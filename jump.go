package asm

import (
	"encoding/binary"
	"math"
)

// Mnemonic        Condition tested      Description
// ----------------------------------------------------------------------------
// jo              OF = 1                overflow
// jno             OF = 0                not overflow
// jc, jb, jnae    CF = 1                carry / below / not above nor equal
// jnc, jae, jnb   CF = 0                not carry / above or equal / not below
// je, jz          ZF = 1                equal / zero
// jne, jnz        ZF = 0                not equal / not zero
// jbe, jna        CF or ZF = 1          below or equal / not above
// ja, jnbe        CF or ZF = 0          above / not below or equal
// js              SF = 1                sign
// jns             SF = 0                not sign
// jp, jpe         PF = 1                parity / parity even
// jnp, jpo        PF = 0                not parity / parity odd
// jl, jnge        SF xor OF = 1         less / not greater nor equal
// jge, jnl        SF xor OF = 0         greater or equal / not less
// jle, jng        (SF xor OF) or ZF = 1 less or equal / not greater
// jg, jnle        (SF xor OF) or ZF = 0 greater / not less nor equal

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
	a.jump(0xeb, []byte{0xe9}, label)
}

// JumpIfLess jumps if the result was less.
func (a *Assembler) JumpIfLess(label string) {
	a.jump(0x7c, []byte{0x0f, 0x8c}, label)
}

// jump implements program flow jumps.
func (a *Assembler) jump(shortCode byte, nearCode []byte, label string) {
	instructionPosition := a.Len()
	pointerPosition := instructionPosition + 1
	pointerSize := uint8(1)
	absoluteAddress, exists := a.Labels[label]

	if !exists {
		// TODO: Support 32-bit jumps for unknown labels
		pointer := jumpPointer{pointerPosition, pointerSize}
		a.undefinedJumpLabels[label] = append(a.undefinedJumpLabels[label], pointer)
		a.WriteBytes(shortCode)
		a.WriteBytes(0)
		return
	}

	offset := int32(absoluteAddress - (pointerPosition + uint32(pointerSize)))

	// Near jump (32-bit)
	if offset < math.MinInt8 || offset > math.MaxInt8 {
		a.WriteBytes(nearCode...)
		_ = binary.Write(a, binary.LittleEndian, offset)
		return
	}

	// Short jump (8-bit)
	a.WriteBytes(shortCode)
	a.WriteBytes(byte(offset))
}

// Return transfers program control to a return address located on the top of the stack.
// The address is usually placed on the stack by a Call instruction.
func (a *Assembler) Return() {
	a.WriteBytes(0xc3)
}
