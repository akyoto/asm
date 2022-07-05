package asm

import (
	"encoding/binary"
	"fmt"
)

// Assembler implements machine-code encoding.
type Assembler struct {
	code            []byte
	data            []byte
	labels          map[string]Address
	jumps           []jump
	pointers        []Pointer
	EnableOptimizer bool
}

// New creates a new assembler.
func New() *Assembler {
	a := &Assembler{}
	a.Reset()
	return a
}

// Reset deletes the entire contents of the assembler.
func (a *Assembler) Reset() {
	a.code = a.code[:0]
	a.data = a.data[:0]
	a.labels = map[string]Address{}
	a.jumps = a.jumps[:0]
	a.pointers = a.pointers[:0]
	a.EnableOptimizer = true
}

// Code returns the machine code.
func (a *Assembler) Code() []byte {
	return a.code
}

// Data returns the data that is needed for the code to run.
func (a *Assembler) Data() []byte {
	return a.data
}

// Pointers returns the list of data references in the code.
func (a *Assembler) Pointers() []Pointer {
	return a.pointers
}

// Position returns the current address.
func (a *Assembler) Position() Address {
	return Address(len(a.code))
}

// Write writes the given byte slice to the machine code.
func (a *Assembler) Write(code []byte) (int, error) {
	a.code = append(a.code, code...)
	return len(code), nil
}

// WriteBytes writes the given bytes to the machine code.
func (a *Assembler) WriteBytes(someBytes ...byte) {
	a.code = append(a.code, someBytes...)
}

// WriteUint16 writes an unsigned 16-bit integer in little endian format to the machine code.
func (a *Assembler) WriteUint16(number uint16) {
	code := make([]byte, 2)
	binary.LittleEndian.PutUint16(code, number)
	a.code = append(a.code, code...)
}

// WriteUint32 writes an unsigned 32-bit integer in little endian format to the machine code.
func (a *Assembler) WriteUint32(number uint32) {
	code := make([]byte, 4)
	binary.LittleEndian.PutUint32(code, number)
	a.code = append(a.code, code...)
}

// WriteUint64 writes an unsigned 64-bit integer in little endian format to the machine code.
func (a *Assembler) WriteUint64(number uint64) {
	code := make([]byte, 8)
	binary.LittleEndian.PutUint64(code, number)
	a.code = append(a.code, code...)
}

// AddData writes the given byte slice to the data segment and returns the address.
func (a *Assembler) AddData(data []byte) Address {
	address := Address(len(a.data))
	a.data = append(a.data, data...)
	return address
}

// Compile compiles the code.
func (a *Assembler) Compile() error {
	for _, jump := range a.jumps {
		toAddress, exists := a.labels[jump.toLabel]

		if !exists {
			return fmt.Errorf("Undefined label: %s", jump.toLabel)
		}

		nextInstructionAddress := jump.addressPosition + 4
		distance := int32(toAddress) - int32(nextInstructionAddress)
		slice := a.code[jump.addressPosition:nextInstructionAddress]
		binary.LittleEndian.PutUint32(slice, uint32(distance))
	}

	return nil
}

// Merge combines the contents of the assembler with another one.
func (a *Assembler) Merge(b *Assembler) {
	codeOffset := uint32(len(a.code))
	dataOffset := uint32(len(a.data))

	// Add code and data
	a.code = append(a.code, b.code...)
	a.data = append(a.data, b.data...)

	// Add labels
	for name, address := range b.labels {
		a.AddLabelAt(name, codeOffset+address)
	}

	// Add jumps
	for _, jmp := range b.jumps {
		newJump := jump{
			addressPosition: codeOffset + jmp.addressPosition,
			toLabel:         jmp.toLabel,
			shortCode:       jmp.shortCode,
			nearCode:        jmp.nearCode,
		}

		a.jumps = append(a.jumps, newJump)
	}

	// Add pointers
	for _, pointer := range b.pointers {
		newPointer := Pointer{
			Address:  dataOffset + pointer.Address,
			Position: codeOffset + pointer.Position,
		}

		slice := a.code[newPointer.Position : newPointer.Position+4]
		binary.LittleEndian.PutUint32(slice, newPointer.Address)
		a.pointers = append(a.pointers, newPointer)
	}
}
