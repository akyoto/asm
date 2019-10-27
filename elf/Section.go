package elf

// Section includes the text data and initialized variables.
type Section struct {
	Header  SectionHeader64
	Padding []byte
	Data    []byte
}
