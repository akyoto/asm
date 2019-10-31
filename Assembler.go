package asm

import (
	"bytes"

	"github.com/akyoto/asm/sections"
)

type Assembler struct {
	bytes.Buffer
	Strings        *sections.Strings
	StringPointers []sections.Pointer
}

func New() *Assembler {
	return &Assembler{
		Strings: sections.NewStrings(),
	}
}

func (a *Assembler) AddString(msg string) int64 {
	address := a.Strings.Add(msg)

	a.StringPointers = append(a.StringPointers, sections.Pointer{
		Address:  address,
		Position: a.Len(),
	})

	return address
}

func (a *Assembler) WriteBytes(someBytes ...byte) {
	for _, b := range someBytes {
		a.WriteByte(b)
	}
}
