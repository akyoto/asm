package asm

import (
	"fmt"
	"log"

	"github.com/akyoto/asm/sections"
)

type Assembler struct {
	code                []byte
	Strings             *sections.Strings
	StringPointers      []sections.Pointer64
	Labels              map[string]int32
	undefinedCallLabels map[string][]int32
}

func New() *Assembler {
	return &Assembler{
		Strings:             sections.NewStrings(),
		Labels:              map[string]int32{},
		undefinedCallLabels: map[string][]int32{},
	}
}

func (a *Assembler) AddString(msg string) int64 {
	address := a.Strings.Add(msg)

	a.StringPointers = append(a.StringPointers, sections.Pointer64{
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

func (a *Assembler) Len() int32 {
	return int32(len(a.code))
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
