package asm

import (
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

func (a *Assembler) Reset() {
	a.EnableOptimizer = true
	a.Strings = sections.NewStrings()
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
