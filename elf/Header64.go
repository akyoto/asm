package elf

// Header64Size is equal to the size of the ELF header in bytes.
const Header64Size = 64

// Header contains general information.
type Header64 struct {
	Magic                       [4]byte
	Class                       byte
	Endianness                  byte
	Version                     byte
	OSABI                       byte
	ABIVersion                  byte
	_                           [7]byte
	Type                        int16
	Architecture                int16
	FileVersion                 int32
	EntryPointInMemory          int64
	ProgramHeaderOffset         int64
	SectionHeaderOffset         int64
	Flags                       int32
	Size                        int16
	ProgramHeaderEntrySize      int16
	ProgramHeaderEntryCount     int16
	SectionHeaderEntrySize      int16
	SectionHeaderEntryCount     int16
	SectionNameStringTableIndex int16
}
