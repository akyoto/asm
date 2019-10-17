package elf

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"
)

const (
	Header64Size        = 64
	ProgramHeader64Size = 56
	SectionHeader64Size = 64
)

// ELF64 represents an ELF 64-bit file.
type ELF64 struct {
	Header64
	ProgramHeaders      []ProgramHeader64
	SectionHeaders      []SectionHeader64
	InstructionsPadding []byte
	Instructions        []byte
	SectionsPadding     []byte
	Sections            []byte
}

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

// ProgramHeader points to the executable part of our program.
type ProgramHeader64 struct {
	Type            int32
	Flags           int32
	Offset          int64
	VirtualAddress  int64
	PhysicalAddress int64
	SizeInFileImage int64
	SizeInMemory    int64
	Align           int64
}

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

// New creates a new 64-bit ELF binary.
func New(instructions []byte) *ELF64 {
	const (
		address      = 0x400000
		programAlign = 16
		sectionAlign = 16
	)

	program := &ELF64{
		Header64: Header64{
			Magic:                  [4]byte{0x7F, 'E', 'L', 'F'},
			Class:                  2,
			Endianness:             1, // Little endianness
			Version:                1,
			Type:                   0x02,
			Architecture:           0x3E, // x86-64
			FileVersion:            1,
			Size:                   Header64Size,
			ProgramHeaderEntrySize: ProgramHeader64Size,
			SectionHeaderEntrySize: SectionHeader64Size,
			ProgramHeaderOffset:    Header64Size,
		},
		ProgramHeaders: []ProgramHeader64{
			ProgramHeader64{
				Type:            1,     // Loadable segment
				Flags:           4 + 1, // Readable & Executable
				VirtualAddress:  address,
				PhysicalAddress: address,
				SizeInFileImage: int64(len(instructions)),
				SizeInMemory:    int64(len(instructions)),
				Align:           programAlign,
			},
		},
		SectionHeaders: []SectionHeader64{
			SectionHeader64{
				NameOffset:      0,
				Type:            1,
				Flags:           2,
				SizeInFileImage: int64(len("Hello World\n")),
				Align:           sectionAlign,
			},
		},
		Instructions: instructions,
		Sections:     []byte("Hello World\n"),
	}

	// Entry point
	program.ProgramHeaderEntryCount = int16(len(program.ProgramHeaders))
	program.SectionHeaderEntryCount = int16(len(program.SectionHeaders))
	program.SectionHeaderOffset = program.ProgramHeaderOffset + int64(program.ProgramHeaderEntryCount)*int64(program.ProgramHeaderEntrySize)
	endOfHeaders := program.SectionHeaderOffset + int64(program.SectionHeaderEntryCount)*int64(program.SectionHeaderEntrySize)
	entryPointInFile := endOfHeaders

	// Padding for instructions
	padding := entryPointInFile % programAlign
	entryPointInFile += padding
	program.InstructionsPadding = bytes.Repeat([]byte{0}, int(padding))
	program.EntryPointInMemory = address + entryPointInFile
	program.ProgramHeaders[0].Offset = entryPointInFile
	program.ProgramHeaders[0].VirtualAddress = program.EntryPointInMemory
	program.ProgramHeaders[0].PhysicalAddress = program.EntryPointInMemory

	// Padding for sections
	endOfInstructions := entryPointInFile + int64(len(instructions))
	padding = endOfInstructions % sectionAlign
	endOfInstructions += padding
	program.SectionsPadding = bytes.Repeat([]byte{0}, int(padding))
	program.SectionHeaders[0].Offset = endOfInstructions
	program.SectionHeaders[0].VirtualAddress = address + (endOfInstructions - entryPointInFile)

	if program.SectionHeaderEntryCount == 0 {
		program.SectionHeaderOffset = 0
	}

	return program
}

// WriteToFile writes the ELF binary to a file.
func (program *ELF64) WriteToFile(fileName string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	program.writeTo(writer)
	return file.Close()
}

//nolint:errcheck
func (program *ELF64) writeTo(writer *bufio.Writer) {
	binary.Write(writer, binary.LittleEndian, &program.Header64)

	for _, programHeader := range program.ProgramHeaders {
		binary.Write(writer, binary.LittleEndian, &programHeader)
	}

	for _, sectionHeader := range program.SectionHeaders {
		binary.Write(writer, binary.LittleEndian, &sectionHeader)
	}

	writer.Write(program.InstructionsPadding)
	writer.Write(program.Instructions)
	writer.Write(program.SectionsPadding)
	writer.Write(program.Sections)
	writer.Flush()
}
