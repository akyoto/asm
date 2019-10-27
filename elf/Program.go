package elf

// Program includes the machine-code instructions.
type Program struct {
	Header  ProgramHeader64
	Padding []byte
	Data    []byte
}
