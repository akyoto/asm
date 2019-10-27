package elf

// ProgramHeader64Size is equal to the size of a program header in bytes.
const ProgramHeader64Size = 56

// ProgramHeader points to the executable part of our program.
type ProgramHeader64 struct {
	Type            ProgramType
	Flags           ProgramFlags
	Offset          int64
	VirtualAddress  int64
	PhysicalAddress int64
	SizeInFileImage int64
	SizeInMemory    int64
	Align           int64
}
