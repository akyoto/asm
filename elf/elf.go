package elf

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"

	"github.com/akyoto/asm/sections"
)

const (
	baseAddress  = 0x400000
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
func New(instructions []byte, strings *sections.Strings, stringPointers []sections.Pointer) *ELF64 {
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

	elf.AddProgram(instructions, ProgramTypeLOAD, ProgramFlagsExecutable)
	elf.AddSection(strings.Bytes(), SectionTypePROGBITS, SectionFlagsAllocate)
	// elf.AddSection(nil, SectionTypeNOBITS, SectionFlagsAllocate|SectionFlagsWritable)

	// Header count
	elf.ProgramHeaderEntryCount = int16(len(elf.Programs))
	elf.SectionHeaderEntryCount = int16(len(elf.Sections))

	// Entry point
	endOfProgramHeaders := elf.ProgramHeaderOffset + int64(elf.ProgramHeaderEntryCount)*int64(elf.ProgramHeaderEntrySize)
	elf.SectionHeaderOffset = endOfProgramHeaders
	endOfSectionHeaders := elf.SectionHeaderOffset + int64(elf.SectionHeaderEntryCount)*int64(elf.SectionHeaderEntrySize)
	offset := endOfSectionHeaders

	// Padding for programs
	for _, program := range elf.Programs {
		padding := calculatePadding(offset, programAlign)
		offset += padding

		if elf.EntryPointInMemory == 0 {
			elf.EntryPointInMemory = baseAddress + offset
		}

		program.Padding = bytes.Repeat([]byte{0}, int(padding))
		program.Header.Offset = offset
		program.Header.VirtualAddress = baseAddress + offset
		program.Header.PhysicalAddress = baseAddress + offset
		offset += int64(len(program.Data))
	}

	// Padding for sections
	for _, section := range elf.Sections {
		padding := calculatePadding(offset, sectionAlign)
		offset += padding
		section.Padding = bytes.Repeat([]byte{0}, int(padding))
		section.Header.Offset = offset
		section.Header.VirtualAddress = baseAddress + offset
		offset += int64(len(section.Data))
	}

	// Special case so that readelf doesn't complain
	if elf.SectionHeaderEntryCount == 0 {
		elf.SectionHeaderOffset = 0
	}

	// Add section offset to all string addresses
	for _, pointer := range stringPointers {
		oldAddressSlice := instructions[pointer.Position : pointer.Position+8]
		newAddress := uint64(baseAddress + elf.Sections[0].Header.Offset + pointer.Address)
		binary.LittleEndian.PutUint64(oldAddressSlice, newAddress)
	}

	return elf
}

// AddProgram adds a section to the ELF file.
func (elf *ELF64) AddProgram(data []byte, typ ProgramType, flags ProgramFlags) {
	elf.Programs = append(elf.Programs, &Program{
		Header: ProgramHeader64{
			Type:            typ,
			Flags:           flags,
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
