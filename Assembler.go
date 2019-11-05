package asm

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/akyoto/asm/sections"
)

type Assembler struct {
	EnableOptimizer     bool
	Labels              map[string]sections.Address
	Strings             *sections.Strings
	StringPointers      []sections.Pointer
	code                []byte
	undefinedCallLabels map[string][]sections.Address
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
	for name, addressList := range b.undefinedCallLabels {
		address, exists := a.Labels[name]

		if exists {
			for _, position := range addressList {
				position += offset
				slice := a.code[position : position+4]
				binary.LittleEndian.PutUint32(slice, address-(position+4))
			}
		} else {
			for index := range addressList {
				addressList[index] += offset
			}

			a.undefinedCallLabels[name] = addressList
		}
	}
}

func (a *Assembler) Reset() {
	a.EnableOptimizer = true
	a.Strings = sections.NewStrings()
	a.StringPointers = a.StringPointers[:0]
	a.Labels = map[string]sections.Address{}
	a.code = a.code[:0]
	a.undefinedCallLabels = map[string][]sections.Address{}
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

func (a *Assembler) WriteBytes(someBytes ...byte) {
	a.code = append(a.code, someBytes...)
}

func (a *Assembler) Len() uint32 {
	return uint32(len(a.code))
}

func (a *Assembler) Bytes() []byte {
	if len(a.undefinedCallLabels) > 0 {
		errorMessage := ""

		for label := range a.undefinedCallLabels {
			errorMessage += fmt.Sprintf("Undefined label: %s\n", label)
		}

		log.Fatal(errorMessage)
	}

	return a.code
}
