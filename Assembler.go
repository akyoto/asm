package asm

import (
	"encoding/binary"
	"fmt"

	"github.com/akyoto/asm/sections"
)

// Assembler implements machine-code encoding.
type Assembler struct {
	EnableOptimizer     bool
	Labels              map[string]sections.Address
	Strings             *sections.Strings
	StringPointers      []sections.Pointer
	code                []byte
	undefinedJumpLabels map[string][]jumpPointer
}

func New() *Assembler {
	a := &Assembler{}
	a.Reset()
	return a
}

func (a *Assembler) Merge(b *Assembler) {
	offset := a.Len()

	// Add code
	a.code = append(a.code, b.code...)

	// Add labels
	for name, address := range b.Labels {
		a.AddLabelAt(name, offset+address)
	}

	// Add strings
	sectionOffset := a.Strings.Len()
	a.Strings.Merge(b.Strings)

	for _, pointer := range b.StringPointers {
		newPointer := sections.Pointer{
			Address:  sectionOffset + pointer.Address,
			Position: offset + pointer.Position,
		}

		slice := a.code[newPointer.Position : newPointer.Position+4]
		binary.LittleEndian.PutUint32(slice, newPointer.Address)
		a.StringPointers = append(a.StringPointers, newPointer)
	}

	// Copy the undefined label only if "a" does not have the label
	for name, pointerList := range b.undefinedJumpLabels {
		address, exists := a.Labels[name]

		if exists {
			for _, pointer := range pointerList {
				pointer.Position += offset
				slice := a.code[pointer.Position : pointer.Position+uint32(pointer.Size)]
				binary.LittleEndian.PutUint32(slice, address-(pointer.Position+uint32(pointer.Size)))
			}
		} else {
			for index := range pointerList {
				pointerList[index].Position += offset
			}

			a.undefinedJumpLabels[name] = pointerList
		}
	}
}

func (a *Assembler) Reset() {
	a.EnableOptimizer = true
	a.Strings = sections.NewStrings()
	a.StringPointers = a.StringPointers[:0]
	a.Labels = map[string]sections.Address{}
	a.code = a.code[:0]
	a.undefinedJumpLabels = map[string][]jumpPointer{}
}

func (a *Assembler) AddString(msg string) sections.Address {
	address := a.Strings.Add(msg)

	a.StringPointers = append(a.StringPointers, sections.Pointer{
		Address:  address,
		Position: a.Len(),
	})

	return address
}

func (a *Assembler) Write(code []byte) (int, error) {
	a.code = append(a.code, code...)
	return len(code), nil
}

func (a *Assembler) WriteUint16(number uint16) {
	code := make([]byte, 2)
	binary.LittleEndian.PutUint16(code, number)
	a.code = append(a.code, code...)
}

func (a *Assembler) WriteUint32(number uint32) {
	code := make([]byte, 4)
	binary.LittleEndian.PutUint32(code, number)
	a.code = append(a.code, code...)
}

func (a *Assembler) WriteUint64(number uint64) {
	code := make([]byte, 8)
	binary.LittleEndian.PutUint64(code, number)
	a.code = append(a.code, code...)
}

func (a *Assembler) WriteBytes(someBytes ...byte) {
	a.code = append(a.code, someBytes...)
}

func (a *Assembler) Len() uint32 {
	return uint32(len(a.code))
}

func (a *Assembler) Verify() []error {
	if len(a.undefinedJumpLabels) == 0 {
		return nil
	}

	errors := make([]error, 0, len(a.undefinedJumpLabels))

	for label := range a.undefinedJumpLabels {
		errors = append(errors, fmt.Errorf("Undefined label: %s", label))
	}

	return errors
}

func (a *Assembler) Bytes() []byte {
	return a.code
}
