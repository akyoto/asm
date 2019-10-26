package elf

// SectionHeader64Size is equal to the size of a section header in bytes.
const SectionHeader64Size = 64

// SectionHeader points to the data sections of our program.
type SectionHeader64 struct {
	NameOffset      int32
	Type            int32
	Flags           int64
	VirtualAddress  int64
	Offset          int64
	SizeInFileImage int64
	Link            int32
	Info            int32
	Align           int64
	EntrySize       int64
}
