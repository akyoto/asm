package elf

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"

	"github.com/akyoto/asm/stringtable"
	"github.com/akyoto/asm/utils"
)

const (
	address      = 0x400000
	programAlign = 16
	sectionAlign = 16
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

// New creates a new 64-bit ELF binary.
func New(instructions []byte, strings *stringtable.StringTable, sectionPointers []utils.Pointer) *ELF64 {
	elf := &ELF64{
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
		SectionHeaders: []SectionHeader64{},
		Instructions:   instructions,
	}

	elf.AddSection(strings.Bytes())

	// Entry point
	elf.ProgramHeaderEntryCount = int16(len(elf.ProgramHeaders))
	elf.SectionHeaderEntryCount = int16(len(elf.SectionHeaders))
	elf.SectionHeaderOffset = elf.ProgramHeaderOffset + int64(elf.ProgramHeaderEntryCount)*int64(elf.ProgramHeaderEntrySize)
	endOfHeaders := elf.SectionHeaderOffset + int64(elf.SectionHeaderEntryCount)*int64(elf.SectionHeaderEntrySize)
	entryPointInFile := endOfHeaders

	// Padding for instructions
	padding := programAlign - (entryPointInFile % programAlign)
	entryPointInFile += padding
	elf.InstructionsPadding = bytes.Repeat([]byte{0}, int(padding))
	elf.EntryPointInMemory = address + entryPointInFile
	elf.ProgramHeaders[0].Offset = entryPointInFile
	elf.ProgramHeaders[0].VirtualAddress = elf.EntryPointInMemory
	elf.ProgramHeaders[0].PhysicalAddress = elf.EntryPointInMemory

	// Padding for sections
	endOfInstructions := entryPointInFile + int64(len(instructions))
	padding = sectionAlign - (endOfInstructions % sectionAlign)
	endOfInstructions += padding
	elf.SectionsPadding = bytes.Repeat([]byte{0}, int(padding))
	elf.SectionHeaders[0].Offset = endOfInstructions
	elf.SectionHeaders[0].VirtualAddress = address + endOfInstructions

	if elf.SectionHeaderEntryCount == 0 {
		elf.SectionHeaderOffset = 0
	}

	// Apply offsets to all section addresses
	for _, pointer := range sectionPointers {
		binary.LittleEndian.PutUint64(instructions[pointer.Position:pointer.Position+8], uint64(address+endOfInstructions+pointer.Address))
	}

	return elf
}

// AddSection adds a section to the ELF file.
func (elf *ELF64) AddSection(data []byte) {
	elf.SectionHeaders = append(elf.SectionHeaders, SectionHeader64{
		NameOffset:      0,
		Type:            1,
		Flags:           2,
		SizeInFileImage: int64(len(data)),
		Align:           sectionAlign,
	})

	elf.Sections = append(elf.Sections, data...)
}

// WriteToFile writes the ELF binary to a file.
func (elf *ELF64) WriteToFile(fileName string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	elf.writeTo(writer)
	return file.Close()
}

//nolint:errcheck
func (elf *ELF64) writeTo(writer *bufio.Writer) {
	binary.Write(writer, binary.LittleEndian, &elf.Header64)

	for _, programHeader := range elf.ProgramHeaders {
		binary.Write(writer, binary.LittleEndian, &programHeader)
	}

	for _, sectionHeader := range elf.SectionHeaders {
		binary.Write(writer, binary.LittleEndian, &sectionHeader)
	}

	writer.Write(elf.InstructionsPadding)
	writer.Write(elf.Instructions)
	writer.Write(elf.SectionsPadding)
	writer.Write(elf.Sections)
	writer.Flush()
}
