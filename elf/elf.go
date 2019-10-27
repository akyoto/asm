package elf

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"

	"github.com/akyoto/asm/sections"
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
	Programs []*Program
	Sections []*Section
}

// New creates a new 64-bit ELF binary.
func New(instructions []byte, strings *sections.Strings, stringPointers []utils.Pointer) *ELF64 {
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
	}

	elf.AddProgram(instructions, ProgramTypeLOAD, ProgramFlagsExecutable|ProgramFlagsReadable)
	elf.AddSection(strings.Bytes(), SectionTypePROGBITS, SectionFlagsAllocate)

	// Entry point
	elf.ProgramHeaderEntryCount = int16(len(elf.Programs))
	elf.SectionHeaderEntryCount = int16(len(elf.Sections))
	elf.SectionHeaderOffset = elf.ProgramHeaderOffset + int64(elf.ProgramHeaderEntryCount)*int64(elf.ProgramHeaderEntrySize)
	endOfHeaders := elf.SectionHeaderOffset + int64(elf.SectionHeaderEntryCount)*int64(elf.SectionHeaderEntrySize)
	entryPointInFile := endOfHeaders

	// Padding for instructions
	padding := calculatePadding(entryPointInFile, programAlign)
	entryPointInFile += padding
	elf.Programs[0].Padding = bytes.Repeat([]byte{0}, int(padding))
	elf.EntryPointInMemory = address + entryPointInFile
	elf.Programs[0].Header.Offset = entryPointInFile
	elf.Programs[0].Header.VirtualAddress = elf.EntryPointInMemory
	elf.Programs[0].Header.PhysicalAddress = elf.EntryPointInMemory

	// Padding for sections
	endOfInstructions := entryPointInFile + int64(len(instructions))
	padding = calculatePadding(endOfInstructions, sectionAlign)
	endOfInstructions += padding
	elf.Sections[0].Padding = bytes.Repeat([]byte{0}, int(padding))
	elf.Sections[0].Header.Offset = endOfInstructions
	elf.Sections[0].Header.VirtualAddress = address + endOfInstructions

	if elf.SectionHeaderEntryCount == 0 {
		elf.SectionHeaderOffset = 0
	}

	// Apply offsets to all string addresses
	for _, pointer := range stringPointers {
		oldAddress := instructions[pointer.Position : pointer.Position+8]
		newAddress := uint64(address + endOfInstructions + pointer.Address)
		binary.LittleEndian.PutUint64(oldAddress, newAddress)
	}

	return elf
}

// AddProgram adds a section to the ELF file.
func (elf *ELF64) AddProgram(data []byte, typ ProgramType, flags ProgramFlags) {
	elf.Programs = append(elf.Programs, &Program{
		Header: ProgramHeader64{
			Type:            typ,
			Flags:           flags,
			VirtualAddress:  address,
			PhysicalAddress: address,
			SizeInFileImage: int64(len(data)),
			SizeInMemory:    int64(len(data)),
			Align:           programAlign,
		},
		Data: data,
	})
}

// AddSection adds a section to the ELF file.
func (elf *ELF64) AddSection(data []byte, typ SectionType, flags SectionFlags) {
	elf.Sections = append(elf.Sections, &Section{
		Header: SectionHeader64{
			NameOffset:      0,
			Type:            typ,
			Flags:           flags,
			SizeInFileImage: int64(len(data)),
			Align:           sectionAlign,
		},
		Data: data,
	})
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

	for _, program := range elf.Programs {
		binary.Write(writer, binary.LittleEndian, &program.Header)
	}

	for _, section := range elf.Sections {
		binary.Write(writer, binary.LittleEndian, &section.Header)
	}

	for _, program := range elf.Programs {
		writer.Write(program.Padding)
		writer.Write(program.Data)
	}

	for _, section := range elf.Sections {
		writer.Write(section.Padding)
		writer.Write(section.Data)
	}

	writer.Flush()
}

func calculatePadding(n int64, align int64) int64 {
	return align - (n % align)
}
